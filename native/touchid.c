//go:build darwin

#include <stdio.h>
#include <stdbool.h>
#include <Foundation/Foundation.h>
#include <LocalAuthentication/LocalAuthentication.h>

#define AUTH_SUCCESS             1
#define AUTH_FAILED_OR_CANCELLED 0
#define AUTH_NOT_AVAILABLE      -1
#define AUTH_ERROR_INTERNAL     -2

#define TOUCHID_DEFAULT_MSG "Touch ID"

int IsAvailable() {
    @try {
        LAContext *context = [[LAContext alloc] init];
        NSError *authError = nil;
        if ([context canEvaluatePolicy:LAPolicyDeviceOwnerAuthenticationWithBiometrics error:&authError]) {
            return 1;
        }
    }
    @catch (id exception) { // safety wildcard catch to prevent errors in the main program
        return 1;
    }
    return 0;
}

int AuthenticateUser(char* prompt, bool allowPassword) {
    @try {
        LAContext *context = [[LAContext alloc] init];
        NSError *authError = nil;
        NSString *msg = nil;

        // Check for empty prompt string
        if (prompt != NULL) {
            msg = [NSString stringWithUTF8String:prompt];
        }
        if (msg == nil || [msg length] == 0) {
            msg = @TOUCHID_DEFAULT_MSG;
        }

        // Evaluate policy (biometrics & password or biometrics only)
        LAPolicy policy;
        if (allowPassword) {
            policy = LAPolicyDeviceOwnerAuthentication;
            context.localizedFallbackTitle = nil;
        } else {
            policy = LAPolicyDeviceOwnerAuthenticationWithBiometrics;
            context.localizedFallbackTitle = @"";   // Remove the "Use password..." button
        }

        dispatch_semaphore_t sema = dispatch_semaphore_create(0);
        __block int result = AUTH_ERROR_INTERNAL;

        if ([context canEvaluatePolicy:policy error:&authError]) {
            [context evaluatePolicy:policy
                    localizedReason:msg
                              reply:^(BOOL success, NSError *error) {

                if (success) {
                    result = AUTH_SUCCESS;
                } else {
                    if (error.code == LAErrorUserCancel || error.code == LAErrorSystemCancel) {
                        result = AUTH_FAILED_OR_CANCELLED;
                    } else if (error.code == LAErrorUserFallback) {
                        result = AUTH_FAILED_OR_CANCELLED;
                    } else if (error.code == LAErrorBiometryNotAvailable || error.code == LAErrorBiometryNotEnrolled) {
                        result = AUTH_NOT_AVAILABLE;
                    } else {
                        result = AUTH_ERROR_INTERNAL;
                    }
                }
                dispatch_semaphore_signal(sema);
            }];
        } else {
            return AUTH_NOT_AVAILABLE;
        }

        dispatch_semaphore_wait(sema, DISPATCH_TIME_FOREVER);
        return result;

    }
    @catch (id exception) { // safety wildcard catch to prevent errors in the main program
        return AUTH_ERROR_INTERNAL;
    }
}
