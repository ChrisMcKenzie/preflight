package preflight

import (
	context "context"
	"fmt"
	"sync"
)

type WalkFunc func(Node) error

type Walker struct {
	Callback WalkFunc

	g *Graph

	q    chan Node
	wg   sync.WaitGroup
	errc chan error
	ctx  context.Context
}

func NewWalker(cb WalkFunc) *Walker {
	return &Walker{
		Callback: cb,
		q:        make(chan Node),
		errc:     make(chan error),
	}
}

func (w *Walker) Run(ctx context.Context, g *Graph) error {
	w.g = g
	w.g.TransitiveReduction()
	go w.start(ctx)

	for _, n := range g.Nodes() {
		if g.From(n).Len() == 0 {
			go w.runNode(ctx, n, &w.wg)
		}
	}

	w.wg.Wait()
	return nil
}

func (w *Walker) Close() error {
	return nil
}

func (w *Walker) start(ctx context.Context) {
	defer w.Close()

	for {
		select {
		case err := <-w.errc:
			fmt.Println(err)
			return
		}
	}
}

func (w *Walker) runNode(ctx context.Context, n Node, wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		wg.Done()
	}(wg)

	err := w.Callback(n)
	if err != nil {
		w.errc <- err
		return
	}

	var dwg sync.WaitGroup
	for _, n := range w.g.To(n).List() {
		dwg.Add(1)
		go w.runNode(ctx, n, &dwg)
	}
	dwg.Wait()
}
