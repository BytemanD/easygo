package math

func SumInt(number ...int) int {
	sum := 0
	for _, num := range number {
		sum += num
	}
	return sum
}
func SumIntArray(numbers []int) int {
	return SumInt(numbers...)
}

func SumFloat32(number ...float32) float32 {
	sum := float32(0)
	for _, num := range number {
		sum += num
	}
	return sum
}
func SumFloat32Array(numbers []float32) float32 {
	return SumFloat32(numbers...)
}
func SumFloat64(number ...float64) float64 {
	sum := float64(0)
	for _, num := range number {
		sum += num
	}
	return sum
}
func SumFloat64Array(numbers []float64) float64 {
	return SumFloat64(numbers...)
}
