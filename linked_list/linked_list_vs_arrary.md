# Linked List vs Array Notes

These are my misc. notes comparing linked lists to arrays
| array | linked_list |
| :------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------ |
| elements have contiguous memory locations | elements have non-contiguous memory locations |
| array size cannot change at runtime. we must know the upper limit on the number of elements in advance. | linked list size can change at runtime |
| any element in an array can be directly accessed with its index | all the previous elements in linked list must be traversed to reach any element |
| better cache locality in arrays (due to contiguous memory allocation) makes modifying a certain element in an array faster | random access is not allowed in linked lists. So we cannot do a binary search with linked lists. |
| inserting/deleting an element in an array is expensive because room has to be created for the new elements and old elements have to be shifted | inserting/ deleting an element is faster in linked lists (has O(1)), though it does take time to traverse and change the pointers. |
