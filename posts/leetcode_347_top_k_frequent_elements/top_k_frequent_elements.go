package top_k_frequent_elements

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
		for n, c := range m {
			if c > max_num.Count {
				max_num = N{n, c}
			}
		}
		delete(m, max_num.Number)
		res[i] = max_num.Number
	}

	return res
}

func BucketSort(nums []int, k int) []int {
	// Get Count Map
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	// Convert Map to Bucket Sort Array
	freq := make([][]int, len(nums)+1)
	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	// Get Largest From Bucket Sorted Array
	var res []int
	for i := len(freq) - 1; i >= 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
