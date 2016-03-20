package unionFind

/*
 Uses same the same []int data structure as QuickUnion
 but interprets it as a tree.

 Unions are done by setting the root of one tree to the
 root of the other. Find is done by checking if both
 ids have the same root.
*/

func NewQuickUnion(size int) *QuickUnion {
	q := &QuickUnion{}
	q.ids = make([]int, size)
	for i := range q.ids {
		q.ids[i] = i
	}
	return q
}

type QuickUnion struct {
	ids []int
}

func (q *QuickUnion) Root(a int) int {
	for a != q.ids[a] {
		a = q.ids[a]
	}
	return a
}

func (q *QuickUnion) Union(a, b int) {
	q.ids[q.Root(a)] = q.Root(b)
}

func (q *QuickUnion) Find(a, b int) bool {
	return q.Root(a) == q.Root(b)
}
