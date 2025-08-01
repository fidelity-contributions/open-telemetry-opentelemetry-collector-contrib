// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/log/log.go

package log // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/log"

// Fields is a map that is used to populated logging context.
type Fields map[string]any

type nilLogger struct{}

func (nilLogger) Debug(string) {
}

func (nilLogger) Warn(string) {
}

func (nilLogger) Error(string) {
}

func (nilLogger) Info(string) {
}

func (nilLogger) Panic(string) {
}

func (nilLogger) WithFields(Fields) Logger {
	return nilLogger{}
}

func (nilLogger) WithError(error) Logger {
	return nilLogger{}
}

// Nil logger is a silent logger interface.
var Nil = nilLogger{}

var _ Logger = (*nilLogger)(nil)

// Logger is generic logging interface.
type Logger interface {
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Info(msg string)
	Panic(msg string)
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}
