package piscine

func isAlphaNumeric(a rune) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

func Capitalize(s string) string {
	arr := []rune(s)

	shouldCapitalize := true

	for i := 0; i < len(arr); i++ {
		if isAlphaNumeric(arr[i]) && shouldCapitalize {
			if arr[i] >= 'a' && arr[i] <= 'z' {
				arr[i] = arr[i] - 'a' + 'A'
			}
			shouldCapitalize = false
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] - 'A' + 'a'
		} else if !isAlphaNumeric(arr[i]) {
			shouldCapitalize = true
		}
	}
	return string(arr)
}
