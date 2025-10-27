---
author: Jacob O'Neill
date: 12/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/group-anagrams", "github": "https://github.com/jacoboneill/blog/blob/main/posts/leetcode_049_group_anagrams/group_anagrams.go"}
---

# Leetcode 049 Group Anagrams

A continuation on [Leetcode's 242 Valid Anagram](../242_valid_anagram/README.md). This problem has you returning a 2D array of strings where all of the strings that are anagrams are grouped together

> Given an array of `strs`, group the [anagrams](https://gcide.gnu.org.ua/?q=Anagram&define=1) together. You can return the answer in any **order**

**Example 1:**
```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
    - There is no string in strs that can be rearranged to form "bat"
    - The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
    - The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
```

**Example 2:**
```
Input: strs = [""]
Output: [[""]]
```

**Example 3:**
```
Input: strs = ["a"]
Output: [["a"]]
```

## My First Approach

My first instinct was *"Ah, we need to group things together in regards to a common theme, this sounds like a job for a hash map"*.

So if the common theme is an anagram, and we learnt from [valid anagram](../242_valid_anagram/README.md) that the way to validate an anagram is to get the count of all characters.

In [valid anagram](../242_valid_anagram/README.md) we used a `hash set` to get a count of all of the characters, but in Go (and many other languages), a `hash set` can not be used as a key for a `hash map` because it is not a comparable data type.

The solution? It feels a bit cheat-y but I'm going to sort the string, this will therefore give us a key that is a representation of the count of the characters (because once the strings are sorted, they will be the same):

```
Input: ["eat", "tea"]

eat &rarr; aet
tea &rarr; aet

Output: {"aet": ["eat", "tea"]}
```

We can then output an array of all the values in the `hashmap` by doing this.

Here was my implementation of this:
```go
import "slices"

func Sorting(strs []string) [][]string {
    // Utility function to sort string into alphabetical order.
    sortString := func(s string) string {
        b := []byte(s)
        slices.Sort(b)
        return string(b)
    }

    // Sort strings into groups of anagrams
    m := make(map[string][]string)

    for _, s := range strs {
        sorted := sortString(s)
        m[sorted] = append(m[sorted], s)
    }

    // Extract groups from hash map
    res := make([][]string, 0, len(strs))
    for _, v := range m {
        res = append(res, v)
    }

    return res
}
```

A couple notes on the program:
- I created a utility `sortString` function to make the code more readable
- Instead of using the `sorting` builtin library, I am using the `slices` library. This is because, firstly, there is no builtin string sorting in Go, so you have to sort an array of either `rune`s or `byte`s and then convert back to a `string`. Secondly, I am using the `slices` library to do the sorting because the Go LSP suggested this as a more modern approach.
O
This results in an $O(m \cdot n \cdot \log{n})$ time complexity and a $O(m \cdot n)$ space complexity, where $m$ is the number of strings, and $n$ is the average length of a string.

## Is There a Better Way?

After doing some research I found that my solution is surprisingly the most common solution! However, there is another solution that, depending on the language, could be slightly more optimal. This is because of the constraints of the problem:

> `strs[i]` consists of lowercase English letters.

This means that the `hashing` of the count could be constructed differently. [Neetcode](https://www.youtube.com/watch?v=vzdNOK2oB2E) suggests using an array of `int`s, where the index is the position of the letter in the alphabet, and the value is the count of that letter, requiring a 26 length `int` array:

```
Input: ["eat", "tea"]

eat &rarr; [1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0]
tea &rarr; [1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0]

Output: {[1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0]: ["eat", "tea"]}
```

This means we don't need to have the complexity of sorting the strings. Meaning that our time and space complexity should both go to $O(m \cdot n)$.

```go
func MapArray(strs []string) [][]string {
    // Sort strings into groups of anagrams
    m := make(map[[26]int][]string)
    for _, s := range strs {
        var count [26]int

        for _, c := range s {
            count[c-'a']++
        }

        m[count] = append(m[count], s)
    }

    // Extract groups from hash map
    res := make([][]string, 0, len(strs))
    for _, v := range m {
        res = append(res, v)
    }

    return res
}
```

This is a slightly more difficult concept to keep in your head compared to the sorting, but removing the need for the *sorting* utility function makes the algorithm look a lot cleaner.

## So Which Is Better?

So to test the performance of both, I wrote a quick benchmark test in the `group_anagrams_test.go` file:

```go
func BenchmarkGroupAnagrams(b *testing.B) {
    for _, function := range functions {
        b.Rune(function.Name, func(b *testing.B) {
            for _, tt := range cases {
                for range b.N {
                    function.F(tt.Input)
                }
            }
        })
    }
}
```

These are the results of that:
```
> go test -bench=.

...

BenchmarkGroupAnagrams/Sorting-10                2656352               430.8 ns/op
BenchmarkGroupAnagrams/Map_Array-10              2444944               495.9 ns/op
PASS
ok      049_group_anagrams      3.729s
```

So how does that work? We found earlier that the `MapArray` function has a better time complexity than the `Sorting` function? Well, this is why we benchmark.

If we look at our test cases (and most real world cases), most of our words are below 10 characters each. This means when we hash using the `Sorting` algorithm, our keys are typically less than 40 bytes (each `rune` is 4 bytes). This is dissimilar to our `MapArray` algorithm that will always have a key memory size of 104 bytes (26 `int`'s where each `int` is 4 bytes long).

So our `MapArray` will eventually be better, but not until we hit something like 100 characters. So while the `MapArray` is theoretically more time and space efficient, in the real world, with Go's very efficient sorting algorithms, the `Sorting` array is more performant.

On Leetcode, our statistics fluctuate, but typically both the `Sorting` and `MapArray` solutions score the same.
