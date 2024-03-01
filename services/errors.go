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
	ErrInvalidTokenType              = errors.New("invalid JWT type")
	ErrInvalidAPIKeyDescription      = NewValidationError("invalid API key description")
	ErrInvalidUserId                 = NewValidationError("invalid user id")
	ErrInvalidIPAddress              = NewValidationError("invalid IP address")
	ErrInvalidEmail                  = NewValidationError("invalid email")
	ErrInvalidUsername               = NewValidationError("invalid username")
	ErrInvalidRoleID                 = NewValidationError("invalid role id")
	ErrInvalidRoleName               = NewValidationError("invalid role name")
	ErrInvalidPassword               = NewValidationError("invalid password")
	ErrInvalidPermission             = NewValidationError("invalid permission")
	ErrOldUsernameDoesNotMatch       = NewValidationError("old username does not match current one")
	ErrNewUsernameEqualsOld          = NewValidationError("old username equals new one")
	ErrOldPasswordDoesNotMatch       = NewValidationError("old password does not match current one")
	ErrNewPasswordEqualsOld          = NewValidationError("old password equals new one")
	ErrOldEmailDoesNotMatch          = NewValidationError("old email does not match current one")
	ErrNewEmailEqualsOld             = NewValidationError("old email equals new one")
	ErrUserNotFound                  = errors.New("user not found")
	ErrWrongPassword                 = errors.New("wrong password")
	ErrUsernameAlreadyTaken          = NewValidationError("username already taken")
	ErrEmailAlreadyTaken             = NewValidationError("email already taken")
	ErrUnexpectedJWTSigningMethod    = NewValidationError("unexpected JWT signing method")
	ErrUserDeleted                   = errors.New("user deleted")
	ErrRoleDeleted                   = errors.New("role deleted")
	ErrMissingAPIKey                 = NewValidationError("missing API key")
	ErrModuleNotFound                = errors.New("module not found")
	ErrInvalidFqdn                   = NewValidationError("invalid fqdn")
	ErrInvalidPrivateKeyOrPassphrase = NewValidationError("invalid ssh private key or passphrase")
	ErrInvalidPublicKey              = NewValidationError("invalid ssh public key")
	ErrFailedConnectingToSSH         = errors.New("failed connecting to ssh server")
	ErrFailedCreatingSSHSession      = errors.New("failed creating ssh session")
	ErrFailedConnectingToSkyhook     = errors.New("failed connecting to skyhook")
	ErrFailedGettingHardwareInfo     = errors.New("failed getting skyhook hardware info")
	ErrNodeNotFound                  = errors.New("node not found")
	ErrNodeAlreadyExists             = NewValidationError("node already exists")
	ErrServerNotFound                = errors.New("server not found")
	ErrServerAlreadyExists           = NewValidationError("server already exists")
	ErrUnsupportedVariant            = NewValidationError("unsupported crater variant")
	ErrNoFreeNodesFound              = NewValidationError("no free nodes satisfying the requirements exist")
	ErrFailedRemovingContainer       = errors.New("failed removing Docker container")
	ErrApiFailure                    = errors.New("API failed")
	ErrModuleHasNoReleases           = errors.New("module has no files to download")
)
