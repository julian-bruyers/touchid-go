package touchid

import (
	"context"
	"time"
)

type authOptions struct {
	message       string
	timeout       time.Duration
	context       context.Context
	allowPassword bool
}

type authResult struct {
	success bool
	err     error
}

type Option func(*authOptions)

func WithMsg(msg string) Option {
	return func(option *authOptions) {
		option.message = msg
	}
}

func WithContext(ctx context.Context) Option {
	return func(option *authOptions) {
		option.context = ctx
	}
}

func WithTimeout(drt time.Duration) Option {
	return func(option *authOptions) {
		option.timeout = drt
	}
}

func WithPassword(allow bool) Option {
	return func(option *authOptions) {
		option.allowPassword = allow
	}
}
