package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// api is a general API.
type api struct {
	handler http.Handler

	server    *http.Server
	serveErrC chan error

	closeOnce       sync.Once
	shutdownTimeout *time.Duration
}

// New constructs new API.
func New(handler http.Handler) *api {
	return &api{
		handler: handler,
	}
}

// Serve starts serving on a specified host and port.
func (a *api) Serve(listen string) error {
	a.server = &http.Server{Addr: listen, Handler: a.handler}

	if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Stop stops the API.
func (a *api) Stop() error {
	err := fmt.Errorf("already stopped")

	a.closeOnce.Do(func() {
		if a.server == nil {
			err = fmt.Errorf("API is not started")
			return
		}

		ctx := context.Background()
		if a.shutdownTimeout != nil {
			var cancel func()
			ctx, cancel = context.WithTimeout(ctx, *a.shutdownTimeout)
			defer cancel()
		}

		err = a.server.Shutdown(ctx)
	})

	return err
}

// Run creates new API instance and starts it up.
func Run(listen string, handler http.Handler, opts ...Option) error {
	api := New(handler)
	applyOptions(api, opts)

	serveErrC := make(chan error)
	go func() {
		serveErrC <- api.Serve(listen)
		close(serveErrC)
	}()

	signalC := make(chan os.Signal)
	signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalC)

	select {
	case err := <-serveErrC:
		return err
	case <-signalC:
	}

	return api.Stop()
}
