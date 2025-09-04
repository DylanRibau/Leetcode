# Leetcode Questions
Some of my solutions for the leetcode [Top 100 Liked](https://leetcode.com/studyplan/top-100-liked/) questions

## Binary Tree - Symmetric Tree
[Solution](/Binary%20Tree/Symmetric%20Tree/main.go)

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).


Example 1:

    Input: root = [1,2,2,3,4,4,3]
    Output: true

Example 2:

    Input: root = [1,2,2,null,3,null,3]
    Output: false
 

Constraints:

    The number of nodes in the tree is in the range [1, 1000].
    -100 <= Node.val <= 100


## Binary Tree Inorder Traversal
[Solution](/Binary%20Tree/BinaryTreeInorderTraversal/main.go)

Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

    Input: root = [1,null,2,3]

    Output: [1,3,2]


Example 2:

    Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

    Output: [4,2,6,5,7,1,3,9,8]


Example 3:

    Input: root = []

    Output: []

Example 4:

    Input: root = [1]

    Output: [1]


Constraints:

    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100

## Letter Combinations of a Phone Number
[Solution](/GenerateParentheses/main.go)

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

    Input: n = 3
    Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

    Input: n = 1
    Output: ["()"]


Constraints:

    1 <= n <= 8


## Letter Combinations of a Phone Number
[Solution](/LetterCombinationsPhoneNumber/main.go)

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


Example 1:

    Input: digits = "23"
    Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

    Input: digits = ""
    Output: []

Example 3:

    Input: digits = "2"
    Output: ["a","b","c"]

 

Constraints:

    0 <= digits.length <= 4
    digits[i] is a digit in the range ['2', '9'].



## Find Minimum in Rotated Sorted Array
[Solution](/MinimumRotatedSortedArray/main.go)

Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

    [4,5,6,7,0,1,2] if it was rotated 4 times.
    [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

## Search a 2D Matrix
[Solution](/Search2DMatrix/main.go)

You are given an m x n integer matrix matrix with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

## Find First and Last Position of Element in Sorted Array
[Solution](/FindFirstAndLastPositionElementSortedArray/main.go)

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.