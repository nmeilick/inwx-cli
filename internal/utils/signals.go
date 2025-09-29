package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

type GracefulShutdown struct {
	shutdownFuncs []func() error
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewGracefulShutdown() *GracefulShutdown {
	ctx, cancel := context.WithCancel(context.Background())
	return &GracefulShutdown{
		shutdownFuncs: make([]func() error, 0),
		ctx:           ctx,
		cancel:        cancel,
	}
}

func (gs *GracefulShutdown) AddShutdownFunc(fn func() error) {
	gs.shutdownFuncs = append(gs.shutdownFuncs, fn)
}

func (gs *GracefulShutdown) Context() context.Context {
	return gs.ctx
}

func (gs *GracefulShutdown) Start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Info().Msg("Received interrupt signal, shutting down gracefully...")

		// Execute shutdown functions in reverse order
		for i := len(gs.shutdownFuncs) - 1; i >= 0; i-- {
			if err := gs.shutdownFuncs[i](); err != nil {
				log.Error().Err(err).Msg("Error during shutdown")
			}
		}

		gs.cancel()
	}()
}

func (gs *GracefulShutdown) Wait() {
	<-gs.ctx.Done()
}
