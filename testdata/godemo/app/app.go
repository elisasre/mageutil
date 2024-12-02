package app

import (
	"context"
	"errors"
	"fmt"
	"godemo/api"
	"godemo/docs"
	"net"
	"net/http"
	"os"
)

type App struct {
	srv *http.Server
	ln  net.Listener
}

func (a *App) Init() error {
	spec, err := docs.LoadSpec()
	if err != nil {
		return fmt.Errorf("failed to load spec: %w", err)
	}

	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: api.NewRouter(spec),
	}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return fmt.Errorf("failed to init listener: %w", err)
	}

	a.srv = srv
	a.ln = ln
	return nil
}

func (a *App) Run() error {
	err := a.srv.Serve(a.ln)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (a *App) Close() error {
	return a.srv.Shutdown(context.Background())
}
