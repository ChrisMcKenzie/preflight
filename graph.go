package preflight

import (
	fmt "fmt"
	"sort"
	"strings"
	"sync"
)

var (
	nodeMutex sync.Mutex
)

type DepthWalkFunc func(Node, int) error

type Node interface{}

type NamedNode interface {
	Node

	Name() string
}

type Graph struct {
	nodes *Set
	from  map[interface{}]*Set
	to    map[interface{}]*Set
}

func NewGraph() *Graph {
	return &Graph{
		nodes: new(Set),
		from:  make(map[interface{}]*Set),
		to:    make(map[interface{}]*Set),
	}
}

func (g *Graph) AddNode(n Node) error {
	if ok := g.nodes.Include(n); ok {
		return fmt.Errorf("node already exists")
	}
	g.nodes.Add(n)
	return nil
}

func (g *Graph) SetEdge(e Edge) {
	var (
		from = e.Source()
		to   = e.Target()

		fromCode = hashcode(from)
		toCode   = hashcode(to)
	)

	if from == to {
		return
	}

	if s, ok := g.from[toCode]; ok && s.Include(from) {
		return
	}

	if !g.Has(from) {
		g.AddNode(from)
	}

	if !g.Has(to) {
		g.AddNode(to)
	}

	s, ok := g.from[toCode]
	if !ok {
		s = new(Set)
		g.from[toCode] = s
	}
	s.Add(from)

	s, ok = g.to[fromCode]
	if !ok {
		s = new(Set)
		g.to[fromCode] = s
	}
	s.Add(to)
}

func (g *Graph) Nodes() []Node {
	list := g.nodes.List()
	nodes := make([]Node, 0, len(list))
	for _, n := range list {
		nodes = append(nodes, n)
	}

	return nodes
}

func (g *Graph) From(n Node) *Set {
	s, ok := g.from[hashcode(n)]
	if !ok {
		return nil
	}

	return s
}

func (g *Graph) To(n Node) *Set {
	s, ok := g.to[hashcode(n)]
	if !ok {
		return nil
	}

	return s
}

func (g *Graph) Has(n Node) bool {
	return g.nodes.Include(n)
}

func (g *Graph) String() string {

	var buf strings.Builder

	nodes := g.Nodes()
	names := make([]string, 0, len(nodes))
	mapping := make(map[string]Node, len(nodes))
	for _, n := range nodes {
		name := NodeName(n)
		names = append(names, name)
		mapping[name] = n
	}

	sort.Strings(names)

	for _, name := range names {
		n := mapping[name]
		buf.WriteString(name)

		buf.WriteString(" => { ")
		from := g.from[hashcode(n)]
		deps := make([]string, 0, from.Len())
		for _, target := range from.List() {
			deps = append(deps, NodeName(target))
		}

		sort.Strings(deps)

		buf.WriteString(strings.Join(deps, ", "))
		buf.WriteString(" }\n")
	}

	return buf.String()
}

func (g *Graph) Neighbors(t Node) []Node {
	fromEdges := g.From(t).List()
	toEdges := g.To(t).List()

	neighbors := make([]Node, 0, len(fromEdges)+len(toEdges))

	for _, v := range fromEdges {
		neighbors = append(neighbors, v)
	}
	for _, v := range toEdges {
		neighbors = append(neighbors, v)
	}

	return neighbors
}

func (g *Graph) Root() (Node, error) {
	results := make([]Node, 0, 1)
	for _, n := range g.Nodes() {
		if g.From(n).Len() == 0 {
			results = append(results, n)
		}
	}

	if len(results) > 1 {
		return nil, fmt.Errorf("encountered multiple roots: %v", results)
	}

	return results[0], nil
}

func (g *Graph) RemoveEdge(e Edge) {
	if s, ok := g.from[e.Target()]; ok {
		s.Delete(e.Source())
	}

	if s, ok := g.to[e.Source()]; ok {
		s.Delete(e.Target())
	}
}

func (g *Graph) TransitiveReduction() {
	g.depthFirstWalk(g.Nodes(), true, func(n Node, d int) error {
		nTargets := g.To(n)
		ns := AsNodeList(nTargets)

		g.depthFirstWalk(ns, true, func(v Node, d int) error {
			shared := nTargets.Intersection(g.From(v))
			for range AsNodeList(shared) {
				g.RemoveEdge(BasicEdge(n, v))
			}

			return nil
		})

		return nil
	})
}

type vertexDepth struct {
	Node  Node
	Depth int
}

func (g *Graph) depthFirstWalk(start []Node, sorted bool, f DepthWalkFunc) error {
	visited := make(map[Node]struct{})
	viewpoint := make([]vertexDepth, len(start))
	for _, n := range start {
		viewpoint = append(viewpoint, vertexDepth{
			Node:  n,
			Depth: 0,
		})
	}

	for len(viewpoint) > 0 {
		i := len(viewpoint)
		current := viewpoint[i-1]
		viewpoint = viewpoint[:i-1]

		if _, ok := visited[current.Node]; ok {
			continue
		}
		visited[current.Node] = struct{}{}

		if err := f(current.Node, current.Depth); err != nil {
			return err
		}

		targets := AsNodeList(g.From(current.Node))

		if sorted {
			sort.Sort(byNodeName(targets))
		}

		for _, t := range targets {
			viewpoint = append(viewpoint, vertexDepth{
				Node:  t,
				Depth: current.Depth + 1,
			})
		}
	}

	return nil
}

func AsNodeList(s *Set) []Node {
	raw := s.List()
	vList := make([]Node, 0, len(raw))
	for _, v := range raw {
		vList = append(vList, v)
	}

	return vList
}

// NodeName returns the name of a vertex.
func NodeName(raw Node) string {
	switch v := raw.(type) {
	case NamedNode:
		return v.Name()
	case fmt.Stringer:
		return fmt.Sprintf("%s", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

type byNodeName []Node

func (b byNodeName) Len() int      { return len(b) }
func (b byNodeName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byNodeName) Less(i, j int) bool {
	return NodeName(b[i]) < NodeName(b[j])
}
