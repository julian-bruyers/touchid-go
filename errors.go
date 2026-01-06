package touchid

import "errors"

var (
	ErrOsNotSupported   = errors.New("touchid for Go is only available for darwin systems")
	ErrArchNotSupported = errors.New("touchid-go is only supported for amd64 and ARM architectures")
	ErrNotAvailable     = errors.New("biometric authentication is not available or configured on this system")
	ErrUserCanceled     = errors.New("the user canceled the authentication")
	ErrInternal         = errors.New("internal authentication error")
)
