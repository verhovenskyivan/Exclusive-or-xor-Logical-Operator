package kata

func Xor(a, b bool) bool {
	if a == b {
		return false
	} else if a || b == true {
		return true
	} else {
		return false
	}
}
