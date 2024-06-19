package compare

func IsType[T any](i interface{}) bool {
	switch i.(type) {
	case T:
		return true
	default:
		return false
	}
}
