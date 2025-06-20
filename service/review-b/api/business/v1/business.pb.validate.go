// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: business/v1/business.proto

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

// Validate checks the field values on ReplyReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReplyReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReplyReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReplyReviewRequestMultiError, or nil if none found.
func (m *ReplyReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReplyReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetReviewId(); val <= 0 || val >= 9223372036854775807 {
		err := ReplyReviewRequestValidationError{
			field:  "ReviewId",
			reason: "value must be inside range (0, 9223372036854775807)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetStoreId(); val <= 0 || val >= 9223372036854775807 {
		err := ReplyReviewRequestValidationError{
			field:  "StoreId",
			reason: "value must be inside range (0, 9223372036854775807)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 1000 {
		err := ReplyReviewRequestValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPicInfo()); l < 1 || l > 1000 {
		err := ReplyReviewRequestValidationError{
			field:  "PicInfo",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetVideoInfo()); l < 1 || l > 1000 {
		err := ReplyReviewRequestValidationError{
			field:  "VideoInfo",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReplyReviewRequestMultiError(errors)
	}

	return nil
}

// ReplyReviewRequestMultiError is an error wrapping multiple validation errors
// returned by ReplyReviewRequest.ValidateAll() if the designated constraints
// aren't met.
type ReplyReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReplyReviewRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReplyReviewRequestMultiError) AllErrors() []error { return m }

// ReplyReviewRequestValidationError is the validation error returned by
// ReplyReviewRequest.Validate if the designated constraints aren't met.
type ReplyReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplyReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplyReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplyReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplyReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplyReviewRequestValidationError) ErrorName() string {
	return "ReplyReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReplyReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReplyReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplyReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplyReviewRequestValidationError{}

// Validate checks the field values on ReplyReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReplyReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReplyReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReplyReviewReplyMultiError, or nil if none found.
func (m *ReplyReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ReplyReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ReplyReviewReplyMultiError(errors)
	}

	return nil
}

// ReplyReviewReplyMultiError is an error wrapping multiple validation errors
// returned by ReplyReviewReply.ValidateAll() if the designated constraints
// aren't met.
type ReplyReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReplyReviewReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReplyReviewReplyMultiError) AllErrors() []error { return m }

// ReplyReviewReplyValidationError is the validation error returned by
// ReplyReviewReply.Validate if the designated constraints aren't met.
type ReplyReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplyReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplyReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplyReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplyReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplyReviewReplyValidationError) ErrorName() string { return "ReplyReviewReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReplyReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReplyReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplyReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplyReviewReplyValidationError{}
