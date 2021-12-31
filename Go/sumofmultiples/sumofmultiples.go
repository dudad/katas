package sumofmultiples

func sumArithProgression(last int, progress int) int {
	last = ((last - 1) / progress) * progress
	return (last * (last/progress + 1)) / 2
}

func Multiple3And5(number int) int {
	if number < 4 {
		return 0
	}

	return sumArithProgression(number, 3) + sumArithProgression(number, 5) - sumArithProgression(number, 15)
}
