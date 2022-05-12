// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/signal.proto

package flyteidl.service;

public final class Signal {
  private Signal() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\035flyteidl/service/signal.proto\022\020flyteid" +
      "l.service\032\034google/api/annotations.proto\032" +
      "\033flyteidl/admin/signal.proto\032,protoc-gen" +
      "-swagger/options/annotations.proto2\270\004\n\rS" +
      "ignalService\022\302\002\n\014CreateSignal\022#.flyteidl" +
      ".admin.SignalCreateRequest\032$.flyteidl.ad" +
      "min.SignalCreateResponse\"\346\001\202\323\344\223\002\024\"\017/api/" +
      "v1/signals:\001*\222A\310\001\032\033Create a signal defin" +
      "ition.JB\n\003400\022;\n9Returned for bad reques" +
      "t that may have failed validation.Je\n\00340" +
      "9\022^\n\\Returned for a request that referen" +
      "ces an identical entity that has already" +
      " been registered.\022\341\001\n\tGetSignal\022 .flytei" +
      "dl.admin.SignalGetRequest\032\026.flyteidl.adm" +
      "in.Signal\"\231\001\202\323\344\223\002j\022h/api/v1/signals/{id." +
      "execution_id.project}/{id.execution_id.d" +
      "omain}/{id.execution_id.name}/{id.signal" +
      "_id}\222A&\032$Retrieve an existing node execu" +
      "tion.B9Z7github.com/flyteorg/flyteidl/ge" +
      "n/pb-go/flyteidl/serviceb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          flyteidl.admin.SignalOuterClass.getDescriptor(),
          grpc.gateway.protoc_gen_swagger.options.Annotations.getDescriptor(),
        }, assigner);
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    registry.add(grpc.gateway.protoc_gen_swagger.options.Annotations.openapiv2Operation);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    flyteidl.admin.SignalOuterClass.getDescriptor();
    grpc.gateway.protoc_gen_swagger.options.Annotations.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}