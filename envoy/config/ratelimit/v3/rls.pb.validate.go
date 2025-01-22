//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/ratelimit/v3/rls.proto

package ratelimitv3

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

	v3 "github.com/sefaphlvn/versioned-go-control-plane/envoy/config/core/v3"
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

	_ = v3.ApiVersion(0)
)

// Validate checks the field values on RateLimitServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitServiceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitServiceConfigMultiError, or nil if none found.
func (m *RateLimitServiceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitServiceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetGrpcService() == nil {
		err := RateLimitServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitServiceConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := v3.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := RateLimitServiceConfigValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RateLimitServiceConfigMultiError(errors)
	}

	return nil
}

// RateLimitServiceConfigMultiError is an error wrapping multiple validation
// errors returned by RateLimitServiceConfig.ValidateAll() if the designated
// constraints aren't met.
type RateLimitServiceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitServiceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitServiceConfigMultiError) AllErrors() []error { return m }

// RateLimitServiceConfigValidationError is the validation error returned by
// RateLimitServiceConfig.Validate if the designated constraints aren't met.
type RateLimitServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitServiceConfigValidationError) ErrorName() string {
	return "RateLimitServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitServiceConfigValidationError{}
