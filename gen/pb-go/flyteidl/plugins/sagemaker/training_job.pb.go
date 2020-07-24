// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/training_job.proto

package plugins

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

type InputMode int32

const (
	InputMode_FILE InputMode = 0
	InputMode_PIPE InputMode = 1
)

var InputMode_name = map[int32]string{
	0: "FILE",
	1: "PIPE",
}

var InputMode_value = map[string]int32{
	"FILE": 0,
	"PIPE": 1,
}

func (x InputMode) String() string {
	return proto.EnumName(InputMode_name, int32(x))
}

func (InputMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0}
}

type AlgorithmName int32

const (
	AlgorithmName_CUSTOM  AlgorithmName = 0
	AlgorithmName_XGBOOST AlgorithmName = 1
)

var AlgorithmName_name = map[int32]string{
	0: "CUSTOM",
	1: "XGBOOST",
}

var AlgorithmName_value = map[string]int32{
	"CUSTOM":  0,
	"XGBOOST": 1,
}

func (x AlgorithmName) String() string {
	return proto.EnumName(AlgorithmName_name, int32(x))
}

func (AlgorithmName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1}
}

type AlgorithmSpecification struct {
	// The input mode can be either PIPE or FILE
	InputMode InputMode `protobuf:"varint,1,opt,name=input_mode,json=inputMode,proto3,enum=flyteidl.plugins.sagemaker.InputMode" json:"input_mode,omitempty"`
	// The algorithm name is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmName AlgorithmName `protobuf:"varint,2,opt,name=algorithm_name,json=algorithmName,proto3,enum=flyteidl.plugins.sagemaker.AlgorithmName" json:"algorithm_name,omitempty"`
	// The algorithm version field is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmVersion string `protobuf:"bytes,3,opt,name=algorithm_version,json=algorithmVersion,proto3" json:"algorithm_version,omitempty"`
	// A list of metric definitions for SageMaker to evaluate/track on the progress of the training job
	// https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics.html
	MetricDefinitions    []*AlgorithmSpecification_MetricDefinition `protobuf:"bytes,4,rep,name=metric_definitions,json=metricDefinitions,proto3" json:"metric_definitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AlgorithmSpecification) Reset()         { *m = AlgorithmSpecification{} }
func (m *AlgorithmSpecification) String() string { return proto.CompactTextString(m) }
func (*AlgorithmSpecification) ProtoMessage()    {}
func (*AlgorithmSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0}
}

func (m *AlgorithmSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmSpecification.Unmarshal(m, b)
}
func (m *AlgorithmSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmSpecification.Marshal(b, m, deterministic)
}
func (m *AlgorithmSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmSpecification.Merge(m, src)
}
func (m *AlgorithmSpecification) XXX_Size() int {
	return xxx_messageInfo_AlgorithmSpecification.Size(m)
}
func (m *AlgorithmSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmSpecification proto.InternalMessageInfo

func (m *AlgorithmSpecification) GetInputMode() InputMode {
	if m != nil {
		return m.InputMode
	}
	return InputMode_FILE
}

func (m *AlgorithmSpecification) GetAlgorithmName() AlgorithmName {
	if m != nil {
		return m.AlgorithmName
	}
	return AlgorithmName_CUSTOM
}

func (m *AlgorithmSpecification) GetAlgorithmVersion() string {
	if m != nil {
		return m.AlgorithmVersion
	}
	return ""
}

func (m *AlgorithmSpecification) GetMetricDefinitions() []*AlgorithmSpecification_MetricDefinition {
	if m != nil {
		return m.MetricDefinitions
	}
	return nil
}

type AlgorithmSpecification_MetricDefinition struct {
	// User-defined name of the metric
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SageMaker hyperparameter tuning parses your algorithm’s stdout and stderr streams to find algorithm metrics
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlgorithmSpecification_MetricDefinition) Reset() {
	*m = AlgorithmSpecification_MetricDefinition{}
}
func (m *AlgorithmSpecification_MetricDefinition) String() string { return proto.CompactTextString(m) }
func (*AlgorithmSpecification_MetricDefinition) ProtoMessage()    {}
func (*AlgorithmSpecification_MetricDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0, 0}
}

func (m *AlgorithmSpecification_MetricDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmSpecification_MetricDefinition.Unmarshal(m, b)
}
func (m *AlgorithmSpecification_MetricDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmSpecification_MetricDefinition.Marshal(b, m, deterministic)
}
func (m *AlgorithmSpecification_MetricDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmSpecification_MetricDefinition.Merge(m, src)
}
func (m *AlgorithmSpecification_MetricDefinition) XXX_Size() int {
	return xxx_messageInfo_AlgorithmSpecification_MetricDefinition.Size(m)
}
func (m *AlgorithmSpecification_MetricDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmSpecification_MetricDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmSpecification_MetricDefinition proto.InternalMessageInfo

func (m *AlgorithmSpecification_MetricDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlgorithmSpecification_MetricDefinition) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type TrainingJobConfig struct {
	// The number of ML compute instances to use. For distributed training, provide a value greater than 1.
	// This is for multi-node training, not multi-GPU training
	InstanceCount int64 `protobuf:"varint,1,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	// The ML compute instance type
	InstanceType string `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// The size of the ML storage volume that you want to provision.
	VolumeSizeInGb       int64    `protobuf:"varint,3,opt,name=volume_size_in_gb,json=volumeSizeInGb,proto3" json:"volume_size_in_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainingJobConfig) Reset()         { *m = TrainingJobConfig{} }
func (m *TrainingJobConfig) String() string { return proto.CompactTextString(m) }
func (*TrainingJobConfig) ProtoMessage()    {}
func (*TrainingJobConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1}
}

