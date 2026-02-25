package model

import "fmt"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrUnauthorized           = &AppError{Code: "UNAUTHORIZED", Message: "인증이 필요합니다", Status: 401}
	ErrForbidden              = &AppError{Code: "FORBIDDEN", Message: "권한이 없습니다", Status: 403}
	ErrNotFound               = &AppError{Code: "NOT_FOUND", Message: "리소스를 찾을 수 없습니다", Status: 404}
	ErrValidation             = &AppError{Code: "VALIDATION_ERROR", Message: "입력값이 올바르지 않습니다", Status: 400}
	ErrInvalidStatusTransition = &AppError{Code: "INVALID_STATUS_TRANSITION", Message: "잘못된 상태 전이입니다", Status: 400}
	ErrInternal               = &AppError{Code: "INTERNAL_ERROR", Message: "서버 내부 오류가 발생했습니다", Status: 500}
)

func NewAppError(code, message string, status int) *AppError {
	return &AppError{Code: code, Message: message, Status: status}
}

func WrapNotFound(entity string) *AppError {
	return &AppError{Code: "NOT_FOUND", Message: fmt.Sprintf("%s을(를) 찾을 수 없습니다", entity), Status: 404}
}
