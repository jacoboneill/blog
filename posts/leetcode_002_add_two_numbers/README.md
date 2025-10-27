---
author: Jacob O'Neill
date: 03/10/2025
tags: ["leetcode"]
urls: { "leetcode": "https://leetcode.com/problems/add-two-numbers", "github": "https://github.com/jacoboneill/blog/blob/main/posts/leetcode_002_add_two_numbers/add_two_numbers.go"}
---

# Leetcode 001 Two Sum

## The Problem
> You are given two non-empty linked-lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked-list.
> 
> You may assume the two numbers do not contain any leading zero, except the number 0 itself.

This is the explanation that *Leetcode* gives, but I'll be honest it took me three iterations and talking to Claude to fully understand what this question was even asking so, let me posit this in my own words:

> You are given two non-empty linked-lists representing two non-negative integers where each node is an integer. The numbers are stored in reversed order, i.e. 1 &rarr; 2 &rarr; 3 = 321.
> 
> Given these two linked-lists: sum them, and return the number as a linked-list in reversed-order.
> 
> You may assume the two numbers do not contain any leading zero, except the number 0 itself.

### Example
> l1: 1 &rarr; 2 &rarr; 3 = 321
>
> l2: 4 &rarr; 5 &rarr; 6 = 654
>
> sum: 975 = 5 &rarr; 7 &rarr; 9

## The Naive Approach
The naive approach that I took initially is to decompose the problem and think of it as a human. This is how I thought of it in my head:
1. Traverse `l1` in reverse and compute the integer from the values
2. Traverse `l2` in reverse and compute the integer from the values
3. Sum the two integers
4. Generate a linked-list in the reverse of the sum
5. Return the first node of the linked-list

I decomposed this into two problems: 
- Generate an integer based on a linked-list in reversed order.
- Generate a linked-list based on an integer in reversed order.

### Generate an Integer Based on a Linked-List in Reversed Order
To do this, I kept track of the power of 10 used, incrementing each time, and summed the values while traversing the linked-list
```py
def traverse_list_reversed(list_node: ListNode) -> int:
    # Initiate values
    power_of_ten: int = 0
    result: int = 0

    # Loop
    while list_node is not None:
        result += pow(10, power_of_ten) * list_node.val
        power_of_ten += 1
        list_node = list_node.next

    return result
```

So for a linked-list that looks like
> 1 &rarr; 2 &rarr; 3

The function would look like this for each iteration:
```
Iteration 1:
result = 0 + (pow(10, 0) * 1)
result = 0 + (1 * 1)
result = 0 + 1
result = 1

Iteration 2:
result = 1 + (pow(10, 1) * 2)
result = 1 + (10 * 2)
result = 1 + 20
result = 21

Iteration 3:
result = 21 + (pow(10, 2) * 3)
result = 21 + (100 * 3)
result = 21 + 300
result = 321
```

### Generate a Linked-List Based on an Integer in Reversed Order
To do this, I generated a linked-list using the modulo of 10, and then floored it by ten, sort of replicating a stack with an integer.
```py
def generate_linked_list(n: int) -> ListNode:
    first = ListNode(val = n % 10)
    previous = first
    n //= 10

    while n != 0:
        previous.next = ListNode(val=n%10)
        previous = previous.next
        n //= 10

    return first
```

So for an integer of `321`, the function would look like this for each iteration:
```
Setup:
first = (321 % 10) -> None
first = (1) -> None

previous = (1) -> None

n = 321 // 10
n = 32

Iteration 1:
previous.next ((1) -> None) = (32 % 10)
previous.next ((1) -> None) = (2)
previous.next ((1) -> (2))
previous = previous.next ((2) -> None)
n = 32 // 10
n = 3

Iteration 2:
previous.next ((2) -> None) = (3 % 10)
previous.next ((2) -> None) = (3)
previous.next ((2) -> (3))
previous = previous.next ((3) -> None)
n = 3 // 10
n = 0

Result:
return (1) -> (2) -> (3)
```

