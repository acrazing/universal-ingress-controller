// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: universal_ingress_controller/envoy_subscriber/v1/config.proto

package envoyv1

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

// Validate checks the field values on EnvoyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EnvoyConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnvoyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnvoyConfigMultiError, or
// nil if none found.
func (m *EnvoyConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *EnvoyConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	_EnvoyConfig_TargetListeners_Unique := make(map[string]struct{}, len(m.GetTargetListeners()))

	for idx, item := range m.GetTargetListeners() {
		_, _ = idx, item

		if _, exists := _EnvoyConfig_TargetListeners_Unique[item]; exists {
			err := EnvoyConfigValidationError{
				field:  fmt.Sprintf("TargetListeners[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_EnvoyConfig_TargetListeners_Unique[item] = struct{}{}
		}

		if _, ok := _EnvoyConfig_TargetListeners_NotInLookup[item]; ok {
			err := EnvoyConfigValidationError{
				field:  fmt.Sprintf("TargetListeners[%v]", idx),
				reason: "value must not be in list []",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetClusters()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Clusters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Clusters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClusters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Clusters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetListeners()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Listeners",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Listeners",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetListeners()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Listeners",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntimes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Runtimes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Runtimes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Runtimes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRoutes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Routes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Routes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoutes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Routes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetScopedRoutes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "ScopedRoutes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "ScopedRoutes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetScopedRoutes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "ScopedRoutes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHosts()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Hosts",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Hosts",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHosts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Hosts",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndpoints()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Endpoints",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Endpoints",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpoints()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Endpoints",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSecrets()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Secrets",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvoyConfigValidationError{
					field:  "Secrets",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSecrets()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvoyConfigValidationError{
				field:  "Secrets",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EnvoyConfigMultiError(errors)
	}

	return nil
}

// EnvoyConfigMultiError is an error wrapping multiple validation errors
// returned by EnvoyConfig.ValidateAll() if the designated constraints aren't met.
type EnvoyConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvoyConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvoyConfigMultiError) AllErrors() []error { return m }

// EnvoyConfigValidationError is the validation error returned by
// EnvoyConfig.Validate if the designated constraints aren't met.
type EnvoyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvoyConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvoyConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvoyConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvoyConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvoyConfigValidationError) ErrorName() string { return "EnvoyConfigValidationError" }

// Error satisfies the builtin error interface
func (e EnvoyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvoyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvoyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvoyConfigValidationError{}

var _EnvoyConfig_TargetListeners_NotInLookup = map[string]struct{}{
	"": {},
}

// Validate checks the field values on EnvoyConfig_DiscoverySource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EnvoyConfig_DiscoverySource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnvoyConfig_DiscoverySource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EnvoyConfig_DiscoverySourceMultiError, or nil if none found.
func (m *EnvoyConfig_DiscoverySource) ValidateAll() error {
	return m.validate(true)
}

func (m *EnvoyConfig_DiscoverySource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Source.(type) {

	case *EnvoyConfig_DiscoverySource_Inline:

		if all {
			switch v := interface{}(m.GetInline()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnvoyConfig_DiscoverySourceValidationError{
						field:  "Inline",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnvoyConfig_DiscoverySourceValidationError{
						field:  "Inline",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInline()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnvoyConfig_DiscoverySourceValidationError{
					field:  "Inline",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *EnvoyConfig_DiscoverySource_Kube:

		if all {
			switch v := interface{}(m.GetKube()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnvoyConfig_DiscoverySourceValidationError{
						field:  "Kube",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnvoyConfig_DiscoverySourceValidationError{
						field:  "Kube",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKube()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnvoyConfig_DiscoverySourceValidationError{
					field:  "Kube",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EnvoyConfig_DiscoverySourceMultiError(errors)
	}

	return nil
}

// EnvoyConfig_DiscoverySourceMultiError is an error wrapping multiple
// validation errors returned by EnvoyConfig_DiscoverySource.ValidateAll() if
// the designated constraints aren't met.
type EnvoyConfig_DiscoverySourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvoyConfig_DiscoverySourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvoyConfig_DiscoverySourceMultiError) AllErrors() []error { return m }

// EnvoyConfig_DiscoverySourceValidationError is the validation error returned
// by EnvoyConfig_DiscoverySource.Validate if the designated constraints
// aren't met.
type EnvoyConfig_DiscoverySourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvoyConfig_DiscoverySourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvoyConfig_DiscoverySourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvoyConfig_DiscoverySourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvoyConfig_DiscoverySourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvoyConfig_DiscoverySourceValidationError) ErrorName() string {
	return "EnvoyConfig_DiscoverySourceValidationError"
}

// Error satisfies the builtin error interface
func (e EnvoyConfig_DiscoverySourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvoyConfig_DiscoverySource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvoyConfig_DiscoverySourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvoyConfig_DiscoverySourceValidationError{}