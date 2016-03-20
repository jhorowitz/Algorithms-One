package unionFind

import "testing"

type unionFindTestDef interface {
	Union(i, j int)
	Find(i, j int) bool
}

const size = 10

func TestQuickFind(t *testing.T) {
	UnionFindImplementationTest(NewQuickFind(size), t)
}

func TestQuickUnion(t *testing.T) {
	UnionFindImplementationTest(NewQuickUnion(size), t)
}

func UnionFindImplementationTest(q unionFindTestDef, t *testing.T) {
	if q.Find(5, 3) {
		t.Error("Found it before union")
	}
	q.Union(5, 3)
	if !q.Find(5, 3) {
		t.Error("Didn't Find 5, 3 it after union of 5, 3")
	}

	if q.Find(3, 7) {
		t.Error("3 and 7 aren't connected")
	}
	q.Union(5, 7)
	if !q.Find(3, 7) {
		t.Error("Didn't Find 3, 7 it after union of 5, 7")
	}
}
