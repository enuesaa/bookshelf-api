// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/board.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AddBoardRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddBoardRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddBoardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddBoardRequestMultiError, or nil if none found.
func (m *AddBoardRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddBoardRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddBoardRequestMultiError(errors)
	}

	return nil
}

// AddBoardRequestMultiError is an error wrapping multiple validation errors
// returned by AddBoardRequest.ValidateAll() if the designated constraints
// aren't met.
type AddBoardRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddBoardRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddBoardRequestMultiError) AllErrors() []error { return m }

// AddBoardRequestValidationError is the validation error returned by
// AddBoardRequest.Validate if the designated constraints aren't met.
type AddBoardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddBoardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddBoardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddBoardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddBoardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddBoardRequestValidationError) ErrorName() string { return "AddBoardRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddBoardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddBoardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddBoardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddBoardRequestValidationError{}

// Validate checks the field values on AddBoardResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddBoardResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddBoardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddBoardResponseMultiError, or nil if none found.
func (m *AddBoardResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddBoardResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddBoardResponseMultiError(errors)
	}

	return nil
}

// AddBoardResponseMultiError is an error wrapping multiple validation errors
// returned by AddBoardResponse.ValidateAll() if the designated constraints
// aren't met.
type AddBoardResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddBoardResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddBoardResponseMultiError) AllErrors() []error { return m }

// AddBoardResponseValidationError is the validation error returned by
// AddBoardResponse.Validate if the designated constraints aren't met.
type AddBoardResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddBoardResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddBoardResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddBoardResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddBoardResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddBoardResponseValidationError) ErrorName() string { return "AddBoardResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddBoardResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddBoardResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddBoardResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddBoardResponseValidationError{}

// Validate checks the field values on CheckinRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckinRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckinRequestMultiError,
// or nil if none found.
func (m *CheckinRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckinRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckinRequestValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckinRequestValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckinRequestValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CheckinRequestMultiError(errors)
	}

	return nil
}

// CheckinRequestMultiError is an error wrapping multiple validation errors
// returned by CheckinRequest.ValidateAll() if the designated constraints
// aren't met.
type CheckinRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckinRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckinRequestMultiError) AllErrors() []error { return m }

// CheckinRequestValidationError is the validation error returned by
// CheckinRequest.Validate if the designated constraints aren't met.
type CheckinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckinRequestValidationError) ErrorName() string { return "CheckinRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckinRequestValidationError{}

// Validate checks the field values on CheckinResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckinResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckinResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckinResponseMultiError, or nil if none found.
func (m *CheckinResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckinResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CheckinResponseMultiError(errors)
	}

	return nil
}

// CheckinResponseMultiError is an error wrapping multiple validation errors
// returned by CheckinResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckinResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckinResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckinResponseMultiError) AllErrors() []error { return m }

// CheckinResponseValidationError is the validation error returned by
// CheckinResponse.Validate if the designated constraints aren't met.
type CheckinResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckinResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckinResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckinResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckinResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckinResponseValidationError) ErrorName() string { return "CheckinResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckinResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckinResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckinResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckinResponseValidationError{}

// Validate checks the field values on ListTimelineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTimelineRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTimelineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTimelineRequestMultiError, or nil if none found.
func (m *ListTimelineRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTimelineRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	if len(errors) > 0 {
		return ListTimelineRequestMultiError(errors)
	}

	return nil
}

// ListTimelineRequestMultiError is an error wrapping multiple validation
// errors returned by ListTimelineRequest.ValidateAll() if the designated
// constraints aren't met.
type ListTimelineRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTimelineRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTimelineRequestMultiError) AllErrors() []error { return m }

// ListTimelineRequestValidationError is the validation error returned by
// ListTimelineRequest.Validate if the designated constraints aren't met.
type ListTimelineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTimelineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTimelineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTimelineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTimelineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTimelineRequestValidationError) ErrorName() string {
	return "ListTimelineRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTimelineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTimelineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTimelineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTimelineRequestValidationError{}

// Validate checks the field values on ListTimelineResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTimelineResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTimelineResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTimelineResponseMultiError, or nil if none found.
func (m *ListTimelineResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTimelineResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	if len(errors) > 0 {
		return ListTimelineResponseMultiError(errors)
	}

	return nil
}

// ListTimelineResponseMultiError is an error wrapping multiple validation
// errors returned by ListTimelineResponse.ValidateAll() if the designated
// constraints aren't met.
type ListTimelineResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTimelineResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTimelineResponseMultiError) AllErrors() []error { return m }

