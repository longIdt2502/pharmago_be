package common

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type AppError struct {
	StatusCode   int    `json:"status_code"`
	RootErr      error  `json:"-"`
	Message      string `json:"message"`
	MessageTrans string `json:"message_trans"`
	Log          string `json:"log"`
	Key          string `json:"error_key"`
}

func NewErrorResponse(root error, msg, msgt, log, key string) *AppError {
	return &AppError{
		StatusCode:   int(codes.InvalidArgument),
		RootErr:      root,
		Message:      msg,
		MessageTrans: msgt,
		Log:          log,
		Key:          key,
	}
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

func NewFullErrorResponse(statusCode int, root error, msg, msgt, log, key string) *AppError {
	return &AppError{
		StatusCode:   statusCode,
		RootErr:      root,
		Message:      msg,
		MessageTrans: msgt,
		Log:          log,
		Key:          key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode:   http.StatusUnauthorized,
		RootErr:      root,
		Message:      msg,
		MessageTrans: "Lỗi xác thực người dùng",
		Key:          key,
	}
}

func NewCustomError(root error, msg, msgt, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, msgt, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msgt, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", "lỗi truy cập db", err.Error(), ErrorDb)
}

func ErrDBWithMsg(err error, msgt string) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", msgt, err.Error(), ErrorDb)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", "truy cập không hợp lệ", err.Error(), ErrInvalid)
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", "lỗi server", err.Error(), ErrInternalCode)
}

func ErrInternalWithMsg(err error, msgt string) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", msgt, err.Error(), ErrInternalCode)
}
