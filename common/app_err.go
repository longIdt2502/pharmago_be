package common

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

// ProtoMessage implements protoiface.MessageV1.
func (*AppError) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements protoiface.MessageV1.
func (*AppError) Reset() {
	panic("unimplemented")
}

// String implements protoiface.MessageV1.
func (*AppError) String() string {
	panic("unimplemented")
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: int(codes.InvalidArgument),
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) GRPCStatus() *status.Status {
	st := status.New(codes.Code(e.StatusCode), e.Message)
	details, _ := st.WithDetails(e)
	return details
}

// Đệ quy
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), ErrorDb)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), ErrInvalid)
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), ErrInternalCode)
}
