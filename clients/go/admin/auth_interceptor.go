package admin

import (
	"context"
	"fmt"

	"github.com/flyteorg/flyteidl/clients/go/admin/cache"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/flyteorg/flytestdlib/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

// ResetTokenSource will attempt to build a TokenSource given the anonymously available information exposed by the server.
// Once established, it'll invoke CustomHeaderTokenSource.Reset() on tokenSource to populate it with the appropriate values.
func ResetTokenSource(ctx context.Context, cfg *Config, tokenCache cache.TokenCache, tokenSource *CustomHeaderTokenSource) error {
	authMetadataClient, err := InitializeAuthMetadataClient(ctx, cfg)
	if err != nil {
		// EngHabu: I don't know why did we use to panic here. Generally speaking, libraries should error on the side of
		// 	returning errors instead.
		return fmt.Errorf("failed to initialize Auth Metadata Client. Error: %w", err)
	}

	tokenSourceProvider, err := NewTokenSourceProvider(ctx, cfg, tokenCache, authMetadataClient)
	if err != nil {
		return fmt.Errorf("failed to initialize token source provider. Err: %w", err)
	}

	clientMetadata, err := authMetadataClient.GetPublicClientConfig(ctx, &service.PublicClientAuthConfigRequest{})
	if err != nil {
		return fmt.Errorf("failed to fetch client metadata. Error: %v", err)
	}

	tSource, err := tokenSourceProvider.GetTokenSource(ctx)
	if err != nil {
		return err
	}

	tokenSource.Reset(tSource, clientMetadata.AuthorizationMetadataKey, cfg.UseInsecureConnection)
	return nil
}

// newAuthInterceptor creates a new grpc.UnaryClientInterceptor that forwards the grpc call and inspects the error.
// If an auth error is detected, it'll panic the process. This is useful for scenarios when auth is enabled after the client
// has been up or if a client was inadvertently (due to transient network failures... etc.) built without auth credentials.
// It will first invoke the grpc pipeline (to proceed with the request) with no modifications. It's expected for the grpc
// pipeline to already have a grpc.WithPerRPCCredentials() DialOption. If the tokenSource has already been initialized,
// it'll take care of refreshing when it expires... etc.
// If the first invocation fails with an auth error, this interceptor will then attempt to establish a token source once
// more. It'll fail hard if it couldn't do so (i.e. it will no longer attempt to send an unauthenticated request). Once
// a token source has been created, it'll invoke the grpc pipeline again, this time the grpc.PerRPCCredentials should
// be able to find and acquire a valid AccessToken to annotate the request with.
func newAuthInterceptor(cfg *Config, tokenCache cache.TokenCache, tokenSource *CustomHeaderTokenSource) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if st, ok := status.FromError(err); ok {
			// If the error we receive from executing the request expects
			if st.Code() == codes.PermissionDenied || st.Code() == codes.Unauthenticated {
				logger.Debugf(ctx, "Request failed due to [%v]. Attempting to establish an authenticated connection and trying again.", st.Code())
				newErr := ResetTokenSource(ctx, cfg, tokenCache, tokenSource)
				if newErr != nil {
					return fmt.Errorf("authentication error! Original Error: %v, Auth Error: %w", err, newErr)
				}

				return invoker(ctx, method, req, reply, cc, opts...)
			}
		}

		return err
	}
}
