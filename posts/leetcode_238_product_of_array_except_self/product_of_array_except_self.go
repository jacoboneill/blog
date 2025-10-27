package main

func Cheating(nums []int) []int {
	res := make([]int, len(nums))

	// Calcualte the product, if there is a zero ignore it but store the index
	product := 1
	zero_index := -1
	for i, n := range nums {
		if n == 0 {
			zero_index = i
		} else {
			product *= n
		}
	}

	// If there is a zero, then all values are zero except the index with zero itself which is the total product
	if zero_index != -1 {
		res[zero_index] = product
		return res
	}

	// Else divide the product by the current number
	for i, n := range nums {
		res[i] = product / n
	}
	return res
}

func BruteForce(nums []int) []int {
	// Create an slice of the l of the input and set them all to 1
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	// For each num, get the product of everything not including the current index
	for i := range nums {
		for j, n := range nums {
			if i != j {
				res[i] *= n
			}
		}
	}

	return res
}

func PostfixPrefixV1(nums []int) []int {
	// Get prefix
	prefix := make([]int, len(nums))
	for i, n := range nums {
		prev := 1
		if i != 0 {
			prev = prefix[i-1]
		}

		prefix[i] = prev * n
	}

	// Get postfix
	postfix := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		prev := 1
		if i != len(nums)-1 {
			prev = postfix[i+1]
		}

		postfix[i] = prev * nums[i]
	}

	// Combine
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
