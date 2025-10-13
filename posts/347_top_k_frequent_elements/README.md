---
author: Jacob O'Neill
date: 13/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/top-k-frequent-elements", "github": "https://github.com/jacoboneill/blog/blob/main/posts/347_top_k_frequent_elements/top_k_frequent_elements.go"}
---

# Leetcode 347 Top K Frequent Elements

## The Problem

[Leetcode 347 Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements) has you solving the question:

> Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in **any order**.

So what does this mean? Basically we need to look through an array of numbers, get the count of each, and then create an array of the size `k` of the largest counts

**Example 1:**
```
Input: nums = [1,1,1,2,2,3], k=2
Output: [1, 2]
Explanation: 1 has a count of 3, 2 has a count of 2 and 3 has a count of 1, therefore the two largest counts are 1 and 2
```

**Example 2:**
```
Input: nums = [1], k=1
Output: [1]
Explanation: 1 has a count of 1, k is asking for the largest count therefore the answer is 1
```

**Example 1:**
```
Input: nums = [1,2,1,2,1,2,3,1,3,2], k=2
Output: [1, 2]
Explanation: 1 has a count of 4, 2 has a count of 4, 3 has a count of 2, therefore the two largest counts are 1 and 2
```

## My Really Bad Solution

So my initial thinking was *"Instead of generating an array of the number and the count and then sorting, we will create a hashmap of num to count and then loop k times to get the max, and delete that value from the map"*. This was my implementation of that

```go
type N struct {
    Number int
    Count  int
}

func Naive(nums []int, k int) []int {
    // Create hashmap of count
    m := make(map[int]int)
    for _, n := range nums {
        m[n]++
    }

    // Sort hashmap
    res := make([]int, k)

    for i := range k {
        max_num := N{}
        for n, c := range nums {
            if c > max_num.Count {
                max_num = N{n, c}
            }
        }
        delete(m, max_num.Number)
        res[i] = max_num.Number
    }

    return res
}
```

This is how the algorithm would work with `[1,1,1,2,2,3]`:
```
nums = [1,1,1,2,2,3]
k = 2

map = {}

for _, n := range nums
    n = 1: map = {1: 1}
    n = 1: map = {1: 2}
    n = 1: map = {1: 3}
    n = 2: map = {1: 3, 2: 1}
    n = 2: map = {1: 3, 2: 2}
    n = 3: map = {1: 3, 2: 2, 3: 1}

res = [0, 0]
for i := range k {
    i = 0
        max_num := {Number: 0, Count: 0}
        for n, c := range nums {
            n = 1, c = 3: c > max_num.Count => max_num = {Number: 1, Count: 3}
            n = 2, c = 2: c !> max_num.Count
            n = 3, c = 1: c !> max_num.Count
        }
        map = {2: 2, 3: 1}
        res[0] = 1

    i = 1
        max_num := {Number: 0, Count: 0}
        for n, c := range nums {
            n = 2, c = 2: c > max_num.Count => max_num = {Number: 2, Count: 2}    
            n = 3, c = 1: c !> max_num.Count
        }
        map = {3: 1}
        res[1] = 2
}

return res = [1, 2]
```

This works well, and is the most straight forward in my head. From what I can tell, the time complexity is $O(n \cdot k)$ as it will need to go through every `num` `k` times. It will also have a $O(n)$ space complexity as it will store each item once at max.

I believed this would be the right answer before puting it into Leetcode where I got the following stats:

> Runtime: 27ms Beats 5.4%
>
> Memory: 7.84MB Beats 76.9%

That runtime is awful. There must be a faster way.

## WTF is a Bucket Sort?!
So [Neetcode](https://www.youtube.com/watch?v=YPTqKIgVk-k) suggests to use something called a *"Bucket Sort"*. Never heard of that in my life but the general premise is very similar to what a `hashmap` does.

A bucket sort is not really a sorting algorithm, but more of a way to group data together. The diagrams on [Geeks for Geeks](https://www.geeksforgeeks.org/dsa/bucket-sort) makes it look very similar to how students are taught how hash maps work. Basically we take the *frequency* of the count, and use that as our hashing key. The main difference between a `hashmap` and bucket sort is a `hashmap` tries to spread the data as much as possible to avoid collisions, whereas bucket sort specifically aims for collisions, this is what groups the data together.

After watching the video, I created my own implementation in Go:
```go
func BucketSort(nums []int, k int) []int {
    // Create map of value to count
    m := make(map[int]int)
    for _, n := range nums {
        m[n]++
    }

    // Convert Map to Bucket Sort
    bs := make([][]int, len(nums)+1)
    for n, c := range m {
        bs[c] = append(bs[c], n)
    }

    // Get Largest From Bucket Sorted Array k Times
    var res []int
    for i := len(bs) - 1; i >= 0; i-- {
        res = append(res, n)
        if len(res) == k{
            return res
        }
    }

    return res
}
```

Now I am not a massive fan of this function, because in Go, we would not just return the `res` at the end, we would instead return an error and nil value. But because the Leetcode function signature is this, I didn't feel it was right to change it. 

> This is because the Leetcode constraints can't allow for the code to get to this path, but the Go compiler doesn't know that.

This leads to an $O(n)$ time and space complexity as we will only need to store each item once, and when getting the result from the *bucket sorted array*, this operation is also $O(n)$.
