package arrays_slices

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll_longway(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	sums := []int{}
	for _, elem := range numbersToSum {
		if len(elem) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(elem[1:]))
		}
	}
	return sums
}
