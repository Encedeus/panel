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
    ErrInvalidTokenType           = errors.New("invalid JWT  type")
    ErrInvalidAPIKeyDescription   = NewValidationError("invalid API key description")
    ErrInvalidUserId              = NewValidationError("invalid user id")
    ErrInvalidIPAddress           = NewValidationError("invalid IP address")
    ErrInvalidEmail               = NewValidationError("invalid email")
    ErrInvalidUsername            = NewValidationError("invalid username")
    ErrInvalidRoleID              = NewValidationError("invalid role id")
    ErrInvalidRoleName            = NewValidationError("invalid role name")
    ErrInvalidPassword            = NewValidationError("invalid password")
    ErrInvalidPermission          = NewValidationError("invalid permission")
    ErrOldUsernameDoesNotMatch    = NewValidationError("old username does not match current one")
    ErrNewUsernameEqualsOld       = NewValidationError("old username equals new one")
    ErrOldPasswordDoesNotMatch    = NewValidationError("old password does not match current one")
    ErrNewPasswordEqualsOld       = NewValidationError("old password equals new one")
    ErrOldEmailDoesNotMatch       = NewValidationError("old email does not match current one")
    ErrNewEmailEqualsOld          = NewValidationError("old email equals new one")
    ErrUserNotFound               = errors.New("user not found")
    ErrWrongPassword              = errors.New("wrong password")
    ErrUsernameAlreadyTaken       = NewValidationError("username already taken")
    ErrEmailAlreadyTaken          = NewValidationError("email already taken")
    ErrUnexpectedJWTSigningMethod = NewValidationError("unexpected JWT signing method")
    ErrUserDeleted                = errors.New("user deleted")
    ErrRoleDeleted                = errors.New("role deleted")
    ErrMissingAPIKey              = NewValidationError("missing API key")
)
