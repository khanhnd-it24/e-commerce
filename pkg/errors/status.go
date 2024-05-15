package errors

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

const (
	// StatusClientClosed is non-standard http status code,
	// which defined by nginx.
	// https://httpstatus.in/499/
	StatusClientClosed = 499
)

func ToGRPCCode(code int) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.Aborted
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case StatusClientClosed:
		return codes.Canceled
	}
	return codes.Unknown
}

func ToMessage(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "Success"
	case http.StatusBadRequest:
		return "BadRequest"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "NotFound"
	case http.StatusConflict:
		return "Conflict"
	case http.StatusTooManyRequests:
		return "TooManyRequests"
	case http.StatusInternalServerError:
		return "InternalServerError"
	case http.StatusNotImplemented:
		return "NotImplemented"
	case http.StatusServiceUnavailable:
		return "ServiceUnavailable"
	case http.StatusGatewayTimeout:
		return "GatewayTimeout"
	case StatusClientClosed:
		return "ClientClosed"
	}
	return "Unknown"
}

func FromGRPCCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return StatusClientClosed
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

func ToCode(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "success"
	case http.StatusBadRequest:
		return "bad_request"
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not_found"
	case http.StatusConflict:
		return "conflict"
	case http.StatusTooManyRequests:
		return "too_many_requests"
	case http.StatusInternalServerError:
		return "internal_server_error"
	case http.StatusNotImplemented:
		return "not_implemented"
	case http.StatusServiceUnavailable:
		return "service_unavailable"
	case http.StatusGatewayTimeout:
		return "gateway_timeout"
	case StatusClientClosed:
		return "client_closed"
	}
	return "unknown"
}
