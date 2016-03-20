package unionFind

import "fmt"

type weightedQAWithPCompNode struct {
	weight int
	parent int
}

func NewWeightedQAWithPComp(size int) *WeightedQAWithPComp {
	w := &WeightedQAWithPComp{}
	w.ids = make([]weightedQAWithPCompNode, size)
	for i := range w.ids {
		w.ids[i].parent = i
		w.ids[i].weight = 1
	}
	return w
}

type WeightedQAWithPComp struct {
	ids []weightedQAWithPCompNode
}

func (w *WeightedQAWithPComp) Root(a int) int {
	parent := w.ids[a].parent
	for parent != a {
		w.ids[a].parent = w.ids[w.ids[a].parent].parent
		a = parent
	}
	defer fmt.Println(w.ids)
	return a
}

func (w *WeightedQAWithPComp) Union(a, b int) {
	a = w.Root(a)
	b = w.Root(b)
	if w.ids[a].weight < w.ids[b].weight {
		a, b = b, a
	}
	w.ids[a].weight += w.ids[b].weight
	w.ids[b].parent = a
}

func (w *WeightedQAWithPComp) Find(a, b int) bool {
	return w.Root(a) == w.Root(b)
}
