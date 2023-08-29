package runnertest

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

type ErrFn func() error

func Run(t testing.TB, run ErrFn, readyURL string) ErrFn {
	errCh := make(chan error, 1)
	go func() { errCh <- run() }()

	if err := AppReady(readyURL, errCh); err != nil {
		t.Fatal(err)
	}

	return func() error {
		Kill(t)

		const deadline = 5 * time.Second
		select {
		case err := <-errCh:
			return err
		case <-time.After(deadline):
			t.Fatalf("App didn't close within %s deadline", deadline)
			return nil
		}
	}
}

func AppReady(readyURL string, errCh <-chan error) error {
	const deadline = 5 * time.Second

	started := time.Now()
	for !IsReady(readyURL) {
		select {
		case err := <-errCh:
			return err
		default:
			if time.Since(started) > deadline {
				return fmt.Errorf("app didn't start within %s deadline", deadline)
			}
		}
	}
	return nil
}

func IsReady(readyURL string) bool {
	r, err := http.Get(readyURL)
	if err != nil {
		return false
	}
	r.Body.Close()
	return true
}

func Kill(t testing.TB) {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatal(err)
	}
}