func (m *TrainingJobConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJobConfig.Unmarshal(m, b)
}
func (m *TrainingJobConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJobConfig.Marshal(b, m, deterministic)
}
func (m *TrainingJobConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJobConfig.Merge(m, src)
}
func (m *TrainingJobConfig) XXX_Size() int {
	return xxx_messageInfo_TrainingJobConfig.Size(m)
}
func (m *TrainingJobConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJobConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJobConfig proto.InternalMessageInfo

func (m *TrainingJobConfig) GetInstanceCount() int64 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

func (m *TrainingJobConfig) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *TrainingJobConfig) GetVolumeSizeInGb() int64 {
	if m != nil {
		return m.VolumeSizeInGb
	}
	return 0
}

// This option allows the users to specify a limit to how long a training job can run and
// how long the users are willing to wait for a managed spot training job to complete
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StoppingCondition.html
type StoppingCondition struct {
	// The maximum length of time in second that the training job can run.
	// If this value is not specified, the default expiration time will be 1 day
	MaxRuntimeInSeconds int64 `protobuf:"varint,1,opt,name=max_runtime_in_seconds,json=maxRuntimeInSeconds,proto3" json:"max_runtime_in_seconds,omitempty"`
	// The maximum length of time in seconds that the users are willing to wait for a managed spot
	// training job to complete.
	// Note that it is the amount of time spent waiting for Spot capacity plus the amount of time the
	// training job runs, so it must be equal to or greater than max_runtime_in_seconds.
	MaxWaitTimeInSeconds int64    `protobuf:"varint,2,opt,name=max_wait_time_in_seconds,json=maxWaitTimeInSeconds,proto3" json:"max_wait_time_in_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoppingCondition) Reset()         { *m = StoppingCondition{} }
func (m *StoppingCondition) String() string { return proto.CompactTextString(m) }
func (*StoppingCondition) ProtoMessage()    {}
func (*StoppingCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2}
}

func (m *StoppingCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoppingCondition.Unmarshal(m, b)
}
func (m *StoppingCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoppingCondition.Marshal(b, m, deterministic)
}
func (m *StoppingCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoppingCondition.Merge(m, src)
}
func (m *StoppingCondition) XXX_Size() int {
	return xxx_messageInfo_StoppingCondition.Size(m)
}
func (m *StoppingCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_StoppingCondition.DiscardUnknown(m)
}

var xxx_messageInfo_StoppingCondition proto.InternalMessageInfo

func (m *StoppingCondition) GetMaxRuntimeInSeconds() int64 {
	if m != nil {
		return m.MaxRuntimeInSeconds
	}
	return 0
}

func (m *StoppingCondition) GetMaxWaitTimeInSeconds() int64 {
	if m != nil {
		return m.MaxWaitTimeInSeconds
	}
	return 0
}

// The spec of a training job
type TrainingJob struct {
	AlgorithmSpecification *AlgorithmSpecification `protobuf:"bytes,1,opt,name=algorithm_specification,json=algorithmSpecification,proto3" json:"algorithm_specification,omitempty"`
	TrainingJobConfig      *TrainingJobConfig      `protobuf:"bytes,2,opt,name=training_job_config,json=trainingJobConfig,proto3" json:"training_job_config,omitempty"`
	Interruptible          bool                    `protobuf:"varint,3,opt,name=interruptible,proto3" json:"interruptible,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *TrainingJob) Reset()         { *m = TrainingJob{} }
func (m *TrainingJob) String() string { return proto.CompactTextString(m) }
func (*TrainingJob) ProtoMessage()    {}
func (*TrainingJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{3}
}

func (m *TrainingJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJob.Unmarshal(m, b)
}
func (m *TrainingJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJob.Marshal(b, m, deterministic)
}
func (m *TrainingJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJob.Merge(m, src)
}
func (m *TrainingJob) XXX_Size() int {
	return xxx_messageInfo_TrainingJob.Size(m)
}
func (m *TrainingJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJob.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJob proto.InternalMessageInfo

func (m *TrainingJob) GetAlgorithmSpecification() *AlgorithmSpecification {
	if m != nil {
		return m.AlgorithmSpecification
	}
	return nil
}

func (m *TrainingJob) GetTrainingJobConfig() *TrainingJobConfig {
	if m != nil {
		return m.TrainingJobConfig
	}
	return nil
}

func (m *TrainingJob) GetInterruptible() bool {
	if m != nil {
		return m.Interruptible
	}
	return false
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputMode", InputMode_name, InputMode_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.AlgorithmName", AlgorithmName_name, AlgorithmName_value)
	proto.RegisterType((*AlgorithmSpecification)(nil), "flyteidl.plugins.sagemaker.AlgorithmSpecification")
	proto.RegisterType((*AlgorithmSpecification_MetricDefinition)(nil), "flyteidl.plugins.sagemaker.AlgorithmSpecification.MetricDefinition")
	proto.RegisterType((*TrainingJobConfig)(nil), "flyteidl.plugins.sagemaker.TrainingJobConfig")
	proto.RegisterType((*StoppingCondition)(nil), "flyteidl.plugins.sagemaker.StoppingCondition")
	proto.RegisterType((*TrainingJob)(nil), "flyteidl.plugins.sagemaker.TrainingJob")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/training_job.proto", fileDescriptor_6a68f64d8fd9fe30)
}

var fileDescriptor_6a68f64d8fd9fe30 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x4f, 0xdb, 0x4c,
	0x10, 0xc5, 0x84, 0x8f, 0x8f, 0x4c, 0x4a, 0x94, 0x2c, 0x88, 0x46, 0x5c, 0x8a, 0xd2, 0x22, 0x05,
	0x2a, 0x6c, 0x29, 0xa8, 0x3d, 0xf5, 0x52, 0x02, 0x45, 0xa9, 0x4a, 0x41, 0x4e, 0xfa, 0x43, 0x95,
	0xaa, 0xd5, 0xda, 0xd9, 0x98, 0x29, 0xde, 0x5d, 0xcb, 0x5e, 0xd3, 0x84, 0x53, 0xcf, 0xbd, 0xf4,
	0x4f, 0x6e, 0xe5, 0x75, 0x62, 0x92, 0xd0, 0xa2, 0xf6, 0x36, 0xfb, 0x66, 0xe6, 0xed, 0xcc, 0x3c,
	0xcd, 0xc0, 0xc1, 0x30, 0x1c, 0x6b, 0x8e, 0x83, 0xd0, 0x89, 0xc2, 0x34, 0x40, 0x99, 0x38, 0x09,
	0x0b, 0xb8, 0x60, 0x57, 0x3c, 0x76, 0x74, 0xcc, 0x50, 0xa2, 0x0c, 0xe8, 0x17, 0xe5, 0xd9, 0x51,
	0xac, 0xb4, 0x22, 0xdb, 0xd3, 0x70, 0x7b, 0x12, 0x6e, 0x17, 0xe1, 0xcd, 0x1f, 0x25, 0xd8, 0x7a,
	0x19, 0x06, 0x2a, 0x46, 0x7d, 0x29, 0x7a, 0x11, 0xf7, 0x71, 0x88, 0x3e, 0xd3, 0xa8, 0x24, 0x39,
	0x06, 0x40, 0x19, 0xa5, 0x9a, 0x0a, 0x35, 0xe0, 0x0d, 0x6b, 0xc7, 0x6a, 0x55, 0xdb, 0xbb, 0xf6,
	0x9f, 0xb9, 0xec, 0x6e, 0x16, 0x7d, 0xa6, 0x06, 0xdc, 0x2d, 0xe3, 0xd4, 0x24, 0x17, 0x50, 0x65,
	0x53, 0x7e, 0x2a, 0x99, 0xe0, 0x8d, 0x65, 0xc3, 0xb4, 0x77, 0x1f, 0x53, 0x51, 0xd1, 0x5b, 0x26,
	0xb8, 0xbb, 0xce, 0x66, 0x9f, 0xe4, 0x29, 0xd4, 0x6f, 0x19, 0xaf, 0x79, 0x9c, 0xa0, 0x92, 0x8d,
	0xd2, 0x8e, 0xd5, 0x2a, 0xbb, 0xb5, 0xc2, 0xf1, 0x3e, 0xc7, 0x49, 0x0c, 0x44, 0x70, 0x1d, 0xa3,
	0x4f, 0x07, 0x7c, 0x88, 0x12, 0xb3, 0xce, 0x92, 0xc6, 0xca, 0x4e, 0xa9, 0x55, 0x69, 0x77, 0xfe,
	0xaa, 0x84, 0xb9, 0xa1, 0xd8, 0x67, 0x86, 0xec, 0xb8, 0xe0, 0x72, 0xeb, 0x62, 0x01, 0x49, 0xb6,
	0x5f, 0x40, 0x6d, 0x31, 0x8c, 0x10, 0x58, 0x31, 0xcd, 0x5b, 0xa6, 0x4e, 0x63, 0x93, 0x4d, 0xf8,
	0x2f, 0xe6, 0x01, 0x1f, 0x99, 0x89, 0x94, 0xdd, 0xfc, 0xd1, 0xfc, 0x6e, 0x41, 0xbd, 0x3f, 0x11,
	0xf1, 0xb5, 0xf2, 0x3a, 0x4a, 0x0e, 0x31, 0x20, 0xbb, 0x50, 0x45, 0x99, 0x68, 0x26, 0x7d, 0x4e,
	0x7d, 0x95, 0x4a, 0x6d, 0x98, 0x4a, 0xee, 0xfa, 0x14, 0xed, 0x64, 0x20, 0x79, 0x0c, 0x05, 0x40,
	0xf5, 0x38, 0xe2, 0x13, 0xea, 0x07, 0x53, 0xb0, 0x3f, 0x8e, 0x38, 0xd9, 0x83, 0xfa, 0xb5, 0x0a,
	0x53, 0xc1, 0x69, 0x82, 0x37, 0x9c, 0xa2, 0xa4, 0x81, 0x67, 0x06, 0x58, 0x72, 0xab, 0xb9, 0xa3,
	0x87, 0x37, 0xbc, 0x2b, 0x4f, 0xbd, 0xe6, 0x37, 0x0b, 0xea, 0x3d, 0xad, 0xa2, 0x08, 0x65, 0xd0,
	0x51, 0x72, 0x90, 0x37, 0x73, 0x08, 0x5b, 0x82, 0x8d, 0x68, 0x9c, 0x4a, 0x8d, 0xc2, 0x10, 0x24,
	0xdc, 0x57, 0x72, 0x90, 0x4c, 0x8a, 0xda, 0x10, 0x6c, 0xe4, 0xe6, 0xce, 0xae, 0xec, 0xe5, 0x2e,
	0xf2, 0x1c, 0x1a, 0x59, 0xd2, 0x57, 0x86, 0x9a, 0x2e, 0xa6, 0x2d, 0x9b, 0xb4, 0x4d, 0xc1, 0x46,
	0x1f, 0x18, 0xea, 0xfe, 0x6c, 0x5e, 0xf3, 0xa7, 0x05, 0x95, 0x99, 0x79, 0x90, 0x2b, 0x78, 0x78,
	0x2b, 0x7f, 0x32, 0x2b, 0x8e, 0xf9, 0xbd, 0xd2, 0x6e, 0xff, 0xbb, 0xac, 0xee, 0x16, 0xfb, 0xfd,
	0x0e, 0x7c, 0x86, 0x8d, 0xd9, 0x85, 0xa2, 0xbe, 0x51, 0xc3, 0xd4, 0x5b, 0x69, 0x1f, 0xdc, 0xf7,
	0xd1, 0x1d, 0x09, 0xdd, 0xba, 0xbe, 0xa3, 0xea, 0x93, 0x4c, 0x2e, 0xcd, 0xe3, 0x38, 0x8d, 0x34,
	0x7a, 0x21, 0x37, 0x2a, 0xac, 0xb9, 0xf3, 0xe0, 0xfe, 0x23, 0x28, 0x17, 0xab, 0x45, 0xd6, 0x60,
	0xe5, 0x55, 0xf7, 0xcd, 0x49, 0x6d, 0x29, 0xb3, 0x2e, 0xba, 0x17, 0x27, 0x35, 0x6b, 0xbf, 0x05,
	0xeb, 0x73, 0x1b, 0x43, 0x00, 0x56, 0x3b, 0xef, 0x7a, 0xfd, 0xf3, 0xb3, 0xda, 0x12, 0xa9, 0xc0,
	0xff, 0x1f, 0x4f, 0x8f, 0xce, 0xcf, 0x7b, 0xfd, 0x9a, 0x75, 0xf4, 0xec, 0xd3, 0x61, 0x80, 0xfa,
	0x32, 0xf5, 0x6c, 0x5f, 0x09, 0x27, 0x1c, 0x0f, 0xb5, 0x53, 0xdc, 0x92, 0x80, 0x4b, 0x27, 0xf2,
	0x0e, 0x02, 0xe5, 0x2c, 0x9e, 0x17, 0x6f, 0xd5, 0x1c, 0x92, 0xc3, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0xda, 0xd0, 0x9e, 0x79, 0x04, 0x00, 0x00,
}
