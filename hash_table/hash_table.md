# Hash Table

These are my misc. notes on the hash table structure

## What is a hash table?

We use a hash table to quickly lookup if a value exists in a set of data.
For this implementation, we've got 3 main characters

1. value we want to store and lookup
2. an array of linked lists
3. a hash function

Steps:

```
1. We take the value and put it in the hash function.

2. This hash function returns the index where that value should live in the array.

3. We add the value to the linked list at that index
or
We lookup the value in the linked list at that index
```

### Fancie definition from wikipedia

A hash table is a data structure that implements an associative array abstract data type, a structure that can map keys to values.
An associative array (aka map or dictionary) is composed of a collection of (key, value) pairs, such that each possible key appears at most once in the collection. A hash table uses a hash function to compute an index, also called a hash code, into an array of buckets or slots, from which the desired value can be found.

During lookup, the key is hashed and the resulting hash indicates where the corresponding value is stored.

Best case O(1)
Worst case O(n) - with separate chaining all the keys are at one index in a linked list

Hash table part (array)
structure - hash table
method - insert ()
method - search()
method - delete()

bucket part (linked list)
structure - bucket
structure - bucketNode
method - insert
method - search
method - delete
