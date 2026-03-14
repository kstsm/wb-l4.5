package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/slog"
	"github.com/kstsm/wb-l4.5/config"
	"github.com/kstsm/wb-l4.5/internal/handler"
	"github.com/kstsm/wb-l4.5/internal/service"
)

const (
	httpServerShutdownTimeout = 5 * time.Second
	readHeaderTimeout         = 5 * time.Second
)

func Run(log *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.GetConfig()

	svc := service.NewService(log)
	h := handler.NewHandler(svc, log)

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:           h.NewRouter(),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	errChan := make(chan error, 2)

	go func() {
		log.Infof("HTTP server started on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("main server error: %w", err)
		}
	}()

	go func() {
		pprofAddr := fmt.Sprintf("%s:%d", cfg.Pprof.Host, cfg.Pprof.Port)
		log.Infof("pprof starting on %s", pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("pprof server error: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Info("Shutting down servers due to signal...")
	case err := <-errChan:
		log.Errorf("Server error occurred: %v", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), httpServerShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Errorf("Shutdown error: %v", err)
		return err
	}

	return nil
}
