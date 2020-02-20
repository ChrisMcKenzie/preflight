package preflight

import (
	"context"
	fmt "fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraphWalk(t *testing.T) {
	g := newBasicGraph()

	w := NewWalker(func(n Node) error {
		fmt.Println(NodeName(n))
		rand.Seed(time.Now().Unix())
		r := rand.Intn(1)
		time.Sleep(time.Duration(r) * time.Millisecond)
		return fmt.Errorf("failed")
		// return nil
	})

	err := w.Run(context.TODO(), g)
	if err != nil {
		t.Fatal(err)
	}
}
