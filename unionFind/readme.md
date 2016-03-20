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
 <pre>
----------------------------------------------------------------------------------------
|                  Algorithm                  | Init |      Union     |      Find      |
|--------------------------------------------------------------------------------------|
|                 Quick Find                  |   N  |        N       |       1        |
|--------------------------------------------------------------------------------------|
|                 Quick Union                 |   N  |       *N       |       1        |
|--------------------------------------------------------------------------------------|
|            Weighted Quick Union             |   N  |      log(n)    |     log(n)     |
|--------------------------------------------------------------------------------------|
|  Weighted Quick Union With Path Compression |   N  |   log(log(n))  |   log(log(n))  |
---------------------------------------------------------------------------------------- 
</pre>

  ```*N -> Worst and unusual case.```
