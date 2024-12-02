package runner

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type Runnable interface {
	Init() error
	Run() error
	Close() error
}

func Run(ctx context.Context, r Runnable) error {
	if err := r.Init(); err != nil {
		return fmt.Errorf("init failed: %w", err)
	}

	wg := &multierror.Group{}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Go(func() (err error) {
		defer cancel()
		return r.Run()
	})

	<-ctx.Done()
	wg.Go(r.Close)

	return wg.Wait().ErrorOrNil()
}
