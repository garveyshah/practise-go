package piscine

func TrimAtoi(s string) int {
	var negative bool = false
	var empty bool = true
	var result int = 0
	for _, value := range s {
		if empty && !negative && value == '-' {
			negative = true
		} else if value >= '0' && value <= '9' {
			result *= 10
			result += int(value - '0')
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if negative {
			return -result
		} else {
			return result
		}
	}
}
