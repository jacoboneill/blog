---
author: Jacob O'Neill
date: 09/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/valid-anagram", "github": "https://github.com/jacoboneill/blog/blob/main/posts/242_valid_anagram/valid_anagram.go"}
---

# Leetcode 242 Valid Anagram

This problem has you making a function to test if two strings are an anagram of one another.

> Given two strings `s` and `t`, return `true` if `t` is an [anagram](https://gcide.gnu.org.ua/?q=Anagram&define=1) of `s`, and `false` otherwise.

**Example 1**:
```
Input: s = "anagram", t = "nagaram"
Output: true
```

**Example 2**:
```
Input: s = "rat", t = "car"
Output: false
```

## Naive solution

My initial solution without doing any research was the following:

1. Record the count of all characters in the `s` and store them in a `hashmap`
2. Takeaway the count of all the characters in `t`. If the count is `0` then delete the entry from the `hashmap`, if the character doesn't exist in the `hashmap` then return `false` early
3. Check the length of the `hashmap`. If there is anything left, return `false` otherwise return `true`.

Here is my implementation:
```go
func Naive(s, t string) bool {
    // Create the Hash Map
    m := make(map[rune]int)

    // Add all the characters from `s`
    for _, r := range s {
        if _, ok := m[r]; !ok {
            m[r] = 1
        } else {
            m[r]++
        }
    }

    // Remove all the characters from `t`, checking for early returns
    for _, r := range t {
        if _, ok := m[r]; !ok {
            return false
        } else if m[r] == 1 {
            delete(m, r)
        } else {
            m[r]--
        }
    }

    // Check if there are any records left, if there are then it should return false
    return len(m) == 0
}
```

This works, and passes all of our tests but doesn't feel very efficient. We have to check (at maximum) each string once. This does technically mean that the time and space complexities are $O(n)$ where $n$ is the length of `s` and `t`.

Looking at the results on Leetcode however, it seems that there must be a better solution as these are my statistics after submitting:

> Runtime: 6ms | Beats 33.1%
>
> Memory: 4.68 MB | Beats 99.48%

I think there must be a trick I am missing

## The Trick I Was Missing

So there was a sneaky trick that I was missing. If you check the top 1% of answers on Leetcode, they seem to use an `[26]int`. This is because, in the Leetcode specification:

> Constraints:
> 
> * `1 <= s.length, t.length <= 5 * 10^4`
> 
> * `s` and `t` consist of lowercase English letters

This means that we can remove the overhead of a `hashmap` and just use an array of 26 length and index into the order of the character and have the count be the value


```go
func TrickSolution(s, t string) bool {
    var arr [26]int

    for _, ch := range s {
        arr[ch-'a']++
    }

    for _, ch := range t {
        arr[ch-'a']--
    }

    for _, count := range arr {
        if count != 0 {
            return false
        }
    }

    return true
}
```

This still gives us a time and space complexity of $O(n)$ where $n$ is still the length of `s` and `t`. But because of Go's implementation, and reduced overhead of a `hashmap`, it gives us these statistics:

```
BenchmarkValidAnagram/Naive-10           151.7 ns/op
BenchmarkValidAnagram/Trick_Solution-10  26.46 ns/op
```

This shows that the `TrickSolution` is almost **5.75**x faster than the `Naive` solution!

## Conclusion

So what did we learn. Well the academically correct solution is probably the `hashmap` solution, and Leetcode even says at the bottom of the description:

> **Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

The trick solution of using an array to remove the overhead instead would soon diminish, the more characters you have. But it is also important to remember language implementations to know what the best solution is depending on what you are doing!
