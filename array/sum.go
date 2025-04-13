package array

func Sum(numbers []int) int {
	var sum int

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	// range allow you to loop through an array
	for _, number := range numbers {
		sum += number
	}

	return sum
}
