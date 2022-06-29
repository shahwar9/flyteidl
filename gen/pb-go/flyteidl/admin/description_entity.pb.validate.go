// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/description_entity.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _description_entity_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DescriptionEntity with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DescriptionEntity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ShortDescription

	if v, ok := interface{}(m.GetLongDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescriptionEntityValidationError{
				field:  "LongDescription",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescriptionEntityValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSourceCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescriptionEntityValidationError{
				field:  "SourceCode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescriptionEntityValidationError is the validation error returned by
// DescriptionEntity.Validate if the designated constraints aren't met.
type DescriptionEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescriptionEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescriptionEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescriptionEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescriptionEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescriptionEntityValidationError) ErrorName() string {
	return "DescriptionEntityValidationError"
}

// Error satisfies the builtin error interface
func (e DescriptionEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescriptionEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescriptionEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescriptionEntityValidationError{}

// Validate checks the field values on LongDescription with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LongDescription) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Values

	// no validation rules for Uri

	// no validation rules for LongFormat

	// no validation rules for IconLink

	return nil
}

// LongDescriptionValidationError is the validation error returned by
// LongDescription.Validate if the designated constraints aren't met.
type LongDescriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LongDescriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LongDescriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LongDescriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LongDescriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LongDescriptionValidationError) ErrorName() string { return "LongDescriptionValidationError" }

// Error satisfies the builtin error interface
func (e LongDescriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLongDescription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LongDescriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LongDescriptionValidationError{}

// Validate checks the field values on SourceCode with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SourceCode) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Link

	return nil
}

// SourceCodeValidationError is the validation error returned by
// SourceCode.Validate if the designated constraints aren't met.
type SourceCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceCodeValidationError) ErrorName() string { return "SourceCodeValidationError" }

// Error satisfies the builtin error interface
func (e SourceCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceCodeValidationError{}

// Validate checks the field values on DescriptionEntityCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescriptionEntityCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescriptionEntityCreateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescriptionEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescriptionEntityCreateRequestValidationError{
				field:  "DescriptionEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescriptionEntityCreateRequestValidationError is the validation error
// returned by DescriptionEntityCreateRequest.Validate if the designated
// constraints aren't met.
type DescriptionEntityCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescriptionEntityCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescriptionEntityCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescriptionEntityCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescriptionEntityCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescriptionEntityCreateRequestValidationError) ErrorName() string {
	return "DescriptionEntityCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescriptionEntityCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescriptionEntityCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescriptionEntityCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescriptionEntityCreateRequestValidationError{}

// Validate checks the field values on DescriptionEntityCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescriptionEntityCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DescriptionEntityCreateResponseValidationError is the validation error
// returned by DescriptionEntityCreateResponse.Validate if the designated
// constraints aren't met.
type DescriptionEntityCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescriptionEntityCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescriptionEntityCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescriptionEntityCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescriptionEntityCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescriptionEntityCreateResponseValidationError) ErrorName() string {
	return "DescriptionEntityCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescriptionEntityCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescriptionEntityCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescriptionEntityCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescriptionEntityCreateResponseValidationError{}
