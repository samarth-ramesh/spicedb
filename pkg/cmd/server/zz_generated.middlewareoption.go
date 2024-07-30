// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package server

import (
	dispatch "github.com/authzed/spicedb/internal/dispatch"
	defaults "github.com/creasty/defaults"
	helpers "github.com/ecordell/optgen/helpers"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	zerolog "github.com/rs/zerolog"
)

type MiddlewareOptionOption func(m *MiddlewareOption)

// NewMiddlewareOptionWithOptions creates a new MiddlewareOption with the passed in options set
func NewMiddlewareOptionWithOptions(opts ...MiddlewareOptionOption) *MiddlewareOption {
	m := &MiddlewareOption{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// NewMiddlewareOptionWithOptionsAndDefaults creates a new MiddlewareOption with the passed in options set starting from the defaults
func NewMiddlewareOptionWithOptionsAndDefaults(opts ...MiddlewareOptionOption) *MiddlewareOption {
	m := &MiddlewareOption{}
	defaults.MustSet(m)
	for _, o := range opts {
		o(m)
	}
	return m
}

// ToOption returns a new MiddlewareOptionOption that sets the values from the passed in MiddlewareOption
func (m *MiddlewareOption) ToOption() MiddlewareOptionOption {
	return func(to *MiddlewareOption) {
		to.Logger = m.Logger
		to.AuthFunc = m.AuthFunc
		to.EnableVersionResponse = m.EnableVersionResponse
		to.DispatcherForMiddleware = m.DispatcherForMiddleware
		to.EnableRequestLog = m.EnableRequestLog
		to.EnableResponseLog = m.EnableResponseLog
		to.DisableGRPCHistogram = m.DisableGRPCHistogram
		to.unaryDatastoreMiddleware = m.unaryDatastoreMiddleware
		to.streamDatastoreMiddleware = m.streamDatastoreMiddleware
	}
}

// DebugMap returns a map form of MiddlewareOption for debugging
func (m MiddlewareOption) DebugMap() map[string]any {
	debugMap := map[string]any{}
	debugMap["EnableVersionResponse"] = helpers.DebugValue(m.EnableVersionResponse, false)
	debugMap["EnableRequestLog"] = helpers.DebugValue(m.EnableRequestLog, false)
	debugMap["EnableResponseLog"] = helpers.DebugValue(m.EnableResponseLog, false)
	debugMap["DisableGRPCHistogram"] = helpers.DebugValue(m.DisableGRPCHistogram, false)
	return debugMap
}

// MiddlewareOptionWithOptions configures an existing MiddlewareOption with the passed in options set
func MiddlewareOptionWithOptions(m *MiddlewareOption, opts ...MiddlewareOptionOption) *MiddlewareOption {
	for _, o := range opts {
		o(m)
	}
	return m
}

// WithOptions configures the receiver MiddlewareOption with the passed in options set
func (m *MiddlewareOption) WithOptions(opts ...MiddlewareOptionOption) *MiddlewareOption {
	for _, o := range opts {
		o(m)
	}
	return m
}

// WithLogger returns an option that can set Logger on a MiddlewareOption
func WithLogger(logger zerolog.Logger) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.Logger = logger
	}
}

// WithAuthFunc returns an option that can set AuthFunc on a MiddlewareOption
func WithAuthFunc(authFunc auth.AuthFunc) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.AuthFunc = authFunc
	}
}

// WithEnableVersionResponse returns an option that can set EnableVersionResponse on a MiddlewareOption
func WithEnableVersionResponse(enableVersionResponse bool) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.EnableVersionResponse = enableVersionResponse
	}
}

// WithDispatcherForMiddleware returns an option that can set DispatcherForMiddleware on a MiddlewareOption
func WithDispatcherForMiddleware(dispatcherForMiddleware dispatch.Dispatcher) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.DispatcherForMiddleware = dispatcherForMiddleware
	}
}

// WithEnableRequestLog returns an option that can set EnableRequestLog on a MiddlewareOption
func WithEnableRequestLog(enableRequestLog bool) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.EnableRequestLog = enableRequestLog
	}
}

// WithEnableResponseLog returns an option that can set EnableResponseLog on a MiddlewareOption
func WithEnableResponseLog(enableResponseLog bool) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.EnableResponseLog = enableResponseLog
	}
}

// WithDisableGRPCHistogram returns an option that can set DisableGRPCHistogram on a MiddlewareOption
func WithDisableGRPCHistogram(disableGRPCHistogram bool) MiddlewareOptionOption {
	return func(m *MiddlewareOption) {
		m.DisableGRPCHistogram = disableGRPCHistogram
	}
}
