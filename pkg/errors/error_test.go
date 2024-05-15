package errors_test

import (
	stderrors "errors"
	"fmt"
	"gobase/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"reflect"
	"testing"
)

type TestError struct{ message string }

func (e *TestError) Error() string { return e.message }

func TestErrors(t *testing.T) {
	var base *errors.Error
	err := errors.Newf(http.StatusBadRequest, "code", "message")
	err2 := errors.Newf(http.StatusBadRequest, "code", "message")
	err3 := err.WithMetadata(map[string]string{
		"foo": "bar",
	})
	wErr := fmt.Errorf("wrap %w", err)

	if stderrors.Is(err, new(errors.Error)) {
		t.Errorf("should not be equal: %v", err)
	}
	if !stderrors.Is(wErr, err) {
		t.Errorf("should be equal: %v", err)
	}
	if !stderrors.Is(wErr, err2) {
		t.Errorf("should be equal: %v", err2)
	}

	if !stderrors.As(err, &base) {
		t.Errorf("should be matches: %v", err)
	}

	if code := err.Code; code != err3.Code {
		t.Errorf("got %s want: %s", code, err)
	}

	if err3.Metadata["foo"] != "bar" {
		t.Error("not expected metadata")
	}

	gs := err.GRPCStatus()
	se := errors.FromError(gs.Err())
	if se.Code != "code" {
		t.Errorf("got %+v want %+v", se, err)
	}

	gs2 := status.New(codes.InvalidArgument, "bad request")
	se2 := errors.FromError(gs2.Err())
	// codes.InvalidArgument should convert to http.StatusBadRequest
	if se2.StatusCode != http.StatusBadRequest {
		t.Errorf("convert code err, got %d want %d", errors.UnknownStatusCode, http.StatusBadRequest)
	}
	if errors.FromError(nil) != nil {
		t.Errorf("FromError(nil) should be nil")
	}
	e := errors.FromError(stderrors.New("test"))
	if !reflect.DeepEqual(e.StatusCode, errors.UnknownStatusCode) {
		t.Errorf("no expect value: %v, but got: %v", e.StatusCode, errors.UnknownStatusCode)
	}
}

func TestIs(t *testing.T) {
	tests := []struct {
		name string
		e    *errors.Error
		err  error
		want bool
	}{
		{
			name: "true",
			e: errors.New(
				404,
				errors.WithCode("test"),
				errors.WithMessage(""),
			),
			err: errors.New(
				http.StatusNotFound,
				errors.WithCode("test"),
				errors.WithMessage(""),
			),
			want: true,
		},
		{
			name: "false",
			e: errors.New(
				0,
				errors.WithCode("test"),
				errors.WithMessage(""),
			),
			err:  stderrors.New("test"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok := tt.e.Is(tt.err); ok != tt.want {
				t.Errorf("Error.Error() = %v, want %v", ok, tt.want)
			}
		})
	}
}

func TestCause(t *testing.T) {
	testError := &TestError{message: "test"}
	err := errors.BadRequest(
		errors.WithCode("foo"),
		errors.WithMessage("bar"),
		errors.WithCause(testError),
	)
	if !stderrors.Is(err, testError) {
		t.Fatalf("want %v but got %v", testError, err)
	}
	if te := new(TestError); stderrors.As(err, &te) {
		if te.message != testError.message {
			t.Fatalf("want %s but got %s", testError.message, te.message)
		}
	}
}

func TestOther(t *testing.T) {
	err := errors.Errorf(10001, "test code 10001", "message")
	// Code
	if !reflect.DeepEqual(errors.StatusCode(nil), http.StatusOK) {
		t.Errorf("Code(nil) = %v, want %v", errors.Code(nil), http.StatusOK)
	}
	if !reflect.DeepEqual(errors.StatusCode(stderrors.New("test")), errors.UnknownStatusCode) {
		t.Errorf(`Code(errors.New("test")) = %v, want %v`, errors.StatusCode(nil), 200)
	}
	if !reflect.DeepEqual(errors.StatusCode(err), 10001) {
		t.Errorf(`Code(err) = %v, want %v`, errors.StatusCode(err), 10001)
	}
	// Reason
	if !reflect.DeepEqual(errors.Code(nil), errors.UnknownCode) {
		t.Errorf(`Reason(nil) = %v, want %v`, errors.Code(nil), errors.UnknownCode)
	}
	if !reflect.DeepEqual(errors.Code(stderrors.New("test")), errors.UnknownCode) {
		t.Errorf(`Code(errors.New("test")) = %v, want %v`, errors.Code(nil), errors.UnknownCode)
	}
	if !reflect.DeepEqual(errors.Code(err), "test code 10001") {
		t.Errorf(`Reason(err) = %v, want %v`, errors.Code(err), "test code 10001")
	}
	// Clone
	err400 := errors.Newf(http.StatusBadRequest, "BAD_REQUEST", "param invalid")
	err400.Metadata = map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	if cErr := errors.Clone(err400); cErr == nil || cErr.Error() != err400.Error() {
		t.Errorf("Clone(err) = %v, want %v", errors.Clone(err400), err400)
	}
	if cErr := errors.Clone(nil); cErr != nil {
		t.Errorf("Clone(nil) = %v, want %v", errors.Clone(err400), err400)
	}
}
