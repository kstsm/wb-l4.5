package main

import (
	"os"

	"github.com/gookit/slog"
	"github.com/kstsm/wb-l4.5/cmd"
	"github.com/kstsm/wb-l4.5/pkg/logger"
)

func main() {
	log := logger.NewSlogLogger()

	if err := cmd.Run(log); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}
