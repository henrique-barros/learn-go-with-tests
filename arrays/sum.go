package arrays

func Sum(numbers []int) int {
	add := func(acc, cur int) int {
		return acc + cur
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	addAll := func(acc, cur []int) []int { return append(acc, Sum(cur)) }
	return Reduce(numbersToSum, addAll, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, cur []int) []int {
		if len(cur) >= 1 {
			return append(acc, Sum(cur[1:]))
		} else {
			return append(acc, 0)
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A, B any](collection []A, fn func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, item := range collection {
		result = fn(result, item)
	}
	return result
}
