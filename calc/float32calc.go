package calc

func Add32(numbers ...float32) float32 {
	var sum float32 = 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Sub32(minuend float32, subtrahends ...float32) float32 {
	var result float32 = minuend

	for _, subtrahend := range subtrahends {
		result -= subtrahend
	}

	return result
}
