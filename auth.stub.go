//go:build !darwin

package touchid

func Available() bool {
	return false
}

func Authenticate(options ...authOptions) (bool, error) {
	return false, ErrOsNotSupported
}
