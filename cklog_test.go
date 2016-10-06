package cklog

import (
	"testing"
	"fmt"
)

func TestStdoutLog(t *testing.T) {
	stdoutLogger1 := NewLogger("debug", "stdout", "")
	stdoutLogger1.Debug("level debug output %s log", "debug")
	stdoutLogger1.Info("level debug output %s log", "info")
	stdoutLogger1.Warn("level debug output %s log", "warn")
	stdoutLogger1.Err("level debug output %s log", "err")

	fmt.Println()

	stdoutLogger2 := NewLogger("info", "stdout", "")
	stdoutLogger2.Debug("level info output %s log", "debug")
	stdoutLogger2.Info("level info output %s log", "info")
	stdoutLogger2.Warn("level info output %s log", "warn")
	stdoutLogger2.Err("level info output %s log", "err")

	fmt.Println()

	stdoutLogger3 := NewLogger("warn", "stdout", "")
	stdoutLogger3.Debug("level warn output %s log", "debug")
	stdoutLogger3.Info("level warn output %s log", "info")
	stdoutLogger3.Warn("level warn output %s log", "warn")
	stdoutLogger3.Err("level warn output %s log", "err")

	fmt.Println()

	stdoutLogger4 := NewLogger("err", "stdout", "")
	stdoutLogger4.Debug("level err output %s log", "debug")
	stdoutLogger4.Info("level err output %s log", "info")
	stdoutLogger4.Warn("level err output %s log", "warn")
	stdoutLogger4.Err("level err output %s log", "err")
}

func TestFileLog(t *testing.T) {
	stdoutLogger1 := NewLogger("debug", "file", "./log-debug")
	stdoutLogger1.Debug("level debug output %s log", "debug")
	stdoutLogger1.Info("level debug output %s log", "info")
	stdoutLogger1.Warn("level debug output %s log", "warn")
	stdoutLogger1.Err("level debug output %s log", "err")

	fmt.Println()

	stdoutLogger2 := NewLogger("info", "file", "./log-info")
	stdoutLogger2.Debug("level info output %s log", "debug")
	stdoutLogger2.Info("level info output %s log", "info")
	stdoutLogger2.Warn("level info output %s log", "warn")
	stdoutLogger2.Err("level info output %s log", "err")

	fmt.Println()

	stdoutLogger3 := NewLogger("warn", "file", "./log-warn")
	stdoutLogger3.Debug("level warn output %s log", "debug")
	stdoutLogger3.Info("level warn output %s log", "info")
	stdoutLogger3.Warn("level warn output %s log", "warn")
	stdoutLogger3.Err("level warn output %s log", "err")

	fmt.Println()

	stdoutLogger4 := NewLogger("err", "file", "./log-err")
	stdoutLogger4.Debug("level err output %s log", "debug")
	stdoutLogger4.Info("level err output %s log", "info")
	stdoutLogger4.Warn("level err output %s log", "warn")
	stdoutLogger4.Err("level err output %s log", "err")
}