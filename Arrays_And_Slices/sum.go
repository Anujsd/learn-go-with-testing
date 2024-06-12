package arrays

func GetSum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func SumAll(numbers ...[]int) []int {

	var sums []int

	for _, numbers := range numbers {
		sums = append(sums, GetSum(numbers))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {

	var sums []int

	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, GetSum(numbers[1:]))
		}
	}

	return sums
}
