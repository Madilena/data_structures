[Designing software for ease of extension and contraction](https://courses.cs.washington.edu/courses/cse503/08wi/parnas-1979.pdf) by D.L. Parnas

## main things discussed

- families of programs in which some members are subsets of other family members or several family members share a common subset
- the stage when one identifies the major components of the system and defines relations between those components

## obstacles commonly encountered in trying to extend or shrink systems fall into four classes

1. Excessive Information Distribution
   e.g. operating system in which an early design decision was that the system would support three conversational languages, and we try to add a fourth or remove one

2. A Chain of Data Transforming Components
   e.g Many programs are structured as a chain of components,each receiving data from the previous component, processing it (and changing the format), before sending the data to the next program in the chain. If one component in this chain is not needed, that code is often hard to remove because the output of its predecessor is not compatible with the input requirements of its successor.

3. Components That Perform More Than One Function
   e.g. having a function that has multiple responsbilities makes it hard to leverage just one of the responsibilities in the future without the other (for performance reasons)

4. Loops in the "Uses"Relation
   one may end up with a system in which nothing works until everything works. For example, while it may seem wise to have an operating system scheduler use the file system to store its data (rather than use its own disk routines), the result will be that the file system must be present and working before any task scheduling is possible. There are users for whom an operating system subset without a file system would be useful.