// ListTimelineResponseValidationError is the validation error returned by
// ListTimelineResponse.Validate if the designated constraints aren't met.
type ListTimelineResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTimelineResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTimelineResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTimelineResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTimelineResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTimelineResponseValidationError) ErrorName() string {
	return "ListTimelineResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListTimelineResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTimelineResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTimelineResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTimelineResponseValidationError{}

// Validate checks the field values on ArchiveBoardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArchiveBoardRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArchiveBoardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArchiveBoardRequestMultiError, or nil if none found.
func (m *ArchiveBoardRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ArchiveBoardRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ArchiveBoardRequestMultiError(errors)
	}

	return nil
}

// ArchiveBoardRequestMultiError is an error wrapping multiple validation
// errors returned by ArchiveBoardRequest.ValidateAll() if the designated
// constraints aren't met.
type ArchiveBoardRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArchiveBoardRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArchiveBoardRequestMultiError) AllErrors() []error { return m }

// ArchiveBoardRequestValidationError is the validation error returned by
// ArchiveBoardRequest.Validate if the designated constraints aren't met.
type ArchiveBoardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArchiveBoardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArchiveBoardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArchiveBoardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArchiveBoardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArchiveBoardRequestValidationError) ErrorName() string {
	return "ArchiveBoardRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ArchiveBoardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArchiveBoardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArchiveBoardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArchiveBoardRequestValidationError{}

// Validate checks the field values on ArchiveBoardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArchiveBoardResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArchiveBoardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArchiveBoardResponseMultiError, or nil if none found.
func (m *ArchiveBoardResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ArchiveBoardResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ArchiveBoardResponseMultiError(errors)
	}

	return nil
}

// ArchiveBoardResponseMultiError is an error wrapping multiple validation
// errors returned by ArchiveBoardResponse.ValidateAll() if the designated
// constraints aren't met.
type ArchiveBoardResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArchiveBoardResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArchiveBoardResponseMultiError) AllErrors() []error { return m }

// ArchiveBoardResponseValidationError is the validation error returned by
// ArchiveBoardResponse.Validate if the designated constraints aren't met.
type ArchiveBoardResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArchiveBoardResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArchiveBoardResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArchiveBoardResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArchiveBoardResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArchiveBoardResponseValidationError) ErrorName() string {
	return "ArchiveBoardResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ArchiveBoardResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArchiveBoardResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArchiveBoardResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArchiveBoardResponseValidationError{}

// Validate checks the field values on UnArchiveBoardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnArchiveBoardRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnArchiveBoardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnArchiveBoardRequestMultiError, or nil if none found.
func (m *UnArchiveBoardRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnArchiveBoardRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UnArchiveBoardRequestMultiError(errors)
	}

	return nil
}

// UnArchiveBoardRequestMultiError is an error wrapping multiple validation
// errors returned by UnArchiveBoardRequest.ValidateAll() if the designated
// constraints aren't met.
type UnArchiveBoardRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnArchiveBoardRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnArchiveBoardRequestMultiError) AllErrors() []error { return m }

// UnArchiveBoardRequestValidationError is the validation error returned by
// UnArchiveBoardRequest.Validate if the designated constraints aren't met.
type UnArchiveBoardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnArchiveBoardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnArchiveBoardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnArchiveBoardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnArchiveBoardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnArchiveBoardRequestValidationError) ErrorName() string {
	return "UnArchiveBoardRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UnArchiveBoardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnArchiveBoardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnArchiveBoardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnArchiveBoardRequestValidationError{}

// Validate checks the field values on UnArchiveBoardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnArchiveBoardResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnArchiveBoardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnArchiveBoardResponseMultiError, or nil if none found.
func (m *UnArchiveBoardResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UnArchiveBoardResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnArchiveBoardResponseMultiError(errors)
	}

	return nil
}

// UnArchiveBoardResponseMultiError is an error wrapping multiple validation
// errors returned by UnArchiveBoardResponse.ValidateAll() if the designated
// constraints aren't met.
type UnArchiveBoardResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnArchiveBoardResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnArchiveBoardResponseMultiError) AllErrors() []error { return m }

// UnArchiveBoardResponseValidationError is the validation error returned by
// UnArchiveBoardResponse.Validate if the designated constraints aren't met.
type UnArchiveBoardResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnArchiveBoardResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnArchiveBoardResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnArchiveBoardResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnArchiveBoardResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnArchiveBoardResponseValidationError) ErrorName() string {
	return "UnArchiveBoardResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UnArchiveBoardResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnArchiveBoardResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnArchiveBoardResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnArchiveBoardResponseValidationError{}