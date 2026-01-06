//go:build !darwin

package touchid

func Authenticate(promptMsg string) (bool, error) {
	return false, ErrOsNotSupported
}
