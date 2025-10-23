package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var response []int

	for _, numbers := range numbersToSum {
		response = append(response, Sum(numbers))
	}

	return response
}

func SumAllTails(numbersToSum ...[]int) []int {
	var response []int

	for _, numbers := range numbersToSum {
		if len(numbers) >= 1 {
			response = append(response, Sum(numbers[1:]))
		} else {
			response = append(response, 0)
		}
	}

	return response
}
