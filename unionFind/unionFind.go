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


 Algorithm Complexities
  |  Algorithm  | Init | Union | Find |
  |-----------------------------------|
  |  QuickFind  |   N  |   N   |  1   |
  |-----------------------------------|
  |  QuickUnion |   N  |  *N   |  1   |
  |-----------------------------------|

  *N -> Worst and unusual case.

*/
