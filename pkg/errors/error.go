package errors

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

const (
	UnknownStatusCode = http.StatusInternalServerError
	UnknownCode       = "unknown"
)

type Error struct {
	// The error code is operational error definition
	Code string
	// Error information is user-readable information and can be used as user prompt content
	Message string
	// Error meta information, add additional extensible information for the error
	Metadata map[string]string
	// The Status Code like HTTP Status Code to categorized error
	StatusCode int
	// Detail error which is used to debug
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %s status_code = %d message = %s metadata = %v cause = %v", e.Code, e.StatusCode, e.Message, e.Metadata, e.cause)
}

// Unwrap provides compatibility for go error chains
func (e *Error) Unwrap() error {
	return e.cause
}

// Is matches each error in the chain with the target value
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code && se.StatusCode == se.StatusCode
	}
	return false
}

// WithCause with the underlying cause of the error.
func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]string) *Error {
	err := Clone(e)
	err.Metadata = md
	return err
}

// GRPCStatus returns the Status represented by se.
func (e *Error) GRPCStatus() *status.Status {
	s, _ := status.New(ToGRPCCode(e.StatusCode), e.Message).
		WithDetails(&errdetails.ErrorInfo{
			Reason:   e.Code,
			Metadata: e.Metadata,
		})
	return s
}

// GRPCError returns the grpc error
func (e *Error) GRPCError() error {
	return e.GRPCStatus().Err()
}

func (e *Error) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

func (e *Error) IsConflict() bool {
	return e.StatusCode == http.StatusConflict
}

func (e *Error) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

func (e *Error) IsForbidden() bool {
	return e.StatusCode == http.StatusForbidden
}

func (e *Error) IsInternalServer() bool {
	return e.StatusCode == http.StatusInternalServerError
}

func (e *Error) IsServiceUnavailable() bool {
	return e.StatusCode == http.StatusServiceUnavailable
}

func (e *Error) IsGatewayTimeout() bool {
	return e.StatusCode == http.StatusGatewayTimeout
}

func (e *Error) IsClientClosed() bool {
	return e.StatusCode == StatusClientClosed
}

func (e *Error) IsServerError() bool {
	return e.StatusCode >= http.StatusInternalServerError
}

func (e *Error) IsClientError() bool {
	return e.StatusCode >= http.StatusBadRequest && e.StatusCode < http.StatusInternalServerError
}

type Option func(*Error)

// New returns an error object for the code, message.
func New(statusCode int, opts ...Option) *Error {
	err := &Error{
		StatusCode: statusCode,
		Code:       ToCode(statusCode),
		Message:    ToMessage(statusCode),
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

// Newf New(code fmt.Sprintf(format, v...))
func Newf(statusCode int, code, format string, v ...interface{}) *Error {
	return New(statusCode, WithCode(code), WithMessage(format, v...))
}

// Errorf returns an error object for the code, message and error info.
func Errorf(statusCode int, code, format string, v ...interface{}) error {
	return New(statusCode, WithCode(code), WithMessage(format, v...))
}

// StatusCode returns the http code for an error.
// It supports wrapped errors.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	return FromError(err).StatusCode
}

// Code returns the reason for a particular error.
// It supports wrapped errors.
func Code(err error) string {
	if err == nil {
		return UnknownCode
	}
	return FromError(err).Code
}

func Clone(err *Error) *Error {
	if err == nil {
		return nil
	}
	metadata := make(map[string]string, len(err.Metadata))
	for k, v := range err.Metadata {
		metadata[k] = v
	}
	return &Error{
		cause:      err.cause,
		Code:       err.Code,
		StatusCode: err.StatusCode,
		Message:    err.Message,
		Metadata:   metadata,
	}
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if !ok {
		return New(UnknownStatusCode, WithCode(UnknownCode), WithMessage(err.Error()))
	}
	ret := New(FromGRPCCode(gs.Code()), WithCode(UnknownCode), WithMessage(gs.Message()))
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Code = d.Reason
			return ret.WithMetadata(d.Metadata)
		}
	}
	return ret
}

// WithCode with the underlying cause of the error.
func WithCode(code string) Option {
	return func(e *Error) {
		e.Code = code
	}
}

// WithMessage with the underlying cause of the error.
func WithMessage(format string, v ...interface{}) Option {
	return func(e *Error) {
		e.Message = fmt.Sprintf(format, v...)
	}
}

// WithCause with the underlying cause of the error.
func WithCause(cause error) Option {
	return func(e *Error) {
		e.cause = cause
	}
}

// WithMetadata with an MD formed by the mapping of key, value.
func WithMetadata(md map[string]string) Option {
	return func(e *Error) {
		e.Metadata = md
	}
}
