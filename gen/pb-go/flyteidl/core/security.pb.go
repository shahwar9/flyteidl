// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/security.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Secret_MountType int32

const (
	// ENV_VAR indicates the secret needs to be mounted as an environment variable.
	Secret_ENV_VAR Secret_MountType = 0
	// FILE indicates the secret needs to be mounted as a file.
	Secret_FILE Secret_MountType = 1
)

var Secret_MountType_name = map[int32]string{
	0: "ENV_VAR",
	1: "FILE",
}

var Secret_MountType_value = map[string]int32{
	"ENV_VAR": 0,
	"FILE":    1,
}

func (x Secret_MountType) String() string {
	return proto.EnumName(Secret_MountType_name, int32(x))
}

func (Secret_MountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{0, 0}
}

// Type of the token requested.
type OAuth2TokenRequest_Type int32

const (
	// CLIENT_CREDENTIALS indicates a 2-legged OAuth token requested using client credentials.
	OAuth2TokenRequest_CLIENT_CREDENTIALS OAuth2TokenRequest_Type = 0
)

var OAuth2TokenRequest_Type_name = map[int32]string{
	0: "CLIENT_CREDENTIALS",
}

var OAuth2TokenRequest_Type_value = map[string]int32{
	"CLIENT_CREDENTIALS": 0,
}

func (x OAuth2TokenRequest_Type) String() string {
	return proto.EnumName(OAuth2TokenRequest_Type_name, int32(x))
}

func (OAuth2TokenRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{3, 0}
}

// Secret encapsulates information about the secret a task needs to proceed. An environment variable
// FLYTE_SECRETS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
// secrets are passed through environment variables.
// FLYTE_SECRETS_DEFAULT_DIR will be passed to indicate the prefix of the path where secrets will be mounted if secrets
// are passed through file mounts.
type Secret struct {
	// The name of the secret to mount. This has to match an existing secret in the system. It's up to the implementation
	// of the secret management system to require case sensitivity.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// mount_requirement is optional. Indicates where the secret has to be mounted. If provided, the execution will fail
	// if the underlying key management system cannot satisfy that requirement. If not provided, the default location
	// will depend on the key management system.
	// +optional
	MountRequirement     Secret_MountType `protobuf:"varint,2,opt,name=mount_requirement,json=mountRequirement,proto3,enum=flyteidl.core.Secret_MountType" json:"mount_requirement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{0}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetMountRequirement() Secret_MountType {
	if m != nil {
		return m.MountRequirement
	}
	return Secret_ENV_VAR
}

// OAuth2Client encapsulates OAuth2 Client Credentials to be used when making calls on behalf of that task.
type OAuth2Client struct {
	// client_id is the public id for the client to use. The system will not perform any pre-auth validation that the
	// secret requested matches the client_id indicated here.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// client_secret is a reference to the secret used to authenticate the OAuth2 client.
	ClientSecret         *Secret  `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuth2Client) Reset()         { *m = OAuth2Client{} }
func (m *OAuth2Client) String() string { return proto.CompactTextString(m) }
func (*OAuth2Client) ProtoMessage()    {}
func (*OAuth2Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{1}
}

func (m *OAuth2Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2Client.Unmarshal(m, b)
}
func (m *OAuth2Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2Client.Marshal(b, m, deterministic)
}
func (m *OAuth2Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2Client.Merge(m, src)
}
func (m *OAuth2Client) XXX_Size() int {
	return xxx_messageInfo_OAuth2Client.Size(m)
}
func (m *OAuth2Client) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2Client.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2Client proto.InternalMessageInfo

func (m *OAuth2Client) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OAuth2Client) GetClientSecret() *Secret {
	if m != nil {
		return m.ClientSecret
	}
	return nil
}

// Identity encapsulates the various security identities a task can run as. It's up to the underlying plugin to pick the
// right identity for the execution environment.
type Identity struct {
	// iam_role references the fully qualified name of Identity & Access Management role to impersonate.
	IamRole string `protobuf:"bytes,1,opt,name=iam_role,json=iamRole,proto3" json:"iam_role,omitempty"`
	// k8s_service_account references a kubernetes service account to impersonate.
	K8SServiceAccount string `protobuf:"bytes,2,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// oauth2_client references an oauth2 client. Backend plugins can use this information to impersonate the client when
	// making external calls.
	Oauth2Client         *OAuth2Client `protobuf:"bytes,3,opt,name=oauth2_client,json=oauth2Client,proto3" json:"oauth2_client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{2}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetIamRole() string {
	if m != nil {
		return m.IamRole
	}
	return ""
}

