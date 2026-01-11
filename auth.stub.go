//go:build !darwin

package touchid

func Available() (bool) {
	return false
}

func Authenticate(promptMsg string) (bool, error) {
	return false, ErrOsNotSupported
}
