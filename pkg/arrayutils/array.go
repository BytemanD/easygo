package arrayutils

// [min], max, [step]
func Range(args ...int) []int {
	var min, max, step int
	switch len(args) {
	case 0:
		min, max, step = 0, 0, 1
	case 1:
		min, max, step = 0, args[0], 1
	case 2:
		min, max, step = args[0], args[1], 1
	default:
		min, max, step = args[0], args[1], args[2]
	}

	items := []int{}
	for i := min; i < max; i += step {
		items = append(items, i)
	}
	return items
}
