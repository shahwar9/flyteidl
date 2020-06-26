// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/admin/quality_of_service.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto();
namespace flyteidl {
namespace admin {
class QualityOfService;
class QualityOfServiceDefaultTypeInternal;
extern QualityOfServiceDefaultTypeInternal _QualityOfService_default_instance_;
class QualityOfServiceSpec;
class QualityOfServiceSpecDefaultTypeInternal;
extern QualityOfServiceSpecDefaultTypeInternal _QualityOfServiceSpec_default_instance_;
}  // namespace admin
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::admin::QualityOfService* Arena::CreateMaybeMessage<::flyteidl::admin::QualityOfService>(Arena*);
template<> ::flyteidl::admin::QualityOfServiceSpec* Arena::CreateMaybeMessage<::flyteidl::admin::QualityOfServiceSpec>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace admin {

enum QualityOfService_Tier {
  QualityOfService_Tier_UNDEFINED = 0,
  QualityOfService_Tier_HIGH = 1,
  QualityOfService_Tier_MEDIUM = 2,
  QualityOfService_Tier_LOW = 3,
  QualityOfService_Tier_QualityOfService_Tier_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  QualityOfService_Tier_QualityOfService_Tier_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool QualityOfService_Tier_IsValid(int value);
const QualityOfService_Tier QualityOfService_Tier_Tier_MIN = QualityOfService_Tier_UNDEFINED;
const QualityOfService_Tier QualityOfService_Tier_Tier_MAX = QualityOfService_Tier_LOW;
const int QualityOfService_Tier_Tier_ARRAYSIZE = QualityOfService_Tier_Tier_MAX + 1;

const ::google::protobuf::EnumDescriptor* QualityOfService_Tier_descriptor();
inline const ::std::string& QualityOfService_Tier_Name(QualityOfService_Tier value) {
  return ::google::protobuf::internal::NameOfEnum(
    QualityOfService_Tier_descriptor(), value);
}
inline bool QualityOfService_Tier_Parse(
    const ::std::string& name, QualityOfService_Tier* value) {
  return ::google::protobuf::internal::ParseNamedEnum<QualityOfService_Tier>(
    QualityOfService_Tier_descriptor(), name, value);
}
// ===================================================================

class QualityOfServiceSpec final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.admin.QualityOfServiceSpec) */ {
 public:
  QualityOfServiceSpec();
  virtual ~QualityOfServiceSpec();

  QualityOfServiceSpec(const QualityOfServiceSpec& from);

  inline QualityOfServiceSpec& operator=(const QualityOfServiceSpec& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  QualityOfServiceSpec(QualityOfServiceSpec&& from) noexcept
    : QualityOfServiceSpec() {
    *this = ::std::move(from);
  }

  inline QualityOfServiceSpec& operator=(QualityOfServiceSpec&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const QualityOfServiceSpec& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const QualityOfServiceSpec* internal_default_instance() {
    return reinterpret_cast<const QualityOfServiceSpec*>(
               &_QualityOfServiceSpec_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(QualityOfServiceSpec* other);
  friend void swap(QualityOfServiceSpec& a, QualityOfServiceSpec& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline QualityOfServiceSpec* New() const final {
    return CreateMaybeMessage<QualityOfServiceSpec>(nullptr);
  }

  QualityOfServiceSpec* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<QualityOfServiceSpec>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const QualityOfServiceSpec& from);
  void MergeFrom(const QualityOfServiceSpec& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(QualityOfServiceSpec* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // uint32 queueing_budget_mins = 1;
  void clear_queueing_budget_mins();
  static const int kQueueingBudgetMinsFieldNumber = 1;
  ::google::protobuf::uint32 queueing_budget_mins() const;
  void set_queueing_budget_mins(::google::protobuf::uint32 value);

  // @@protoc_insertion_point(class_scope:flyteidl.admin.QualityOfServiceSpec)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::uint32 queueing_budget_mins_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto;
};
// -------------------------------------------------------------------

class QualityOfService final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.admin.QualityOfService) */ {
 public:
  QualityOfService();
  virtual ~QualityOfService();

  QualityOfService(const QualityOfService& from);

  inline QualityOfService& operator=(const QualityOfService& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  QualityOfService(QualityOfService&& from) noexcept
    : QualityOfService() {
    *this = ::std::move(from);
  }

  inline QualityOfService& operator=(QualityOfService&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const QualityOfService& default_instance();

  enum DesignationCase {
    kTier = 1,
    kSpec = 2,
    DESIGNATION_NOT_SET = 0,
  };

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const QualityOfService* internal_default_instance() {
    return reinterpret_cast<const QualityOfService*>(
               &_QualityOfService_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(QualityOfService* other);
  friend void swap(QualityOfService& a, QualityOfService& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline QualityOfService* New() const final {
    return CreateMaybeMessage<QualityOfService>(nullptr);
  }

  QualityOfService* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<QualityOfService>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const QualityOfService& from);
  void MergeFrom(const QualityOfService& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(QualityOfService* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  typedef QualityOfService_Tier Tier;
  static const Tier UNDEFINED =
    QualityOfService_Tier_UNDEFINED;
  static const Tier HIGH =
    QualityOfService_Tier_HIGH;
  static const Tier MEDIUM =
    QualityOfService_Tier_MEDIUM;
  static const Tier LOW =
    QualityOfService_Tier_LOW;
  static inline bool Tier_IsValid(int value) {
    return QualityOfService_Tier_IsValid(value);
  }
  static const Tier Tier_MIN =
    QualityOfService_Tier_Tier_MIN;
  static const Tier Tier_MAX =
    QualityOfService_Tier_Tier_MAX;
  static const int Tier_ARRAYSIZE =
    QualityOfService_Tier_Tier_ARRAYSIZE;
  static inline const ::google::protobuf::EnumDescriptor*
  Tier_descriptor() {
    return QualityOfService_Tier_descriptor();
  }
  static inline const ::std::string& Tier_Name(Tier value) {
    return QualityOfService_Tier_Name(value);
  }
  static inline bool Tier_Parse(const ::std::string& name,
      Tier* value) {
    return QualityOfService_Tier_Parse(name, value);
  }

  // accessors -------------------------------------------------------

  // .flyteidl.admin.QualityOfService.Tier tier = 1;
  private:
  bool has_tier() const;
  public:
  void clear_tier();
  static const int kTierFieldNumber = 1;
  ::flyteidl::admin::QualityOfService_Tier tier() const;
  void set_tier(::flyteidl::admin::QualityOfService_Tier value);

  // .flyteidl.admin.QualityOfServiceSpec spec = 2;
  bool has_spec() const;
  void clear_spec();
  static const int kSpecFieldNumber = 2;
  const ::flyteidl::admin::QualityOfServiceSpec& spec() const;
  ::flyteidl::admin::QualityOfServiceSpec* release_spec();
  ::flyteidl::admin::QualityOfServiceSpec* mutable_spec();
  void set_allocated_spec(::flyteidl::admin::QualityOfServiceSpec* spec);

  void clear_designation();
  DesignationCase designation_case() const;
  // @@protoc_insertion_point(class_scope:flyteidl.admin.QualityOfService)
 private:
  class HasBitSetters;
  void set_has_tier();
  void set_has_spec();

  inline bool has_designation() const;
  inline void clear_has_designation();

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  union DesignationUnion {
    DesignationUnion() {}
    int tier_;
    ::flyteidl::admin::QualityOfServiceSpec* spec_;
  } designation_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  ::google::protobuf::uint32 _oneof_case_[1];

  friend struct ::TableStruct_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// QualityOfServiceSpec

// uint32 queueing_budget_mins = 1;
inline void QualityOfServiceSpec::clear_queueing_budget_mins() {
  queueing_budget_mins_ = 0u;
}
inline ::google::protobuf::uint32 QualityOfServiceSpec::queueing_budget_mins() const {
  // @@protoc_insertion_point(field_get:flyteidl.admin.QualityOfServiceSpec.queueing_budget_mins)
  return queueing_budget_mins_;
}
inline void QualityOfServiceSpec::set_queueing_budget_mins(::google::protobuf::uint32 value) {
  
  queueing_budget_mins_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.admin.QualityOfServiceSpec.queueing_budget_mins)
}

// -------------------------------------------------------------------

// QualityOfService

// .flyteidl.admin.QualityOfService.Tier tier = 1;
inline bool QualityOfService::has_tier() const {
  return designation_case() == kTier;
}
inline void QualityOfService::set_has_tier() {
  _oneof_case_[0] = kTier;
}
inline void QualityOfService::clear_tier() {
  if (has_tier()) {
    designation_.tier_ = 0;
    clear_has_designation();
  }
}
inline ::flyteidl::admin::QualityOfService_Tier QualityOfService::tier() const {
  // @@protoc_insertion_point(field_get:flyteidl.admin.QualityOfService.tier)
  if (has_tier()) {
    return static_cast< ::flyteidl::admin::QualityOfService_Tier >(designation_.tier_);
  }
  return static_cast< ::flyteidl::admin::QualityOfService_Tier >(0);
}
inline void QualityOfService::set_tier(::flyteidl::admin::QualityOfService_Tier value) {
  if (!has_tier()) {
    clear_designation();
    set_has_tier();
  }
  designation_.tier_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.admin.QualityOfService.tier)
}

// .flyteidl.admin.QualityOfServiceSpec spec = 2;
inline bool QualityOfService::has_spec() const {
  return designation_case() == kSpec;
}
inline void QualityOfService::set_has_spec() {
  _oneof_case_[0] = kSpec;
}
inline void QualityOfService::clear_spec() {
  if (has_spec()) {
    delete designation_.spec_;
    clear_has_designation();
  }
}
inline ::flyteidl::admin::QualityOfServiceSpec* QualityOfService::release_spec() {
  // @@protoc_insertion_point(field_release:flyteidl.admin.QualityOfService.spec)
  if (has_spec()) {
    clear_has_designation();
      ::flyteidl::admin::QualityOfServiceSpec* temp = designation_.spec_;
    designation_.spec_ = nullptr;
    return temp;
  } else {
    return nullptr;
  }
}
inline const ::flyteidl::admin::QualityOfServiceSpec& QualityOfService::spec() const {
  // @@protoc_insertion_point(field_get:flyteidl.admin.QualityOfService.spec)
  return has_spec()
      ? *designation_.spec_
      : *reinterpret_cast< ::flyteidl::admin::QualityOfServiceSpec*>(&::flyteidl::admin::_QualityOfServiceSpec_default_instance_);
}
inline ::flyteidl::admin::QualityOfServiceSpec* QualityOfService::mutable_spec() {
  if (!has_spec()) {
    clear_designation();
    set_has_spec();
    designation_.spec_ = CreateMaybeMessage< ::flyteidl::admin::QualityOfServiceSpec >(
        GetArenaNoVirtual());
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.admin.QualityOfService.spec)
  return designation_.spec_;
}

inline bool QualityOfService::has_designation() const {
  return designation_case() != DESIGNATION_NOT_SET;
}
inline void QualityOfService::clear_has_designation() {
  _oneof_case_[0] = DESIGNATION_NOT_SET;
}
inline QualityOfService::DesignationCase QualityOfService::designation_case() const {
  return QualityOfService::DesignationCase(_oneof_case_[0]);
}
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace admin
}  // namespace flyteidl

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::flyteidl::admin::QualityOfService_Tier> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::admin::QualityOfService_Tier>() {
  return ::flyteidl::admin::QualityOfService_Tier_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fadmin_2fquality_5fof_5fservice_2eproto
