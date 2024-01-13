# Algorithms in Go - Heaps

This project explores the implementation and application of heaps in the Go programming language.

## Basic Knowledge of Max Heaps

Heap: A specialized tree-based data structure that satisfies the heap property:

- Max-heap: Every parent node is greater than or equal to its children.
- Min-heap: Every parent node is less than or equal to its children.

## Common Operations and Time Complexity:

| Operation | Time Complexity | Description                                                 |
| --------- | --------------- | ----------------------------------------------------------- |
| Insert    | O(log n)        | Adds a new element to the heap                              |
| Extract   | O(log n)        | Removes and returns the largest element (max-heap)          |
| Peek Max  | O(1)            | Returns the largest element without removing it (max-heap). |
| Heapify   | O(log n)        | Maintains the heap property after a change.                 |

## Index Calculations:

Heaps are commonly represented using arrays. Given the index of a node i in the array representation:

- Parent index: `(i - 1) / 2` (using integer division)
- Left child index: `2 \* i + 1`
- Right child index: `2 \* i + 2`
