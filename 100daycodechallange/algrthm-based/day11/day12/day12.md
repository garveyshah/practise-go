# 100day Coding Challange.

## Day 12: Write a program to remove duplicates from an array.

## Training Problems

### 1. Basic Functionality
Write a Go program that takes an array of integers and returns a new array with all duplicate elements removed. For example, given the input `[1, 2, 2, 3, 4, 4, 5]`, the program should output `[1, 2, 3, 4, 5]`.

### 2. Handling Edge Cases
Write a Go program to remove duplicates from an array that contains only one element, an empty array, and an array where all elements are the same. For example:
- Input: `[7]` → Output: `[7]`
- Input: `[]` → Output: `[]`
- Input: `[5, 5, 5, 5, 5]` → Output: `[5]`

### 3. Efficiency
Write a Go program to remove duplicates from a large array (e.g., 10,000 elements) and evaluate its time complexity. Describe any trade-offs between time complexity and space complexity in your solution.

### 4. Generic Types
Modify the program to handle an array of strings instead of integers. For example, given the input `["apple", "banana", "apple", "orange", "banana"]`, the program should output `["apple", "banana", "orange"]`.

### 5. Maintaining Order
Write a Go program to remove duplicates from an array while preserving the order of the first occurrence of each element. For example, given the input `[4, 5, 6, 4, 5, 7]`, the program should output `[4, 5, 6, 7]`.

### 6. In-place Modification
Write a Go program that removes duplicates from an array in-place, i.e., without using additional arrays or data structures. The program should modify the input array and return the new length of the array after duplicates are removed. For example, given the input `[3, 3, 2, 1, 4, 2, 4]`, the program should modify the array to `[3, 2, 1, 4]` and return `4`.

### 7. Complex Data Structures
Write a Go program to remove duplicates from an array of structures (e.g., custom structs with multiple fields). For example, given an array of `Person` structs with `Name` and `Age` fields, the program should remove duplicate `Person` objects based on the `Name` field.

### 8. Multidimensional Arrays
Write a Go program to remove duplicate rows from a 2D array (matrix). For example, given the input `[[1, 2], [3, 4], [1, 2], [5, 6], [3, 4]]`, the program should output `[[1, 2], [3, 4], [5, 6]]`.

### 9. Partial Duplicates
Write a Go program to remove elements from an array if they contain duplicate digits. For example, given the input `[121, 234, 321, 455, 789]`, the program should output `[234, 321, 789]` (since 121 and 455 have duplicate digits).

### 10. Conditional Removal
Write a Go program to remove duplicates from an array only if they meet a certain condition. For example, given the input `[12, 24, 12, 36, 48, 24, 60]` and the condition that the element must be divisible by 12, the program should output `[12, 24, 36, 48, 60]`.

### 11. Concurrent Processing
Write a Go program that uses goroutines to parallelize the removal of duplicates from an array. For example, split the array into chunks, process each chunk concurrently to remove duplicates, and then combine the results while ensuring no duplicates across chunks.

### 12. Immutable Data
Write a Go program that treats the input array as immutable and returns a new array with duplicates removed, ensuring that the original array remains unchanged. Discuss the benefits and trade-offs of immutability in this context.

### 13. Sorted Arrays
Write a Go program that removes duplicates from a sorted array. Optimize your solution to run in O(n) time by taking advantage of the sorted order. For example, given the input `[1, 1, 2, 2, 3, 4, 4, 5]`, the program should output `[1, 2, 3, 4, 5]`.

### 14. Handling Null Values
Write a Go program that removes duplicates from an array containing null values (or zero values in Go's case, since Go doesn't have null). For example, given the input `[1, 2, 0, 1, 3, 0, 4, 5]`, the program should output `[1, 2, 0, 3, 4, 5]`.

### 15. Memory Constraints
Write a Go program to remove duplicates from an array with strict memory constraints, limiting additional memory usage to a constant amount (O(1) space complexity). Discuss the limitations and potential performance impacts of this approach.

### 16. Custom Equality Criteria
Write a Go program that removes duplicates based on a custom equality criterion provided as a function. For example, consider two elements equal if their absolute difference is less than a given threshold. For input `[1, 3, 5, 7, 9]` with a threshold of `2`, the output should be `[1, 5, 9]`.

### 17. Time-Stamped Data
Write a Go program to remove
