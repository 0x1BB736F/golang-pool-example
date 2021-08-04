package pool

import (
	"context"

	"github.com/pkg/errors"
)

type BuilderPool struct {
	Elements []*Item
	poolSize int
	poolChan chan int
}

const (
	minimalSize = 1
)

var (
	ErrTimeoutDone = errors.New("error [context]: context timeout done")
)

func New(poolSize int) *BuilderPool {
	p := BuilderPool{}

	validateSize(&poolSize)

	c := make(chan int, poolSize)

	for i := 0; i < poolSize; i++ {
		p.Elements = append(p.Elements, newItem(i, c))
	}

	p.poolChan = c
	p.poolSize = poolSize

	return &p
}

func validateSize(ps *int) {
	if *ps <= 0 {
		*ps = 1
	}
}

func (bp *BuilderPool) Get(ctx context.Context) (*Item, error) {
	select {
	case <-ctx.Done():
		return nil, ErrTimeoutDone
	case idx := <-bp.poolChan:
		return bp.Elements[idx], nil
	}
}