### The problem with this approach

There are unfortunately a few problems with this approach. The first one being is that it is really slow.
You have to remember it has to loop through each linked-list and then generate another list. This means that you are going to be in the $O(3n)$ which when we normally look at time complexity would be simplified to $O(n)$ but in this case there is an actual $O(n)$ solution that we can do instead

The other (pretty major) issue is our conversion of linked-list to integer. In Python this is not too bad as an "integer" can change interchangeably between an `int` and a `long` without the developer even knowing [\[1\]](https://stackoverflow.com/questions/7604966/maximum-and-minimum-values-for-ints). But when I was porting this answer over to Go, that was when I realised this answer would not work (Go's maximum integer is $1.8*10^20$ so therefore the maximum length of a linked-list is 20 before things start to break)[\[2\]](https://go.dev/tour/basics/11)

So what is the *"correct"* way to solve this problem

## The Optimal Way

The optimal way you would have learned in primary school (or elementary school for my readers from over the pond). Calculate each column and make sure to account for the carry.

So in our previous example of:
> l1: 1 &rarr; 2 &rarr; 3
> 
> l2: 4 &rarr; 5 &rarr; 6

We would do:
```py
def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]: 
    dummy = ListNode()
    current = dummy
    carry = 0

    while l1 is not None or l2 is not None or carry != 0:
        v1 = l1.val if l1 else 0
        v2 = l2.val if l2 else 0

        value = v1 + v2 + carry
        carry = value // 10
        value %= 10

        current.next = ListNode(value)
        current = current.next
        
        l1 = l1.next if l1 else None
        l2 = l2.next if l2 else None

    return dummy.next
```

This was my second attempt at the problem after getting the hint of *"treat it like adding two numbers in school"* (thanks [Neetcode](https://www.youtube.com/watch?v=wgFPrzTjm7s) for the hint). After watching his video I found my solution was essentially the same.

## My Port to Go

As I am trying to move away from Python and more into Go, with each of these challenges I will be doing a port to Go.

Here is my port to Go:
```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        val := carry

        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }

        carry = val / 10

        curr.Next = &ListNode{val % 10, nil}
        curr = curr.Next
    }

    return dummy.Next
}
```

The main differences between the Go and Python versions are the checking for `l1` and `l2` being `None`/`nil` or not. If you move around the program flow like I did then it works quite cleanly but trust that I had a very messy first version. Apart from that it seems to mainly be just a syntax difference.

For the [test file](https://github.com/jacoboneill/leetcode/blob/main/002_add_two_numbers/add_two_numbers_test.go) it was actually quite interesting. I implemented `Atoll` "**A**rray **to** **l**inked-**l**ist" and `Lltoa` "**L**inked-**l**ist **to** **a**rray" utility functions so I can write the cases out as integer arrays, like Leetcode does on their website.

```go
func Atoll(arr []int) *ListNode {
	first := &ListNode{arr[0], nil}
	previous := first
	for _, n := range arr[1:] {
		previous.Next = &ListNode{n, nil}
		previous = previous.Next
	}

	return first
}

func Lltoa(l *ListNode) []int {
	var res []int
	for l.Next != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return append(res, l.Val)
}
```

I also created a `AssertLinkedList` function so I can just call that in the evaluation.

```go
func AssertLinkedList(l1 *ListNode, l2 *ListNode) bool {
	for l1.Next != nil || l2.Next != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1.Next == nil && l2.Next == nil
}
```

## Conclusion

Overall, this was an interesting and more difficult task than [001 Two Sum](../001_two_sum/README.md). It taught me a lot about maneuvering around linked-lists and how to think a bit smarter about an optimal solution rather than taking a brute force approach. It also taught me how Python can hide certain things (such as integer memory size) that can make the implementation perfectly fine there but anywhere else and it will not work.
