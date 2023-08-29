package services

import "errors"

type ValidationError struct {
    message string
}

func (e ValidationError) Error() string {
    return e.message
}

func NewValidationError(message string) ValidationError {
    ve := ValidationError{
        message: message,
    }

    return ve
}

func IsValidationError(err error) bool {
    if errors.As(err, &ValidationError{}) {
        return true
    }

    return false
}

var (
    ErrInvalidTokenType         = errors.New("invalid JWT  type")
    ErrInvalidAPIKeyDescription = NewValidationError("invalid API key description")
    ErrInvalidUserId            = NewValidationError("invalid user id")
    ErrInvalidIPAddress         = NewValidationError("invalid IP address")
    ErrInvalidEmail             = NewValidationError("invalid email")
    ErrInvalidUsername          = NewValidationError("invalid username")
    ErrInvalidRoleID            = NewValidationError("invalid role id")
    ErrInvalidRoleName          = NewValidationError("invalid role name")
    ErrInvalidPassword          = NewValidationError("invalid password")
)
