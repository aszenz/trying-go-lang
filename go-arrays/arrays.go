package arrays

func ArraySum(numbers []int) int {
	sum := 0
	for _, value := range(numbers) {
		sum += value
	}
	return sum
}
