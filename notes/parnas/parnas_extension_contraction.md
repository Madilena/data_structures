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

4. Loops in the "Uses" Relation
   one may end up with a system in which nothing works until everything works. For example, while it may seem wise to have an operating system scheduler use the file system to store its data (rather than use its own disk routines), the result will be that the file system must be present and working before any task scheduling is possible. There are users for whom an operating system subset without a file system would be useful.

## Steps towards a better structure

1. Requirements Definition - identify the subsets first
   identification of the potentially desirable subsets is a demanding intellectual exercise in which one first searches for the minimal subset that might conceivably perform a useful service and then searches for a set of minimal increments to the system.

2. Information Hiding - interface and module definition

1) Identification of the items that are likely to change. These items are termed "secrets."
2) Location of the specialized components in separate modules.
3) Designing intermodule interfaces that are insensitive to the anticipated changes. The changeable aspects or "secrets" of the modules are not revealed by the interface.

All data structures that reveal the presence or number of certain components should be included in separate information hiding modules with abstract interfaces

3. the virtual machine concenpt
   Rather than write programs that perform the transformation from input data to output data, we design software machine extensions that will be useful in writing many such programs.
   Where our hardware machine provides us with a set of instructions that operate on a small set of data types, the extended or virtual machine will have additional data types as well a "software instructions" that operate on those data types.

4. designing the uses structure

"Uses" can also be formulated as "requires the presence ofa correct version of."

The "uses" relation and "invokes" very often coincide, but uses differs from invokes in two ways:
a) Certain invocations may not be instances of "uses." If A's specification requires only that A invoke B when certain. conditions occur, then A has fulfilled its specification when it
has generated a correct call to B.

A is correct even if B is incorrect or absent. A proof of correctness of A need only make assumptions about the way to invoke B.

b) A program A may use B even though it never invokes it. The best illustration of this is interrupt handling. Most programs in a computer system are only correct on the assumption that the interrupt handling routine will leave the processor in an acceptable state
