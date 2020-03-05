package config

import (
	fmt "fmt"
	"log"
	strings "strings"
	"sync"

	"github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/goombaio/dag"
)

type Config struct {
	Tasks     map[string]*plugin.TaskItem
	Variables []Variable

	nodes map[string]*dag.Vertex
	graph *dag.DAG
	wg    sync.WaitGroup
}

func (c *Config) GoString() string {
	str := "Config {\n  Tasks: {\n"
	for k, item := range c.Tasks {
		result := fmt.Sprintf("%s: %+v", k, item)
		s, ok := item.Task.(fmt.GoStringer)
		if ok {
			result = s.GoString()
		}
		str += strings.Replace(result, "\n", "    \n", -1) + "\n"
	}

	str += "  }\n}"

	return str
}

func (c *Config) Graph() (*dag.DAG, error) {
	if c.graph == nil {
		c.graph = dag.NewDAG()
	}

	if c.nodes == nil {
		c.nodes = make(map[string]*dag.Vertex)
	}

	if err := c.addNodes(); err != nil {
		return nil, err
	}

	if err := c.addEdges(); err != nil {
		return nil, err
	}

	return c.graph, nil
}

func (c *Config) addNode(id string, i *plugin.TaskItem) {
	n := dag.NewVertex(id, i)
	c.nodes[id] = n
	c.graph.AddVertex(n)
}

func (c *Config) addNodes() error {
	for n, item := range c.Tasks {
		c.addNode(n, item)
	}
	return nil
}

func (c *Config) addEdges() error {
	for n, item := range c.Tasks {
		node := c.nodes[n]
		for _, dep := range item.Meta.Dependencies {
			d, ok := c.nodes[dep]
			if !ok {
				continue
			}

			c.graph.AddEdge(d, node)
		}
	}
	return nil
}

func (c *Config) Traverse(f func(t *plugin.TaskItem)) {
	sv := c.graph.SourceVertices()
	c.resolveVertices(sv, f)
	c.wg.Wait()

}

func (c *Config) resolveVertices(vs []*dag.Vertex, f func(t *plugin.TaskItem)) {
	for _, item := range vs {
		c.wg.Add(1)
		go func(item *dag.Vertex, f func(t *plugin.TaskItem)) {
			if err := c.resolve(item, f); err != nil {
				log.Println(err)
			}
		}(item, f)
	}
}

func (c *Config) resolve(v *dag.Vertex, f func(t *plugin.TaskItem)) error {
	defer c.wg.Done()

	f(v.Value.(*plugin.TaskItem))
	vs, err := c.graph.Successors(v)
	if err != nil {
		return err
	}
	c.resolveVertices(vs, f)

	return nil
}
