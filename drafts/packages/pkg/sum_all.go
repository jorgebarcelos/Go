package assistant

func SumAll(x ...int) int {
	result := 0
	for _, value := range x {
		result += value
	}
	return result
}