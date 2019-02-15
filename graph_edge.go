package preflight

type Edge interface {
	Source() Node
	Target() Node
}

type basicEdge struct {
	s, t Node
}

func BasicEdge(source, target Node) Edge {
	return &basicEdge{s: source, t: target}
}

func (bs *basicEdge) Source() Node {
	return bs.s
}

func (bs *basicEdge) Target() Node {
	return bs.t
}
