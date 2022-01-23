[On the Criteria to Be Used in Decomposing Systems into Modules](https://www.win.tue.nl/~wstomv/edu/2ip30/references/criteria_for_modularization.pdf) by D.L. Parnas

## what is a module?

- module = responsibility assignment, rather than submodule

* we must abandon the assumption that a module is one or more subroutines
* Rather, subroutines and programs are assembled collections of code from various modules

* each module is characterized by its knowledge of a design decision that it hides from others. Its interface or definition was chosen to reveal as little possible about its inner workings

## advisable decompositions

1. A data structure, its internal linkings, accessing procedures and modifying procedures are part of a single module

2. The sequence of instructions necessary to call a given routine and the routine itself are part of the same module.

3. The formats of control blocks used in queues in operating systems and similar programs must be hidden within a "control block module." It is conventional to make such formats the interfaces between various modules.

4. Character codes, alphabetic orderings, and similar data should be hidden in a module for greatest flexibility.

5. The sequence in which certain items will be processed should (as far as practical) be hidden within a single module.

## hierarchy and decomposition

- We have a hierarchical structure if a certain relation may be defined between the modules or programs and that relation is a partial ordering.

* The relation we are concerned with is "uses" or "depends upon." It is better to use a relation between programs since in many cases one module depends upon only part of another module

- hierarchical structure and "clean" decomposition are two desirable but independent properties of a system structure

* worst case: If we had designed a system in which the "low level" modules made some use of the "high level" modules, we would not have the hierarchy, wewouldfind it much harder to remove portions of the system, and "level" would not have much meaning in the system :grimacing:

## how to decompose a system?

- almost always incorrect to begin the decomposition of a system into modules on the basis of a flowchart. ie making each major step in the processing a module. This is wrong.

* Start with a list of difficult design decisions or design decisions which are likely to change

## clean code by Bob Martin mentions this

- objects expose functions and hide their data
  - obejcts hide their data behind abstractions and expose functions that operate on that data
- data structures expose data and have no significant behavior

  - data structures expose their data and have no meaningful functions

- a class does not simply push its variables out through getters and setters. Rather it exposes abstract interfaces that allow its users to maniuplate the essence of the data without having to know its implementation

* interfaces are all about de- coupling. we use interfaces to guarantee some functionality. exactly how that functionality is implemented does not affect the caller of the interface
