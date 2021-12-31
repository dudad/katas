package uniquenumber

func FindUniq(arr []float32) float32 {
	numbers := make(map[float32]int32)
	for _, el := range arr[0:3] {
		numbers[el] += 1
	}
	if len(numbers) > 1 {
		for number, cnt := range numbers {
			if cnt == 1 {
				return number
			}
		}
	}

	for _, el := range arr[3:] {
		if el != arr[0] {
			return el
		}
	}
	return 0
}
