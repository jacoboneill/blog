---
author: Jacob O'Neill
date: 09/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/contains-duplicate", "github": "https://github.com/jacoboneill/blog/blob/main/posts/leetcode_217_contains_duplicate/contains_duplicate.go"}
---

# Leetcode 217 Contains Duplicate

This is my first attempt with using [Neetcode](https://neetcode.io/)'s roadmap. The first step is "Arrays & Hashing" and has Leetcode's 217 as the first question

## The Problem

> Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

**Example 1:**
```
Input: nums = [1, 2, 3, 1]
Output: true

Explanation:
The element 1 occurs at the indices 0 and 3.
```

**Example 2:**
```
Input: nums = [1, 2, 3, 4]
Output: false

Explanation:
All elements are distinct
```

**Example 3:**
```
Input: nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
Output: true

Explanation:
The element 1 occurs at indices 0, 1, and 2, element 3 occurs at indices 3, 4, and 6 and element 4 occurs at indices 5 and 8.
```

## Brute-Force

The brute force way to do this is to use a double `for` loop and check every digit with every other digit. If they match return `true`, else if we go through the entire loop, return `false`.

```go
func BruteForce(nums: []int) bool {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return true
            }
        } 
    }

    return false
}
```

This works, but is not the most optimal. Due to the double `for` loop, it will be $O(n^2)$ time, but due to not storing anything, it will be $O(1)$ space complexity.

## More Time Efficient

There is a more time-efficient technique you can use, by using a `hash set`. This will mean we can instantly look up if a value exists in the set.

> Go doesn't have a builtin `hash set`. So we have to use a `hash map` and lookup via the keys, not caring what the value is.
>
> This can either be done with a `map[int]struct{}`. This is because a `struct{}` is a zero-sized type [\[1\]](https://go.dev/ref/spec#Size_and_alignment_guarantees)

```go
func HashSet(nums []int) bool {
    s := make(map[int]struct{})

    for _, n := range nums {
        if _, ok := s[n]; ok {
            return true
        }

        s[n] = struct{}{}
    }

    return false
}
```

This now brings the time complexity to $O(n)$ as, at worst, it has to check each value once. The space complexity does unfortunatley go up to $O(n)$ as now at worst you have to store each value in the `hash set`.
