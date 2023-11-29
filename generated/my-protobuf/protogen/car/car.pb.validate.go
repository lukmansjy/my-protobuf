// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/car/car.proto

package car

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

// define the regex for a UUID once up-front
var _car_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Car with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Car) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Car with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CarMultiError, or nil if none found.
func (m *Car) ValidateAll() error {
	return m.validate(true)
}

func (m *Car) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = CarValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Car_Brand_Pattern.MatchString(m.GetBrand()) {
		err := CarValidationError{
			field:  "Brand",
			reason: "value does not match regex pattern \"(?i)^Toyota|Honda|Ford|BMW$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetModel()) > 30 {
		err := CarValidationError{
			field:  "Model",
			reason: "value length must be at most 30 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPrice() <= 0 {
		err := CarValidationError{
			field:  "Price",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetManufactureYear(); val < 2020 || val > 2030 {
		err := CarValidationError{
			field:  "ManufactureYear",
			reason: "value must be inside range [2020, 2030]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CarMultiError(errors)
	}

	return nil
}

func (m *Car) _validateUuid(uuid string) error {
	if matched := _car_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CarMultiError is an error wrapping multiple validation errors returned by
// Car.ValidateAll() if the designated constraints aren't met.
type CarMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarMultiError) AllErrors() []error { return m }

// CarValidationError is the validation error returned by Car.Validate if the
// designated constraints aren't met.
type CarValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarValidationError) ErrorName() string { return "CarValidationError" }

// Error satisfies the builtin error interface
func (e CarValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCar.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarValidationError{}

var _Car_Brand_Pattern = regexp.MustCompile("(?i)^Toyota|Honda|Ford|BMW$")