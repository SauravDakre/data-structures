# 347. Top K Frequent Elements

https://leetcode.com/problems/top-k-frequent-elements/

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Ex 1: 
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

# Max Heap: 
A max-heap is a complete binary tree in which the value in each internal node is greater than or equal to the values in the children of that node.

Mapping the elements of a heap into an array is trivial: if a node is stored an index k, then its left child is stored at index 2k+1 and its right child at index 2k+2.

source: https://www.geeksforgeeks.org/max-heap-in-java/
