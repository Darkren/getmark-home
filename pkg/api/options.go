package api

import "time"

// Option is an optional argument for the api.
type Option func(a *api)

// WithShutdownTimeout passes shutdown timeout to the api.
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(a *api) {
		a.shutdownTimeout = &timeout
	}
}

// applyOptions applies all the passed options to the api instance.
func applyOptions(api *api, opts []Option) {
	for _, opt := range opts {
		opt(api)
	}
}
