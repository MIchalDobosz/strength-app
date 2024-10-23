package handlers

import (
	"fmt"
	"runtime"
)

func Error(text error) error {
	_, filename, line, _ := runtime.Caller(1)

	return fmt.Errorf("%s:%d %s", filename, line, text)
}

func Errorf(format string, a ...any) error {
	_, filename, line, _ := runtime.Caller(1)
	a = append([]any{filename, line}, a...)

	return fmt.Errorf("%s:%d "+format, a...)
}
