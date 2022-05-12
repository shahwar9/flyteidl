// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: flyteidl/service/signal.proto
#ifndef GRPC_flyteidl_2fservice_2fsignal_2eproto__INCLUDED
#define GRPC_flyteidl_2fservice_2fsignal_2eproto__INCLUDED

#include "flyteidl/service/signal.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_generic_service.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace grpc_impl {
class Channel;
class CompletionQueue;
class ServerCompletionQueue;
}  // namespace grpc_impl

namespace grpc {
namespace experimental {
template <typename RequestT, typename ResponseT>
class MessageAllocator;
}  // namespace experimental
}  // namespace grpc_impl

namespace grpc {
class ServerContext;
}  // namespace grpc

namespace flyteidl {
namespace service {

// TODO hamersaw - document
class SignalService final {
 public:
  static constexpr char const* service_full_name() {
    return "flyteidl.service.SignalService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // Create and upload a :ref:`ref_flyteidl.admin.Signal` definition
    virtual ::grpc::Status CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::flyteidl::admin::SignalCreateResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>> AsyncCreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>>(AsyncCreateSignalRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>> PrepareAsyncCreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>>(PrepareAsyncCreateSignalRaw(context, request, cq));
    }
    // Fetches a :ref:`ref_flyteidl.admin.Signal`.
    virtual ::grpc::Status GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::flyteidl::admin::Signal* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>> AsyncGetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>>(AsyncGetSignalRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>> PrepareAsyncGetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>>(PrepareAsyncGetSignalRaw(context, request, cq));
    }
    class experimental_async_interface {
     public:
      virtual ~experimental_async_interface() {}
      // Create and upload a :ref:`ref_flyteidl.admin.Signal` definition
      virtual void CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void CreateSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::SignalCreateResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      virtual void CreateSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::SignalCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      // Fetches a :ref:`ref_flyteidl.admin.Signal`.
      virtual void GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::Signal* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      virtual void GetSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::Signal* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
    };
    virtual class experimental_async_interface* experimental_async() { return nullptr; }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>* AsyncCreateSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::SignalCreateResponse>* PrepareAsyncCreateSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>* AsyncGetSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::flyteidl::admin::Signal>* PrepareAsyncGetSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::flyteidl::admin::SignalCreateResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>> AsyncCreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>>(AsyncCreateSignalRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>> PrepareAsyncCreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>>(PrepareAsyncCreateSignalRaw(context, request, cq));
    }
    ::grpc::Status GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::flyteidl::admin::Signal* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>> AsyncGetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>>(AsyncGetSignalRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>> PrepareAsyncGetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>>(PrepareAsyncGetSignalRaw(context, request, cq));
    }
    class experimental_async final :
      public StubInterface::experimental_async_interface {
     public:
      void CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response, std::function<void(::grpc::Status)>) override;
      void CreateSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::SignalCreateResponse* response, std::function<void(::grpc::Status)>) override;
      void CreateSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void CreateSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::SignalCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response, std::function<void(::grpc::Status)>) override;
      void GetSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::Signal* response, std::function<void(::grpc::Status)>) override;
      void GetSignal(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void GetSignal(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::admin::Signal* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit experimental_async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class experimental_async_interface* experimental_async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class experimental_async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>* AsyncCreateSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::SignalCreateResponse>* PrepareAsyncCreateSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalCreateRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>* AsyncGetSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::flyteidl::admin::Signal>* PrepareAsyncGetSignalRaw(::grpc::ClientContext* context, const ::flyteidl::admin::SignalGetRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_CreateSignal_;
    const ::grpc::internal::RpcMethod rpcmethod_GetSignal_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // Create and upload a :ref:`ref_flyteidl.admin.Signal` definition
    virtual ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response);
    // Fetches a :ref:`ref_flyteidl.admin.Signal`.
    virtual ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_CreateSignal() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestCreateSignal(::grpc::ServerContext* context, ::flyteidl::admin::SignalCreateRequest* request, ::grpc::ServerAsyncResponseWriter< ::flyteidl::admin::SignalCreateResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetSignal() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetSignal(::grpc::ServerContext* context, ::flyteidl::admin::SignalGetRequest* request, ::grpc::ServerAsyncResponseWriter< ::flyteidl::admin::Signal>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_CreateSignal<WithAsyncMethod_GetSignal<Service > > AsyncService;
  template <class BaseClass>
  class ExperimentalWithCallbackMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithCallbackMethod_CreateSignal() {
      ::grpc::Service::experimental().MarkMethodCallback(0,
        new ::grpc::internal::CallbackUnaryHandler< ::flyteidl::admin::SignalCreateRequest, ::flyteidl::admin::SignalCreateResponse>(
          [this](::grpc::ServerContext* context,
                 const ::flyteidl::admin::SignalCreateRequest* request,
                 ::flyteidl::admin::SignalCreateResponse* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   return this->CreateSignal(context, request, response, controller);
                 }));
    }
    void SetMessageAllocatorFor_CreateSignal(
        ::grpc::experimental::MessageAllocator< ::flyteidl::admin::SignalCreateRequest, ::flyteidl::admin::SignalCreateResponse>* allocator) {
      static_cast<::grpc::internal::CallbackUnaryHandler< ::flyteidl::admin::SignalCreateRequest, ::flyteidl::admin::SignalCreateResponse>*>(
          ::grpc::Service::experimental().GetHandler(0))
              ->SetMessageAllocator(allocator);
    }
    ~ExperimentalWithCallbackMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class ExperimentalWithCallbackMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithCallbackMethod_GetSignal() {
      ::grpc::Service::experimental().MarkMethodCallback(1,
        new ::grpc::internal::CallbackUnaryHandler< ::flyteidl::admin::SignalGetRequest, ::flyteidl::admin::Signal>(
          [this](::grpc::ServerContext* context,
                 const ::flyteidl::admin::SignalGetRequest* request,
                 ::flyteidl::admin::Signal* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   return this->GetSignal(context, request, response, controller);
                 }));
    }
    void SetMessageAllocatorFor_GetSignal(
        ::grpc::experimental::MessageAllocator< ::flyteidl::admin::SignalGetRequest, ::flyteidl::admin::Signal>* allocator) {
      static_cast<::grpc::internal::CallbackUnaryHandler< ::flyteidl::admin::SignalGetRequest, ::flyteidl::admin::Signal>*>(
          ::grpc::Service::experimental().GetHandler(1))
              ->SetMessageAllocator(allocator);
    }
    ~ExperimentalWithCallbackMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  typedef ExperimentalWithCallbackMethod_CreateSignal<ExperimentalWithCallbackMethod_GetSignal<Service > > ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_CreateSignal() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetSignal() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithRawMethod_CreateSignal() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestCreateSignal(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithRawMethod_GetSignal() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetSignal(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class ExperimentalWithRawCallbackMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithRawCallbackMethod_CreateSignal() {
      ::grpc::Service::experimental().MarkMethodRawCallback(0,
        new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
          [this](::grpc::ServerContext* context,
                 const ::grpc::ByteBuffer* request,
                 ::grpc::ByteBuffer* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   this->CreateSignal(context, request, response, controller);
                 }));
    }
    ~ExperimentalWithRawCallbackMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void CreateSignal(::grpc::ServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class ExperimentalWithRawCallbackMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithRawCallbackMethod_GetSignal() {
      ::grpc::Service::experimental().MarkMethodRawCallback(1,
        new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
          [this](::grpc::ServerContext* context,
                 const ::grpc::ByteBuffer* request,
                 ::grpc::ByteBuffer* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   this->GetSignal(context, request, response, controller);
                 }));
    }
    ~ExperimentalWithRawCallbackMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetSignal(::grpc::ServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_CreateSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_CreateSignal() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler< ::flyteidl::admin::SignalCreateRequest, ::flyteidl::admin::SignalCreateResponse>(std::bind(&WithStreamedUnaryMethod_CreateSignal<BaseClass>::StreamedCreateSignal, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_CreateSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status CreateSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalCreateRequest* request, ::flyteidl::admin::SignalCreateResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedCreateSignal(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::flyteidl::admin::SignalCreateRequest,::flyteidl::admin::SignalCreateResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetSignal : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetSignal() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler< ::flyteidl::admin::SignalGetRequest, ::flyteidl::admin::Signal>(std::bind(&WithStreamedUnaryMethod_GetSignal<BaseClass>::StreamedGetSignal, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetSignal() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetSignal(::grpc::ServerContext* context, const ::flyteidl::admin::SignalGetRequest* request, ::flyteidl::admin::Signal* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetSignal(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::flyteidl::admin::SignalGetRequest,::flyteidl::admin::Signal>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_CreateSignal<WithStreamedUnaryMethod_GetSignal<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_CreateSignal<WithStreamedUnaryMethod_GetSignal<Service > > StreamedService;
};

}  // namespace service
}  // namespace flyteidl


#endif  // GRPC_flyteidl_2fservice_2fsignal_2eproto__INCLUDED