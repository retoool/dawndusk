package errors

import "fmt"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// Common errors
var (
	ErrUnauthorized       = NewAppError("UNAUTHORIZED", "Unauthorized", 401)
	ErrForbidden          = NewAppError("FORBIDDEN", "Forbidden", 403)
	ErrNotFound           = NewAppError("NOT_FOUND", "Resource not found", 404)
	ErrBadRequest         = NewAppError("BAD_REQUEST", "Bad request", 400)
	ErrInternalServer     = NewAppError("INTERNAL_SERVER", "Internal server error", 500)
	ErrEmailAlreadyExists = NewAppError("EMAIL_EXISTS", "Email already exists", 400)
	ErrInvalidCredentials = NewAppError("INVALID_CREDENTIALS", "Invalid credentials", 401)
	ErrUserNotFound       = NewAppError("USER_NOT_FOUND", "User not found", 404)
)
