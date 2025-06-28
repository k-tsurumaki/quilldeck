package errors

import "fmt"

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func New(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(err error, code, message string) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

// 共通エラーコード
const (
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeInternal     = "INTERNAL_ERROR"
)