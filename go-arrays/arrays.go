package arrays

// ArraySum sums up the values of array
func ArraySum(numbers []int) int {
	sum := 0
	for _, value := range(numbers) {
		sum += value
	}
	return sum
}

// SlicesSum sums up multiple slices and returns an array with their sum
func SlicesSum(slices [][]int) []int {
	var sums []int
	for _, arr := range(slices) {
		sums = append(sums, ArraySum(arr))
	}
	return sums
}
