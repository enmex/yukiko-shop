package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(ctx context.Context, port string, handler http.Handler) *http.Server {
	return &http.Server{
		Handler:     handler,
		Addr:        Addr(port),
		BaseContext: BaseContext(ctx),
	}
}

// StartServer starts a server that can respond to signals from the OS and terminate correctly
func StartServer(server *http.Server) error {
	// Make a channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start http server listening for requests.
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	// Blocking and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("running server: %w", err)
	case <-shutdown:
		return server.Close()
	}
}

func BaseContext(ctx context.Context) func(net.Listener) context.Context {
	return func(net.Listener) context.Context {
		return ctx
	}
}

func Addr(port string) string {
	return net.JoinHostPort("", port)
}
