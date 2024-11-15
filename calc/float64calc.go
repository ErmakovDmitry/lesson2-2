package calc

func Add64(numbers ...float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Sub64(minuend float64, subtrahends ...float64) float64 {
	var result float64 = minuend

	for _, subtrahend := range subtrahends {
		result -= subtrahend
	}

	return result
}
