package unionFind

/*
 QuickFind is an "Eager algorithm" for solving the connectivity program.
 - Data structure for the algorithm is simply a []int

 P and Q are connected iff they share the same id
 ___________________
 0|1|2|3|4|5|6|7|8|9
 0|1|1|8|8|0|0|1|8|8

 In the above example, 0, 5, 6 are connected. 1, 2, 7 are connected. 3, 4, 8, 9 are connected.

 As a rough standard, computers can do 10^9 operations per second and have
 10^9 words of memory. This means they can touch all words in roughly 1 second

 This is a huge problem for QuickFind. Processing 10^9 union commands
 will take 30+ years of computer time. In general, quadratic algorithms
 don't scale with technology. New computer may be 10X as fast but it
 has 10X the amount of memory => want to solve a problem 10X as big.
 With quadratic algorithm, that takes 10X as long.

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