func (m *Identity) GetK8SServiceAccount() string {
	if m != nil {
		return m.K8SServiceAccount
	}
	return ""
}

func (m *Identity) GetOauth2Client() *OAuth2Client {
	if m != nil {
		return m.Oauth2Client
	}
	return nil
}

// OAuth2TokenRequest encapsulates information needed to request an OAuth2 token.
// FLYTE_TOKENS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
// tokens are passed through environment variables.
// FLYTE_TOKENS_PATH_PREFIX will be passed to indicate the prefix of the path where secrets will be mounted if tokens
// are passed through file mounts.
type OAuth2TokenRequest struct {
	// name indicates a unique id for the token request within this task token requests. It'll be used as a suffix for
	// environment variables and as a filename for mounting tokens as files.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type indicates the type of the request to make. Defaults to CLIENT_CREDENTIALS.
	Type OAuth2TokenRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=flyteidl.core.OAuth2TokenRequest_Type" json:"type,omitempty"`
	// client references the client_id/secret to use to request the OAuth2 token.
	Client *OAuth2Client `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// idp_discovery_endpoint references the discovery endpoint used to retrieve token endpoint and other related
	// information.
	// +optional
	IdpDiscoveryEndpoint string `protobuf:"bytes,4,opt,name=idp_discovery_endpoint,json=idpDiscoveryEndpoint,proto3" json:"idp_discovery_endpoint,omitempty"`
	// token_endpoint references the token issuance endpoint. If idp_discovery_endpoint is not provided, this parameter is
	// mandatory.
	// +optional
	TokenEndpoint        string   `protobuf:"bytes,5,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuth2TokenRequest) Reset()         { *m = OAuth2TokenRequest{} }
func (m *OAuth2TokenRequest) String() string { return proto.CompactTextString(m) }
func (*OAuth2TokenRequest) ProtoMessage()    {}
func (*OAuth2TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{3}
}

func (m *OAuth2TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2TokenRequest.Unmarshal(m, b)
}
func (m *OAuth2TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2TokenRequest.Marshal(b, m, deterministic)
}
func (m *OAuth2TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2TokenRequest.Merge(m, src)
}
func (m *OAuth2TokenRequest) XXX_Size() int {
	return xxx_messageInfo_OAuth2TokenRequest.Size(m)
}
func (m *OAuth2TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2TokenRequest proto.InternalMessageInfo

func (m *OAuth2TokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuth2TokenRequest) GetType() OAuth2TokenRequest_Type {
	if m != nil {
		return m.Type
	}
	return OAuth2TokenRequest_CLIENT_CREDENTIALS
}

func (m *OAuth2TokenRequest) GetClient() *OAuth2Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *OAuth2TokenRequest) GetIdpDiscoveryEndpoint() string {
	if m != nil {
		return m.IdpDiscoveryEndpoint
	}
	return ""
}

func (m *OAuth2TokenRequest) GetTokenEndpoint() string {
	if m != nil {
		return m.TokenEndpoint
	}
	return ""
}

// SecurityContext holds security attributes that apply to tasks.
type SecurityContext struct {
	// run_as encapsulates the identity a pod should run as. If the task fills in multiple fields here, it'll be up to the
	// backend plugin to choose the appropriate identity for the execution engine the task will run on.
	RunAs *Identity `protobuf:"bytes,1,opt,name=run_as,json=runAs,proto3" json:"run_as,omitempty"`
	// secrets indicate the list of secrets the task needs in order to proceed. Secrets will be mounted/passed to the
	// pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
	// Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
	// to the secret) and to pass it to the remote execution engine.
	Secrets []*Secret `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty"`
	// tokens indicate the list of token requests the task needs in order to proceed. Tokens will be mounted/passed to the
	// pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
	// Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
	// to the secret) and to pass it to the remote execution engine.
	Tokens               []*OAuth2TokenRequest `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SecurityContext) Reset()         { *m = SecurityContext{} }
func (m *SecurityContext) String() string { return proto.CompactTextString(m) }
func (*SecurityContext) ProtoMessage()    {}
func (*SecurityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_0996009b6d39c02f, []int{4}
}

func (m *SecurityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityContext.Unmarshal(m, b)
}
func (m *SecurityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityContext.Marshal(b, m, deterministic)
}
func (m *SecurityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityContext.Merge(m, src)
}
func (m *SecurityContext) XXX_Size() int {
	return xxx_messageInfo_SecurityContext.Size(m)
}
func (m *SecurityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityContext.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityContext proto.InternalMessageInfo

func (m *SecurityContext) GetRunAs() *Identity {
	if m != nil {
		return m.RunAs
	}
	return nil
}

func (m *SecurityContext) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecurityContext) GetTokens() []*OAuth2TokenRequest {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.core.Secret_MountType", Secret_MountType_name, Secret_MountType_value)
	proto.RegisterEnum("flyteidl.core.OAuth2TokenRequest_Type", OAuth2TokenRequest_Type_name, OAuth2TokenRequest_Type_value)
	proto.RegisterType((*Secret)(nil), "flyteidl.core.Secret")
	proto.RegisterType((*OAuth2Client)(nil), "flyteidl.core.OAuth2Client")
	proto.RegisterType((*Identity)(nil), "flyteidl.core.Identity")
	proto.RegisterType((*OAuth2TokenRequest)(nil), "flyteidl.core.OAuth2TokenRequest")
	proto.RegisterType((*SecurityContext)(nil), "flyteidl.core.SecurityContext")
}

func init() { proto.RegisterFile("flyteidl/core/security.proto", fileDescriptor_0996009b6d39c02f) }

var fileDescriptor_0996009b6d39c02f = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xb6, 0xae, 0x1f, 0xb7, 0xed, 0xe8, 0x0c, 0x8c, 0xa0, 0x21, 0x28, 0x91, 0x40, 0x7b,
	0x21, 0x91, 0xba, 0x09, 0x8d, 0x3d, 0x51, 0xba, 0x20, 0x55, 0x2a, 0x45, 0x4a, 0xab, 0x3d, 0xf0,
	0x12, 0xa5, 0xc9, 0xa5, 0xb3, 0xda, 0xd8, 0xc1, 0x71, 0x26, 0xf2, 0x27, 0x78, 0x84, 0x5f, 0xc1,
	0x7f, 0x44, 0xb1, 0x93, 0x7d, 0xa9, 0x08, 0xde, 0x6c, 0x9f, 0x73, 0x7c, 0xcf, 0xfd, 0x82, 0x67,
	0x5f, 0xd7, 0xb9, 0x44, 0x1a, 0xad, 0x9d, 0x90, 0x0b, 0x74, 0x52, 0x0c, 0x33, 0x41, 0x65, 0x6e,
	0x27, 0x82, 0x4b, 0x4e, 0xba, 0x15, 0x6a, 0x17, 0xa8, 0xf5, 0xc3, 0x80, 0xfa, 0x0c, 0x43, 0x81,
	0x92, 0x10, 0xa8, 0xb1, 0x20, 0x46, 0xd3, 0xe8, 0x1b, 0x47, 0x2d, 0x4f, 0x9d, 0xc9, 0x04, 0xf6,
	0x63, 0x9e, 0x31, 0xe9, 0x0b, 0xfc, 0x96, 0x51, 0x81, 0x31, 0x32, 0x69, 0x6e, 0xf7, 0x8d, 0xa3,
	0xbd, 0xc1, 0x0b, 0xfb, 0xce, 0x4f, 0xb6, 0xfe, 0xc5, 0xfe, 0x54, 0xd0, 0xe7, 0x79, 0x82, 0x5e,
	0x4f, 0x29, 0xbd, 0x1b, 0xa1, 0x65, 0x41, 0xeb, 0x1a, 0x26, 0x6d, 0x68, 0xb8, 0xd3, 0x0b, 0xff,
	0x62, 0xe8, 0xf5, 0xb6, 0x48, 0x13, 0x6a, 0x1f, 0xc7, 0x13, 0xb7, 0x67, 0x58, 0x4b, 0xe8, 0x7c,
	0x1e, 0x66, 0xf2, 0x72, 0x30, 0x5a, 0x53, 0x64, 0x92, 0x1c, 0x42, 0x2b, 0x54, 0x27, 0x9f, 0x46,
	0xa5, 0xb5, 0xa6, 0x7e, 0x18, 0x47, 0xe4, 0x0c, 0xba, 0x25, 0x98, 0xaa, 0xe8, 0xca, 0x5a, 0x7b,
	0xf0, 0x78, 0xa3, 0x35, 0xaf, 0xa3, 0xb9, 0xfa, 0x66, 0xfd, 0x32, 0xa0, 0x39, 0x8e, 0x90, 0x49,
	0x2a, 0x73, 0xf2, 0x14, 0x9a, 0x34, 0x88, 0x7d, 0xc1, 0xd7, 0x55, 0xfe, 0x0d, 0x1a, 0xc4, 0x1e,
	0x5f, 0x23, 0xb1, 0xe1, 0xe1, 0xea, 0x34, 0xf5, 0x53, 0x14, 0x57, 0x34, 0x44, 0x3f, 0x08, 0xc3,
	0x22, 0x05, 0x15, 0xa9, 0xe5, 0xed, 0xaf, 0x4e, 0xd3, 0x99, 0x46, 0x86, 0x1a, 0x20, 0xef, 0xa1,
	0xcb, 0x83, 0x22, 0x01, 0x5f, 0x87, 0x33, 0x77, 0x94, 0xa7, 0xc3, 0x7b, 0x9e, 0x6e, 0x27, 0xe9,
	0x75, 0xb4, 0x42, 0xdf, 0xac, 0x9f, 0xdb, 0x40, 0x34, 0x3c, 0xe7, 0x2b, 0x64, 0x45, 0x05, 0x31,
	0xdd, 0xdc, 0x9f, 0x33, 0xa8, 0xc9, 0x3c, 0xc1, 0xb2, 0x25, 0xaf, 0x37, 0xc6, 0xb8, 0xfd, 0x89,
	0xad, 0x3a, 0xa3, 0x34, 0xe4, 0x18, 0xea, 0xff, 0xef, 0xb0, 0xa4, 0x92, 0x13, 0x38, 0xa0, 0x51,
	0xe2, 0x47, 0x34, 0x0d, 0xf9, 0x15, 0x8a, 0xdc, 0x47, 0x16, 0x25, 0x9c, 0x32, 0x69, 0xd6, 0x94,
	0xad, 0x47, 0x34, 0x4a, 0xce, 0x2b, 0xd0, 0x2d, 0x31, 0xf2, 0x0a, 0xf6, 0x64, 0xe1, 0xe2, 0x86,
	0xbd, 0xab, 0xd8, 0x5d, 0xf5, 0x5a, 0xd1, 0xac, 0xe7, 0x50, 0x53, 0xa3, 0x71, 0x00, 0x64, 0x34,
	0x19, 0xbb, 0xd3, 0xb9, 0x3f, 0xf2, 0xdc, 0x73, 0x77, 0x3a, 0x1f, 0x0f, 0x27, 0xb3, 0xde, 0x96,
	0xf5, 0xdb, 0x80, 0x07, 0xb3, 0x72, 0x9c, 0x47, 0x9c, 0x49, 0xfc, 0x2e, 0x89, 0x0d, 0x75, 0x91,
	0x31, 0x3f, 0x48, 0x55, 0x5d, 0xda, 0x83, 0x27, 0xf7, 0xb2, 0xa8, 0x5a, 0xec, 0xed, 0x8a, 0x8c,
	0x0d, 0x53, 0xe2, 0x40, 0x43, 0xcf, 0x4a, 0x6a, 0x6e, 0xf7, 0x77, 0xfe, 0x3e, 0x2c, 0x15, 0x8b,
	0xbc, 0x83, 0xba, 0x72, 0x99, 0x9a, 0x3b, 0x8a, 0xff, 0xf2, 0x9f, 0x45, 0xf6, 0x4a, 0xc1, 0x87,
	0xb7, 0x5f, 0x4e, 0x96, 0x54, 0x5e, 0x66, 0x0b, 0x3b, 0xe4, 0xb1, 0xa3, 0x64, 0x5c, 0x2c, 0x9d,
	0xeb, 0xfd, 0x5c, 0x22, 0x73, 0x92, 0xc5, 0x9b, 0x25, 0x77, 0xee, 0xac, 0xec, 0xa2, 0xae, 0x56,
	0xf5, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xe9, 0xb7, 0xcf, 0xca, 0x03, 0x00, 0x00,
}
