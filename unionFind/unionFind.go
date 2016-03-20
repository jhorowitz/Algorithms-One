package unionFind

/*
Given a set of N objects.
 - Union command: connect two objects
 - Find / connected query: Is there a path connecting the two objects


Applications involve manipulating objects of all kinds
 - Computers in a network
 - Friends in a social network
 - pixels in a photo
 - transistors in a computer chip
 - elements in a mathematical set
 - variable names in a fortran program
 - Metallic sites in a composite system


We assume "is connected to" is an equivalence relation.
 - Reflexive: p is connected to p
 - Symmetric: If p is connected to q then q is connected to p
 - Transitive: If p is connected to q and q is connected to r then p is connected to r

Connected components: Maximal set of objects that are mutually connected

*/

/*
 Algorithm Complexity
  Algorithm | Initialize | Union | Find |
  QuickFind |      N     |   N   |  N   |
*/

/*
"Eager algorithm" for solving the connectivity program.
 - Data structure for the algorithm is simply a []int

 P and Q are connected iff they share the same id
 ___________________
 0|1|2|3|4|5|6|7|8|9
 0|1|1|8|8|0|0|1|8|8

 In the above example, 0, 5, 6 are connected. 1, 2, 7 are connected. 3, 4, 8, 9 are connected.

*/

func NewQuickFind(size int) *QuickFind {
	q := &QuickFind{}
	q.ids = make([]int, size)
	for i := range q.ids {
		q.ids[i] = i
	}
	return q
}

/*
 QuickFind is highly efficient for find and highly inefficient for union.

 Complexity:
 Initialize: N
 Union: N
 Find: 1

*/
type QuickFind struct {
	ids []int
}

/*
 To merge components, change all ids that == ids[j] to ids[i]
*/
func (qf *QuickFind) Union(a, b int) {
	var previousId = qf.ids[b]
	for i, id := range qf.ids {
		if id == previousId {
			qf.ids[i] = qf.ids[a]
		}
	}
}

func (qf *QuickFind) Find(i, j int) bool {
	return qf.ids[i] == qf.ids[j]
}
