package homework

func average(input [15]float32) (result float32) {
	sum := float32(0)
	count := 0
	for _, v := range input {
		if v != 0 {
			sum += v
			count += 1
		}
	}
	result = sum / float32(count)
	return
}
