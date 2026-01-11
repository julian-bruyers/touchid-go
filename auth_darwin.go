//go:build darwin

package touchid

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework LocalAuthentication -framework Foundation
#include <stdlib.h>
#include "native/touchid.c"
*/
import "C"

import (
	"unsafe"
)

func Available() bool {
	return C.IsAvailable() == 1
}

func Authenticate(promptMsg string) (bool, error) {
	cPrompt := C.CString(promptMsg)
	defer C.free(unsafe.Pointer(cPrompt))

	result := C.AuthenticateUser(cPrompt)

	switch result {
	case 1:
		return true, nil
	case 0:
		return false, ErrUserCanceled
	case -1:
		return false, ErrNotAvailable
	case -2:
		return false, ErrInternal
	default:
		return false, ErrInternal
	}
}
