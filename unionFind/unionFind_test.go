package unionFind

import "testing"

func TestQuickFind(t *testing.T) {
	q := NewQuickFind(10)
	if q.Find(5, 3) {
		t.Error("Found it before union")
	}
	q.Union(5, 3)
	if !q.Find(5, 3) {
		t.Error("Didn't Find it after union")
	}
	if q.Find(3, 7) {
		t.Fail()
	}
	q.Union(5, 7)
	if !q.Find(5, 3) {
		t.Error("Didn't Find it after union")
	}
}