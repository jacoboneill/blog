---
author: Jacob O'Neill
date: 02/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/two-sum", "github": "https://github.com/jacoboneill/leetcode/blob/main/001_two_sum/two_sum.go"}
---
# Leetcode 001 Two Sum

Two sum is the first problem that many people first discover when they are learning DSA. It's a great first question that stumps beginner and advanced programmers alike.

There are two main strategies for this question, I like to call these the *"naive approach"* and the *"optimal approach"*. When I first attempted this question I originally solved it with the *naive approach* so no worries if you did too, just thank the editorial section on the Leetcode website for showing us all better solutions.

## The Naive Approach

The naive approach basically is to iterate over every combination in the two arrays until you get an answer:

```go
func naiveSolution(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1}
}
```

```py
def naive_solution(nums: list[int], target: int) -> list[int]:
    for i in range(len(nums) - 1):
        for j in range(i+1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i, j]

    return [-1]
```

This works, but due to the double `for` loops, it gives you a time complexity of $O(n^2)$ which is very inefficient with larger data sets, however the space complexity is $O(1)$ meaning that the same amount of memory will be used for an array of `nums` whether the length is 3 or 300,000.

This is the more obvious approach and for many (including myself) is the solution you first come up with.

## The Optimal Approach

For this approach, you use a `map` (or `dict`) as a lookup table for the numbers you have seen so far. This means you only need one pass of the array, but it means you use more memory.

```go
func optimalSolution(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, n := range nums {
        if j, ok := seen[target - n]; ok {
            return []int{j, i}
        } else {
            seen[target - n] = i
        }
    }

    return []int{-1}
}
```

```py
def optimalSolution(nums: list[int], target: int) -> list[int]:
    seen = dict()
    for i, n in enumerate(nums):
        if target - n in seen:
            return [seen[target - n], i]
        else:
            seen[target - n] = i

    return [-1]
```

This gives you a space complexity of $O(n)$ as, worst case, you will have to store every number in the lookup table, however because it is using a `map`, the lookup time, and therefore, the time complexity gets reduced to $O(n)$, making it far more time efficient for larger data sets.

## Conclusion
Overall, this problem teaches you the importance of lookup tables and not using a brute-force approach naively.
