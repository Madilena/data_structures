# Big O

These are my misc. notes on big O

Big O means order of the function
Big O is the worst case estimate of time and space complexity
Big Theta = best case

## O(1) -> Constant Time

Example: getting an index from an array

The array could have size=1 or 10000 but the function would still require 1 step

## O(n) -> Linear Time

Example: going through each element in an array until you find the desired element

If n = 10 , we have 10 operations. If n=1000, we have 1000 operations. Think slope = 1

## O(log(n))

Example: Binary Search

The most common attributes of logarithmic running-time function are that:

- the choice of the next element on which to perform some action is one of several possibilities, and only one will need to be chosen.
  or
- the elements on which the action is performed are digits of n

![Alt text](big_o.png?raw=true"Title")
