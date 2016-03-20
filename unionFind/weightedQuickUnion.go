package unionFind

/*
 Weighted quick union is similar to QuickUnion except
 it will put shorter trees on the root of longer ones
 to avoid very tall trees.
*/

type weightedQuickUnionNode struct {
	parent int
	weight int
}

func NewWeightedQuickUnion(size int) *weightedQuickUnion {
	w := &weightedQuickUnion{}
	w.ids = make([]weightedQuickUnionNode, size)
	for i := range w.ids {
		w.ids[i].parent = i
		w.ids[i].weight = 1
	}
	return w
}

type weightedQuickUnion struct {
	ids []weightedQuickUnionNode
}

func (w weightedQuickUnion) Root(a int) int {
	parent := w.ids[a].parent
	for a != parent {
		a = parent
	}
	return a
}

func (w weightedQuickUnion) Union(a, b int) {
	a = w.Root(a)
	b = w.Root(b)
	if w.ids[a].weight < w.ids[b].weight {
		t := a
		a = b
		b = t
	}
	w.ids[a].parent = b
	w.ids[b].weight += w.ids[a].weight
}

func (w weightedQuickUnion) Find(a, b int) bool {
	return w.Root(a) == w.Root(b)
}
