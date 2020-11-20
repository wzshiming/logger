package logger

import (
	"context"

	"github.com/go-logr/logr"
)

var log = NewDelegatingLogger(DiscardLogger)

// DiscardLogger is a logr.Logger that does nothing.
var DiscardLogger = logr.Discard()

// Log is the base logger used by kubebuilder.  It delegates
// to another logr.Logger.  You *must* call SetLogger to
// get any actual logging.
var Log logr.Logger = log

// SetLogger sets a concrete logging implementation for all deferred Loggers.
func SetLogger(l logr.Logger) {
	log.Fulfill(l)
}

// FromContext returns a logger with predefined values from a context.Context.
func FromContext(ctx context.Context) logr.Logger {
	if ctx == nil {
		return Log
	}
	log := logr.FromContext(ctx)
	if log == nil {
		return Log
	}
	return log
}

// WithContext takes a context and sets the logger as one of its keys.
// Use FromContext function to retrieve the logger.
func WithContext(ctx context.Context, log logr.Logger) context.Context {
	return logr.NewContext(ctx, log)
}
