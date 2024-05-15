package errutil

type notFoundErr interface {
	IsNotFound() bool
}

type serverErr interface {
	IsServerError() bool
}

type clientErr interface {
	IsClientError() bool
}

type conflictErr interface {
	IsConflict() bool
}

type unauthorizedErr interface {
	IsUnauthorized() bool
}

type forbiddenErr interface {
	IsForbidden() bool
}

type internalErr interface {
	IsInternalServer() bool
}

type serviceUnavailableErr interface {
	IsServiceUnavailable() bool
}

type gatewayTimeoutErr interface {
	IsGatewayTimeout() bool
}

type clientClosedErr interface {
	IsClientClosed() bool
}

func IsNotFound(err error) bool {
	te, ok := err.(notFoundErr)
	return ok && te.IsNotFound()
}

func IsConflict(err error) bool {
	te, ok := err.(conflictErr)
	return ok && te.IsConflict()
}

func IsUnauthorized(err error) bool {
	te, ok := err.(unauthorizedErr)
	return ok && te.IsUnauthorized()
}

func IsForbidden(err error) bool {
	te, ok := err.(forbiddenErr)
	return ok && te.IsForbidden()
}

func IsInternalServer(err error) bool {
	te, ok := err.(internalErr)
	return ok && te.IsInternalServer()
}

func IsServiceUnavailable(err error) bool {
	te, ok := err.(serviceUnavailableErr)
	return ok && te.IsServiceUnavailable()
}

func IsGatewayTimeout(err error) bool {
	te, ok := err.(gatewayTimeoutErr)
	return ok && te.IsGatewayTimeout()
}

func IsClientClosed(err error) bool {
	te, ok := err.(clientClosedErr)
	return ok && te.IsClientClosed()
}

func IsClientErr(err error) bool {
	te, ok := err.(clientErr)
	return ok && te.IsClientError()
}

func IsServerError(err error) bool {
	te, ok := err.(serverErr)
	return ok && te.IsServerError()
}
