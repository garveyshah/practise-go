package sln1

func CanJump(arr []uint) bool {
	for i := 0; i < len(arr); {
		i += int(arr[i])
		if i >= len(arr) {
			return false
		}
		if arr[i] == 0 && i < len(arr)-1 {
			return false
		} else {
			return true
		}
	}
	return true

}
