package helper

import (
	"log/slog"
	"os"
)

func PanicIfError(err error) {
	if err != nil {
		slog.Error("Critical error encountered", "error", err)
		panic(err)
	}
}

func LogError(err error, message string, args ...any) {
	if err != nil {
		slog.Error(message, append(args, "error", err)...)
	}
}

func ExitIfError(err error) {
	if err != nil {
		slog.Error("Fatal error, exiting", "error", err)
		os.Exit(1)
	}
}
