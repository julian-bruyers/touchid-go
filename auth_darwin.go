//go:build darwin || cgo

package touchid

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework LocalAuthentication -framework Foundation
#include <stdlib.h>
#include "native/touchid.c"
*/
import "C"

import (
	"context"
	"unsafe"
)

func Available() bool {
	return C.IsAvailable() == 1
}

func Authenticate(options ...Option) (bool, error) {
	authOpts := &authOptions{
		message:       "",
		context:       context.Background(),
		allowPassword: true,
	}

	for _, current := range options {
		current(authOpts)
	}

	if authOpts.timeout > 0 {
		var cancel context.CancelFunc
		authOpts.context, cancel = context.WithTimeout(authOpts.context, authOpts.timeout)
		defer cancel()
	}

	authResultChannel := make(chan authResult, 1)

	go func() {
		success, err := authenticate(authOpts.message, authOpts.allowPassword)
		authResultChannel <- authResult{success, err}
	}()

	select {
	case result := <-authResultChannel:
		return result.success, result.err
	case <-authOpts.context.Done():
		return false, authOpts.context.Err()
	}

}

func authenticate(promptMsg string, allowPassword bool) (bool, error) {
	cPrompt := C.CString(promptMsg)
	defer C.free(unsafe.Pointer(cPrompt))

	result := C.AuthenticateUser(cPrompt, C.bool(allowPassword))

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
