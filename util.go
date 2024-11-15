package mageutil

import (
	"context"
	"strings"

	"github.com/magefile/mage/mg"
)

type (
	SerialFns   []mg.Fn
	ParallelFns []mg.Fn
)

func (s SerialFns) ID() string {
	ids := make([]string, 0, len(s))
	for _, f := range s {
		ids = append(ids, f.ID())
	}
	return strings.Join(ids, ":")
}

func (s SerialFns) Name() string {
	ids := make([]string, 0, len(s))
	for _, f := range s {
		ids = append(ids, f.Name())
	}
	return strings.Join(ids, ":")
}

func (s SerialFns) Run(ctx context.Context) error {
	fns := make([]interface{}, 0, len(s))
	for _, f := range s {
		fns = append(fns, f)
	}
	mg.SerialCtxDeps(ctx, fns...)
	return nil
}

func (p ParallelFns) ID() string {
	ids := make([]string, 0, len(p))
	for _, f := range p {
		ids = append(ids, f.ID())
	}
	return strings.Join(ids, ":")
}

func (p ParallelFns) Name() string {
	ids := make([]string, 0, len(p))
	for _, f := range p {
		ids = append(ids, f.Name())
	}
	return strings.Join(ids, ":")
}

func (p ParallelFns) Run(ctx context.Context) error {
	fns := make([]interface{}, 0, len(p))
	for _, f := range p {
		fns = append(fns, f)
	}
	mg.CtxDeps(ctx, fns...)
	return nil
}
