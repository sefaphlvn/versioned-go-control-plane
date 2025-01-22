//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/compression/qatzip/compressor/v3alpha/qatzip.proto

package v3alpha

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

// Validate checks the field values on Qatzip with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Qatzip) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Qatzip with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in QatzipMultiError, or nil if none found.
func (m *Qatzip) ValidateAll() error {
	return m.validate(true)
}

func (m *Qatzip) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetCompressionLevel(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 9 {
			err := QatzipValidationError{
				field:  "CompressionLevel",
				reason: "value must be inside range [1, 9]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if _, ok := Qatzip_HardwareBufferSize_name[int32(m.GetHardwareBufferSize())]; !ok {
		err := QatzipValidationError{
			field:  "HardwareBufferSize",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if wrapper := m.GetInputSizeThreshold(); wrapper != nil {

		if val := wrapper.GetValue(); val < 128 || val > 524288 {
			err := QatzipValidationError{
				field:  "InputSizeThreshold",
				reason: "value must be inside range [128, 524288]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetStreamBufferSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1024 || val > 2092032 {
			err := QatzipValidationError{
				field:  "StreamBufferSize",
				reason: "value must be inside range [1024, 2092032]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetChunkSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 4096 || val > 65536 {
			err := QatzipValidationError{
				field:  "ChunkSize",
				reason: "value must be inside range [4096, 65536]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return QatzipMultiError(errors)
	}

	return nil
}

// QatzipMultiError is an error wrapping multiple validation errors returned by
// Qatzip.ValidateAll() if the designated constraints aren't met.
type QatzipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QatzipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QatzipMultiError) AllErrors() []error { return m }

// QatzipValidationError is the validation error returned by Qatzip.Validate if
// the designated constraints aren't met.
type QatzipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QatzipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QatzipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QatzipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QatzipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QatzipValidationError) ErrorName() string { return "QatzipValidationError" }

// Error satisfies the builtin error interface
func (e QatzipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQatzip.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QatzipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QatzipValidationError{}
