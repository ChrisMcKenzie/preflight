package preflight

import (
	"fmt"
	"testing"
)

type MockTask struct{ name string }

func (m *MockTask) Evaluate()    {}
func (m *MockTask) Apply()       { fmt.Println("applying:", m.name) }
func (m *MockTask) Name() string { return "mock-task-" + m.name }

func (m *MockTask) GoString() string { return fmt.Sprintf("%+v", *m) }

func TestGraph(t *testing.T) {
	g := NewGraph()

	task1 := &MockTask{name: "t1"}
	task2 := &MockTask{name: "t2"}
	task3 := &MockTask{name: "t3"}
	task4 := &MockTask{name: "t4"}
	task5 := &MockTask{name: "t5"}
	task6 := &MockTask{name: "t6"}
	task7 := &MockTask{name: "t7"}

	g.AddNode(task1)
	g.AddNode(task2)
	g.AddNode(task3)
	g.AddNode(task4)
	g.AddNode(task5)
	g.AddNode(task6)
	g.AddNode(task7)

	g.SetEdge(BasicEdge(task1, task2))
	g.SetEdge(BasicEdge(task1, task3))
	g.SetEdge(BasicEdge(task2, task3))
	g.SetEdge(BasicEdge(task2, task4))
	g.SetEdge(BasicEdge(task4, task5))
	g.SetEdge(BasicEdge(task5, task6))

	g.TransitiveReduction()
	fmt.Println("Traversing graph...")
	fmt.Printf("%s\n", g.String())
}

func newBasicGraph() *Graph {
	g := NewGraph()

	task1 := &MockTask{name: "t1"}
	task2 := &MockTask{name: "t2"}
	task3 := &MockTask{name: "t3"}
	task4 := &MockTask{name: "t4"}
	task5 := &MockTask{name: "t5"}
	task6 := &MockTask{name: "t6"}
	task7 := &MockTask{name: "t7"}

	g.AddNode(task1)
	g.AddNode(task2)
	g.AddNode(task3)
	g.AddNode(task4)
	g.AddNode(task5)
	g.AddNode(task6)
	g.AddNode(task7)

	g.SetEdge(BasicEdge(task1, task2))
	g.SetEdge(BasicEdge(task1, task3))
	g.SetEdge(BasicEdge(task2, task3))
	g.SetEdge(BasicEdge(task2, task4))
	g.SetEdge(BasicEdge(task4, task5))
	g.SetEdge(BasicEdge(task5, task6))

	return g
}
