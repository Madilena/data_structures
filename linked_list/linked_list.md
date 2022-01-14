# Linked List Notes

These are my misc. notes on the linked list data structure

## What is a linked list?

A linked list is a linear data structure, where the elements are not stored at contiguous memory locations.

A linked list consists of nodes where each node has a data value and reference to the next node in the list.
These nodes are linked using pointers.

```
 Adding nodes to linked list

head 1 has memory address 0xc00010a020
node 2 has memory address 0xc000114040
node 3 has memory address 0xc000114058
node 4 has memory address 0xc000114070
node 5 has memory address 0xc000114088
node 6 has memory address 0xc0001140a0

 Just printing again

head 1 has memory address 0xc00010a028
node 2 has memory address 0xc000114040
node 3 has memory address 0xc000114058
node 4 has memory address 0xc000114070
node 5 has memory address 0xc000114088
node 6 has memory address 0xc0001140a0

 Now deleting node with value 2

head 1 has memory address 0xc00010a030
node 3 has memory address 0xc000114058
node 4 has memory address 0xc000114070
node 5 has memory address 0xc000114088
node 6 has memory address 0xc0001140a0

 Now adding a node with value 7

head 1 has memory address 0xc00010a038
node 3 has memory address 0xc000114058
node 4 has memory address 0xc000114070
node 5 has memory address 0xc000114088
node 6 has memory address 0xc0001140a0
node 7 has memory address 0xc0001140b8
```

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
