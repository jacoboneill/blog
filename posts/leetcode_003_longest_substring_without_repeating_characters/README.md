---
author: Jacob O'Neill
date: 08/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/longest-substring-without-repeating-characters", "github": "https://github.com/jacoboneill/blog/blob/main/posts/leetcode_003_longest_substring_without_repeating_characters/longest_substring_without_repeating_characters.go"}
---

# Leetcode 003 Longest Substring Without Repeating Characters

## The Problem

Given a string `s`, find the length of the **longest** substring without duplicate characters

### Examples

**Example 1:**
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
```

**Example 2:**
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. Notice that the answer must be a substring, "pwke" is a subsequence and not a substring
```

## Naive Approach

### The Beginning

For the first time, I have decided to try and write this entire problem in Go, rather than using Python first and then parsing that into Go.

What I did first was I setup my [`longest_substring_without_repeating_characters_test.go`](https://github.com/jacoboneill/leetcode/blob/main/003_longest_substring_without_repeating_characters/longest_substring_without_repeating_characters_test.go) file:

```go
// longest_substring_without_repeating_characters_test.go
package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	functions := []struct {
		Name     string
		Function func(string) int
	}{
		{"Naive", Naive},
	}

	cases := []struct {
		Input    string
		Expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, f := range functions {
		for _, tt := range cases {
			got := f.Function(tt.Input)
			if got != tt.Expected {
				t.Errorf("%s: Input: %s, Output: %d, Expected: %d", f.Name, tt.Input, got, tt.Expected)
			}
		}
	}
}
```

I copied all of the cases in the description, and am now ready to start developing:

```go
// longest_substring_without_repeating_characters.go

package longest_substring_without_repeating_characters

func Naive(s string) int {
	return 0
}
```

### My Thought Proccess

This is how I imagine the question could be solved. I am thinking of the *brute-force* approach: calculate every possible combination and then return the max of the length of strings

### My First Attempt

Okay, so I ended up doing it a slightly different approach, as we don't need to keep track of what was the longest substring just the length.

As well as this, Go doesn't have a `set` type like Python does, but it seems the standard way to solve this is to use a `map[type]bool` and just check using the syntax:
```go
if _, ok := m[t]; ok {
	// t is a key in the map m so therefore t exists
}
```

The other hurdle I had to deal with was in the *"pwwkew"* case, where the max substring is "wke" or "kew", and you have to include a check for the final character

Here is what I came up with:
```go
// longest_substring_without_repeating_characters.go

package longest_substring_without_repeating_characters

func Naive(s string) int {
	maxLength := 0
	start := 0

	seen := make(map[rune]bool)
	for i, c := range s {
		if _, ok := seen[c]; ok {
			if i-start > maxLength {
				maxLength = i - start
			}
			start = i + 1
			seen = make(map[rune]bool)
		} else if i == len(s)-1 && len(s)-start > maxLength {
			maxLength = len(s) - start
		} else {
			seen[c] = true
		}
	}
	return maxLength
}
```

And this passed all of my cases. Woohoo!

Unfortunatley, when I put this into Leetcode, it did not pass all of the cases:
```
Input
s = "aab"

Output
1

Expected
2
```

... bugger

### Let's Try This Again...

I updated the test file so that means I can have functions that aren't meant to pass (just to keep a record of what I did without it reporting that I have failing tests):
```go
// longest_substring_without_repeating_characters_test.go

package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	functions := []struct {
		Name        string
		Function    func(string) int
		MeantToPass bool
	}{
		{"Naive", Naive, false},
		{"Let's Try This Again", NaiveLetsTryThisAgain, true},
	}

	cases := []struct {
		Input    string
		Expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, f := range functions {
		if !f.MeantToPass {
			continue
		}

		for _, tt := range cases {
			got := f.Function(tt.Input)
			if got != tt.Expected {
				t.Errorf("%s: Input: %s, Output: %d, Expected: %d", f.Name, tt.Input, got, tt.Expected)
			}
		}
	}
}
```

This basically just means that if the `f.MeantToPass` flag is negative, then just don't borther running the test.

So what did I do wrong. Let's do a write up of the failing test "aab"
```
s = "aab", expected = 2

0:
	maxLength = 0
	start = 0
	seen = {}

1: i=0, c="a"
	c is not in seen
	i != len(s) - 1
	seen = {"c": true}

2: i=1, c="a"
	c is in seen
	i - start = 1, which is larger than maxLength
	maxLength = 1
	start = 2
	seen = {}

3: i=2, c="b"
	c not in seen
	i == len(s) - 1 but len(s) - 2 is not larger than maxLength <-- (THIS IS THE PROBLEM)

return maxLength = 1
```

Aha! We aren't keeping track of our start very well. After a bit of tinkering, I found if we update:
```go
start = i + 1
seen = make(map[rune]bool)
```

to
```go
start = i
seen = map[rune]bool{c: true}
```

Then we get a few more answers right. Okay it's not perfect but at least we are getting somewhere... Let's try doing the runthrough again with our new algorithm with a new case that is not working: *"dvdf"*

```
s = "dvdf", expected = 3

0:
	maxLength = 0
	start = 0
	seen = {}

1: i=0, c="d"
	c is not in seen
	i != len(s) - 1 (3)
	seen = {"d": true}

2: i=1, c="v"
	c is not in seen
	i != len(s) - 1 (3)
	seen = {"d": true, "v": true}

3: i=2, c="d"
	c is in seen
	i (2) - start (0) (2) > maxLength(0)
	maxLength = 0
	start = 2
	seen = {"d": true}

4: i=3, c="f"
	c is not in seen
	i == len(s) - 1 (3) and len(s) (4) - start (2) (2) is not > maxLength

return maxLength (2)
```

Okay, I see the problem after writing this. We are not going back to "v" to check to see if there are any substrings there. Which means unfortunatley we are going to need to do a double `for` loop to complete this.

```go
// longest_substring_without_repeating_characters.go

package longest_substring_without_repeating_characters

func NaiveLetsTryThisAgain(s string) int {
	maxLength := 0

	for i := range s {
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			char := s[j]
			if seen[char] {
				break
			}
			seen[char] = true
			currentLength := j - i + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
```

This passes all of our local tests. Let's hope it passes the tests on Leetcode...

WOOHOO! This passes all of the tests on Leetcode. The bad news is the stats:
> Runtime: 293ms (Beats 5.7%)
> 
> Memory: 8.87MB (Beats 12.3%)

One more optimisation I can think of is that if the `maxLength` is bigger than the difference between `i` and `j` then there isn't any point looking anymore because there could never be a bigger combo. Let's see what the stats say for that

```go
// longest_substring_without_repeating_characters.go

package longest_substring_without_repeating_characters

func NaiveLetsTryThisAgain(s string) int {
	maxLength := 0

	for i := range s {
		if len(s)-i <= maxLength {
			break
		}
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			char := s[j]
			if seen[char] {
				break
			}
			seen[char] = true
			currentLength := j - i + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
```

And...
> Runtime: 271ms (Beats 7.6%)
>
> Memory: 9.05MB (Beats 10.6%)

Okay, the runtime is slightly better but the memory is worse. I think this may be a discrepancy in Leetcode's severs for caching? (Run the same code multiple times and you will get different answers every time but always around the area you were.)

Think it's time to do some research...

## One Research Later

So after watching [Neetcode](https://www.youtube.com/watch?v=wiGpQwVHdE0)'s video, they suggest to use the *Sliding Window* Technique. What does this mean exactly?

We start with a range of one item, we check if there is a duplicate, if not then we increase the range until there is a duplicate. Then we move the window across until either we are at the end or we can increase it again. This means that we can have a time and space complexity of $O(n)$ instead of our original $O(n^2)$.

So for *"abcabcbb"* we would do the following:

```
[a]bcabcbb | Max Length: 1, Seen: {a}

[ab]cabcbb | Max Length: 2, Seen: {a, b}

[abc]abcbb | Max Length: 3, Seen: {a, b, c}

[abca]bcbb | X NOT ALLOWED X

a[bca]bcbb | Max Length: 3, Seen: {b, c, a}

a[bcab]cbb | X NOT ALLOWED X

ab[cab]cbb | Max Length: 3, Seen: {c, a, b}

ab[cabc]bb | X NOT ALLOWED X

abc[abc]bb | Max Length: 3, Seen: {a, b, c}

abc[abcb]b | X NOT ALLOWED X

abca[bcb]b | X NOT ALLOWED X

abcab[cb]b | Max Length: 3, Seen: {c, b}

abcab[cbb] | X NOT ALLOWED X

abcabc[bb] | X NOT ALLOWED X

abcabcb[b] | Max Length: 3, Seen: {b}
```

Or for another one of our examples *"dvdf"*:

```
[d]vdf | Max Length: 1, Seen: {d}

[dv]df | Max Length: 2, Seen: {d, v}

[dvd]f | X NOT ALLOWED X

d[vdf] | Max Length: 3, Seen: {v, d, f}

dv[df] | Max Length: 3, Seen: {d, f}

dvd[f] | Max Length: 3, Seen: {f}
```

Now there is an obvious optimisation, which is to break out of the loop early if there is no more length to the string to make the combo bigger than the maximum length, i.e. we don't need to check the cases:
- `dv[df]`
- `dvd[f]`

Because we already know the maximum length was three.

Let's try and implement this in Go:

```go
func AfterSomeResearch(s string) int {
	maxLength := 0
	seen := make(map[byte]bool)

	start := 0

	for end := 0; end < len(s); end++ {
		for _, ok := seen[s[end]]; ok; _, ok = seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		seen[s[end]] = true
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}
```

I'll be honest, I still don't fully understand how the sliding window algorithm works, and the wonky `for` loops in Go that are actually `while` loops in any other language make my brain hurt, but these are the main things you need to know

1. Setup your static variables, `maxLength`, `seen`, and `start`
	1. You may have noticed that `seen` is now a `map[byte]bool` rather than a `map[rune]bool`. This is because when you index into a string in Go, you get returned the byte in ASCII for the character. I.e. `x := "abc"`, `x[0]` would be hex `61`. The reason I changed the data type was so that means we don't have to do the type conversion, because really there is no need to
2. Set the start and end pointer, the end pointer will loop through the entire string
3. Keep looping and moving the start pointer right by one until there are no more duplicates
4. The last if statement is basically just a `max` function, checking if we need to update the `maxLength`.

We can also add our pruning optimisation by changing:
```go
// Before
for end := 0; end < len(s); end++ {...}


// After
for end := 0; end < len(s) && len(s)-start > maxLength; end++ {...}
```

Which is basically saying instead of *"While `end` is smaller than the length of `s`"* it also says *"... and if the `start` is far enough away from the end that the window's range can be bigger than the current `maxLength`"*. Okay, I may have not explained that well but unfortunately for you, I am not going to explain it any better. We live in a day and age of AI, if you can't understand my rambling just plug this into ChatGPT and ask it to summarise.

## Conclusion

Overall this Leetcode problem has been the hardest one I have learned so far, taking me three days to wrap my head around, but here are all the things I learned

- How to debug a problem
- How to create `set`'s in Go using a `hash map` with a key of the value in a traditional `set` and the value being a `boolean`.
- How in Go, indexing into strings returns a `byte` value
- The sliding window technique (sort of, I still don't fully get it).

I think the issue is that I am doing these Leetcode problems in order, rather than learning one thing and then building on that. So from now on I am going to follow the [Neetcode](https://neetcode.io/roadmap)'s roadmap to learn DSA better. This means we will be going back a bit into easier stuff, but it's more concrete.

This means next time we will be learning about **arrays and hashing**, working on *Leetcode 217: Contains Duplicate*.
