# Linked List Notes

These are my misc. notes on the linked list data structure

## What is a linked list?

A linked list is a linear data structure, where the elements are not stored at contiguous memory locations.

A linked list consists of nodes where each node has a data value and reference to the next node in the list.
These nodes are linked using pointers.

http://cslibrary.stanford.edu/103/LinkedListBasics.pdf

![Alt text](linked_list.png?raw=true"Title")

Reminder:

Heap part of memory is long lived (this is not heap data structure)

- static or global variables live here
- has suff that lives past a single function call

Stack part of memory is short lived

- does not have return
- automatic growth/shrinkage
- last in first out (LIFO)
- you can have multiple stacks, concurrently
