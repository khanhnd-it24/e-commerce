package errors

import "net/http"

func BadRequest(options ...Option) *Error {
	return New(http.StatusBadRequest, options...)
}

func Unauthorized(options ...Option) *Error {
	return New(http.StatusUnauthorized, options...)
}

func Forbidden(options ...Option) *Error {
	return New(http.StatusForbidden, options...)
}

func NotFound(options ...Option) *Error {
	return New(http.StatusNotFound, options...)
}

func Conflict(options ...Option) *Error {
	return New(http.StatusConflict, options...)
}

func InternalServer(options ...Option) *Error {
	return New(http.StatusInternalServerError, options...)
}

func ServiceUnavailable(options ...Option) *Error {
	return New(http.StatusServiceUnavailable, options...)
}

func GatewayTimeout(options ...Option) *Error {
	return New(http.StatusGatewayTimeout, options...)
}

func ClientClosed(options ...Option) *Error {
	return New(StatusClientClosed, options...)
}
