# Hash Table

These are my misc. notes on the hash table structure

## What is a hash table?

We use a hash table to quickly lookup if a value exists in a set of data.
For this implementation, we've got 3 main players

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

```
Example
We want to store the value SAMSON in a hash_table.
But where should we put it?

we'll create a hash function that brings together
* properties of the thing we want to store/lookup (the value)
* properties of the space that we want to store it in (the array)

Suppose our array size is 7
SAMSON has ASCII code sum: 465
For the array size of 7, we'll put it in index=3 for the array

...why?

    (ASCII sum of our value) % (array length) -> gives us a remainder

    (465) % 7 -> gives us remainder 3

    465 / 7 = 66 remainder 3 ..... (66 * 7 = 462)

    This remainder is the index we will store SAMSON at
```

What happens when we want to add a value that yields an index=3 from our hash function?
Example:
MAPOTOFU has ASCII code sum: 619 and in our array with size =7, this should be stored at index 3

So, we have 2 values with the same index. This is called a collision.

To address this, we will make the array an array of linked lists.
We will put collided values at the same index, but in the linked list at the index. This is called separate chaining, or open hashing.

Best case O(1)

Worst case O(n) - where, with separate chaining, all the values are at one index in a linked list, and we have to go through each linked list node to find our value

![Alt text](hash_table_separate_chaining.png?raw=true"Title")

### Fancie definition from [wikipedia](https://en.wikipedia.org/wiki/Hash_table)

A hash table is a data structure that implements an associative array abstract data type, a structure that can map keys to values.
An associative array (aka map or dictionary) is composed of a collection of (key, value) pairs, such that each possible key appears at most once in the collection.
A hash table uses a hash function to compute an index, also called a hash code, into an array of buckets or slots, from which the desired value can be found.

During lookup, the key is hashed and the resulting hash indicates where the corresponding value is stored.

Things we need to build the hash table:

```
Hash table part (array)

- structure - hash table
- method - insert
- method - search
- method - delete

bucket part (linked list)

- structure - bucket
- structure - bucketNode
- method - insert
- method - search
- method - delete
```
