package zap

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelStrings = map[string]zapcore.Level{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"dpanic": zap.DPanicLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
}

type encoderFlag struct {
	setFunc func(NewEncoderFunc)
	value   string
}

var _ flag.Value = &encoderFlag{}

func (ev *encoderFlag) String() string {
	return ev.value
}

func (ev *encoderFlag) Type() string {
	return "encoder"
}

func (ev *encoderFlag) Set(flagValue string) error {
	val := strings.ToLower(flagValue)
	switch val {
	case "json":
		ev.setFunc(newJSONEncoder)
	case "console":
		ev.setFunc(newConsoleEncoder)
	default:
		return fmt.Errorf("invalid encoder value %q", flagValue)
	}
	ev.value = flagValue
	return nil
}

type levelFlag struct {
	setFunc func(zapcore.LevelEnabler)
	value   string
}

var _ flag.Value = &levelFlag{}

func (ev *levelFlag) Set(flagValue string) error {
	level, validLevel := levelStrings[strings.ToLower(flagValue)]
	if !validLevel {
		logLevel, err := strconv.Atoi(flagValue)
		if err != nil {
			return fmt.Errorf("invalid log level %q", flagValue)
		}
		if logLevel > 0 {
			intLevel := -1 * logLevel
			ev.setFunc(zap.NewAtomicLevelAt(zapcore.Level(int8(intLevel))))
		} else {
			return fmt.Errorf("invalid log level %q", flagValue)
		}
	} else {
		ev.setFunc(zap.NewAtomicLevelAt(level))
	}
	ev.value = flagValue
	return nil
}

func (ev *levelFlag) String() string {
	return ev.value
}

func (ev *levelFlag) Type() string {
	return "level"
}

type stackTraceFlag struct {
	setFunc func(zapcore.LevelEnabler)
	value   string
}

var _ flag.Value = &stackTraceFlag{}

func (ev *stackTraceFlag) Set(flagValue string) error {
	level, validLevel := levelStrings[strings.ToLower(flagValue)]
	if !validLevel {
		return fmt.Errorf("invalid stacktrace level %q", flagValue)
	}
	ev.setFunc(zap.NewAtomicLevelAt(level))
	ev.value = flagValue
	return nil
}

func (ev *stackTraceFlag) String() string {
	return ev.value
}

func (ev *stackTraceFlag) Type() string {
	return "level"
}
