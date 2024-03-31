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

// Parse string list to substrings
//
// e.g. if size is 2:
//
//	["a" "b" "c"]         -> ["a" "b"] ["c"]
//	["a" "b" "c" "d"]     -> ["a" "b"] ["c" "d"]
//	["a" "b" "c" "d" "e"] -> ["a" "b"] ["c" "d"] ["e"]
func SplitStrings(array []string, size int) [][]string {
	result := [][]string{}
	start := 0
	for i := 0; i < (len(array)+size-1)/size; i++ {
		result = append(result, array[start:min(len(array), start+size)])
		start += size
	}
	return result
}
