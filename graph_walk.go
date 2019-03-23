package preflight

import "sync"

type WalkFunc func(Node)

type Walker struct {
	Callback WalkFunc

	wg sync.WaitGroup
}

func (w *Walker) Run(g *Graph) {

}
