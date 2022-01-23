# Heap Notes

These are my misc. notes on the heap data structure

## What is a heap?

- Heap says, 'I got a bunch of values and I just want to keep track of the min or max value'
- A heap is not a sorted structure. It can be regarded as partially ordered. It's possible to have different heaps for the same set of elements.
- A heap is a complete binary tree (all levels completely filled except maybe last level)
- A heap is backed by an array
- Heap does not allow duplicates
- Heap order property : For min heap, parent is less than its children. For max heap, parent is greater than its children.

## Big O for heap

This explains [optimizing heap building](https://stackoverflow.com/questions/9755721/how-can-building-a-heap-be-on-time-complexity)

Notice that only one node is at the top whereas half the nodes lie in the bottom layer. So, if we have to apply an operation to every node we would prefer to start at the end of the array and sift down (towards index=0) over siftUp. startIndex is last non-leaf node (array_length/2 - 1)
When we implement the sift down method, we have O(n). Sift up is O(logn)

## Details on running heap.go

```
example

starting with this array:
{100, 3, 25, 29, 52, 1, 8, 9, 11, 18, 19, 69, 75, 92, 22}

we create the min heap:
[1 3 8 9 18 25 22 29 11 52 19 69 75 92 100]

we can display how the sift down finds the min value

[100 3 25 29 52 1 8 9 11 18 19 69 75 92 22]
[100 3 25 29 18 1 8 9 11 52 19 69 75 92 22]
[100 3 25 9 18 1 8 29 11 52 19 69 75 92 22]
[100 3 1 9 18 25 8 29 11 52 19 69 75 92 22]
[1 3 100 9 18 25 8 29 11 52 19 69 75 92 22]
[1 3 8 9 18 25 100 29 11 52 19 69 75 92 22]
[1 3 8 9 18 25 22 29 11 52 19 69 75 92 100]

we can also get additional info about the heap nodes:

number of tree levels:  4
tree level is  1 array index is:  0 node value is  1
tree level is  2 array index is:  1 node value is  3
tree level is  2 array index is:  2 node value is  8
tree level is  3 array index is:  3 node value is  9
tree level is  3 array index is:  4 node value is  18
tree level is  3 array index is:  5 node value is  25
tree level is  3 array index is:  6 node value is  22
tree level is  4 array index is:  7 node value is  29
tree level is  4 array index is:  8 node value is  11
tree level is  4 array index is:  9 node value is  52
tree level is  4 array index is:  10 node value is  19
tree level is  4 array index is:  11 node value is  69
tree level is  4 array index is:  12 node value is  75
tree level is  4 array index is:  13 node value is  92
tree level is  4 array index is:  14 node value is  100


we can also display the min heap as a tree

[1]
[3 8]
[9 18 25 22]
[29 11 52 19 69 75 92 100]
```

## Our example with image

![Alt text](heap_example.png?raw=true"Title")
