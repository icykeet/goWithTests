package arrandslice

func Sum(nums []int) int {
	res := 0

	for _, num := range nums {
		res += num
	}

	return res
}

func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, nums := range numberToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, num := range nums {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
