package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	lengthOfNumbers := len(slices)
	for i := 0; i < lengthOfNumbers; i++ {
		sums = append(sums, Sum(slices[i]))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))

		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
