package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers1, numbers2 []int) []int {

	sums := []int{0, 0}

	for _, number := range numbers1 {
		sums[0] += number
	}

	for _, number := range numbers2 {
		sums[1] += number
	}

	return sums
}

func SumAllTail(numbers1, numbers2 []int) []int {
	sums := []int{0, 0}

	for i := 1; i < len(numbers1); i++ {
		sums[0] += numbers1[i]
	}

	for i := 1; i < len(numbers2); i++ {
		sums[1] += numbers2[i]
	}

	return sums
}
