# Binary Search Tree

These are my misc. notes on the binary search tree structure

## What is a binary search tree?

[Visualization tool](https://www.cs.usfca.edu/~galles/visualization/BST.html)

Binary Search Tree is a node-based binary tree data structure which has the following properties:

- The left subtree of a node contains only nodes with keys lesser than the node's key.
- The right subtree of a node contains only nodes with keys greater than the nod's key.
- The left and right subtree each must also be a binary search tree.

## Big O

If bst is bushy, full we have O(logn)

- every step cuts the tree in half. for n, this is at most logn steps
- height of tree is O(log(n)), at most have to take height-of-tree step

If bst is not bushy :(, we have O(n)

- Potentially need to visit every element, just like linked lists

If we know the height of the tree, we say O(h)
