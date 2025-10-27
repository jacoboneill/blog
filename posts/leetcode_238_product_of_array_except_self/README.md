---
author: Jacob O'Neill
date: 27/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/product-of-array-except-self", "github": "https://github.com/jacoboneill/blog/blob/main/posts/leetcode_238_product_of_array_except_self/product_of_array_except_self.go"}
---

# Leetcode 238 Product of Array Except Self

## The Problem

[Leetcode 238 Product of Array Except self](https://leetcode.com/problems/product-of-array-except-self) has you solving the question:

> Given an integer array `nums`, return an array `output` where `output[i]` is the product of all the elements of `nums` except `nums[i]`.

**Example 1:**
```
Input: nums=[1, 2, 4, 6]
Output: [48, 24, 12, 8]

Explanation
i=0: 2*4*6=48
i=1: 1*4*6=24
i=2: 1*2*6=12
i=3: 1*2*4=8
```

A caveat to this challenge is that you are not allowed to use the `/` operator.

## Breaking the Rules

If we break the rules and allow us to use the `/` operator, this challenge becomes a whole lot easier. This is because we can cache the total product, and then divide by the current index:
```go
func Cheating(nums []int) []int {
    res := make([]int, len(nums))
    
    product := 1
    for _, n := range nums {
        product *= n
    }

    for i, n := range nums {
        res[i] = product / n
    }

    return res
}
```

There is one special case where this algorithm doesn't work: If one of the numbers is `0`, then the product will be `0` and we will have to try and divide by `0` at some point. We can fix for this though:


```go
func Cheating(nums []int) []int {
    res := make([]int, len(nums))

	product := 1
	zero_index := -1
	for i, n := range nums {
		if n == 0 {
			zero_index = i
		} else {
			product *= n
		}
	}

	if zero_index != -1 {
		res[zero_index] = product
		return res
	}

	for i, n := range nums {
		res[i] = product / n
	}
	return res
}
```

## Brute-Force

The easiest (but definitely not most optimised) is to use brute-force and not use any tricks and just calculate everything:

```go
func BruteForce(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	for i := range nums {
		for j, n := range nums {
			if i != j {
				res[i] *= n
			}
		}
	}

	return res
}
```

The first part of this algorithm initialises an array of 1's to be the same length of the input. The second part then iterates through the array twice, making sure to only multiply by itself as long as the new index `j` is not the same as the old index `i`.

This works, but is not very optimised for time due to the double `for` loop resulting in $O(n^2)$ time complexity. The space complexity is $O(n)$ as we only store what we have.

## The Optimal Way

Thanks to [Neetcode](https://www.youtube.com/watch?v=bNvIQI2wAjk) they use a process of using a *postfix* and *prefix*. This is sort of how the algorithm works:

### Prefix
For each element, calculate the prefix by multiplying itself to the element calculated previously, if there is no element previous, assume it to be 1:
```go
in := []int{1,2,3,4}
prefix := make([]int, len(in))

for i, n := range in {
	prev := 1
	if i != 0{
		prev = prefix[i-1]
	}

	prefix[i] = prev * n
}

// prefix: {1, 2, 6, 24}
```

### Postfix
The postfix is the reverse of the prefix, so starting from the back, calculate the product by multiplying to the one on the right, if you are on the far right, assume the previous number is 1 again:
```go
in := []int{1,2,3,4}
postfix := make([]int, len(in))

for i := len(in) - 1; i >= 0; i-- {
	prev := 1
	if i != len(in) - 1 {
		prev = postfix[i + 1]
	}

	postfix[i] = prev * in[i]
}

// postfix: {24, 24, 12, 4}
```

### Combining
Finally you combine the *postfix* and *prefix* by taking the *prefix* before the value, and the *postfix* after:
```go
in := []int{1,2,3,4}
prefix := []int{1,2,6,24}
postfix := []int{24,24,12,4}
res := make([]int, len(in))

for i := range in {
	pre := 1
	if i != 0 {
		pre = prefix[i-1]
	}

	post := 1
	if i != len(in) - 1 {
		post = postfix[i+1]
	}

	res[i] = pre*post
}

// res: {24, 12, 8, 6}
```

So let's put this all together in a function:
```go
func PrefixPostfixV1(nums []int) []int {
	prefix := make([]int, len(nums))
	for i, n := range nums {
		prev := 1
		if i != 0 {
			prev = prefix[i-1]
		}
		prefix[i] = prev * n
	}

	postfix := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		prev := 1
		if i != len(nums)-1 {
			prev = postfix[i+1]
		}

		postfix[i] = prev * nums[i]
	}

	res := make([]int, len(nums))
	for i := range nums {
		pre, post := 1, 1
		if i != 0 {
			pre = prefix[i-1]
		}
		if i != len(nums)-1 {
			post = postfix[i+1]
		}

		res[i] = pre * post
	}

	return res
}
```

This works well, and the time complexity is now down to $O(n)$, but we are using so much space up with `prefix`, and `postfix` (`res` does not count towards the space complexity in this problem). We can optimise the space

## PostfixPrefixV2

So how do we combine the two? Well we just calculate the *postfix* and *prefix* in place of our output array:
```go
func PostfixPrefixV2(nums []int) []int {
	res := make([]int, len(nums))

	// Prefix pass
	prefix := 1
	for i, n := range nums {
		res[i] = prefix
		prefix *= n
	}

	// Postfix pass
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res

}
```

This is very similar to how [Neetcode](https://www.youtube.com/watch?v=bNvIQI2wAjk) does his:
```py
class Solution:
	def productExceptSelf(self, nums: List[int]) -> List[int]:
		res = [1] * (len(nums))

		prefix = 1
		for i in range(len(nums)):
			res[i] = prefix
			prefix *= nums[i]

		postfix = 1
		for i in range(len(nums) - 1, -1, -1):
			res[i] *= postfix
			postfix *= nums[i]

		return res
```

> I pinky promise I didn't see the solution until I wrote it just now, I just watched his explanation and implemented it.

## Conclusion

So what did we learn? Well this was one of those questions where I feel like it was a bit of a trick with the *postfix* and *prefix* technique, but one to remember if you are doing interviews. Otherwise the idea of going through a brute-force approach and then moving on to the more optimal approach could be good. That's all lol.
