package main

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	length := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		length++
	}
	result := make([]byte, length+len(sign))
	for i := length - 1; i >= 0; i-- {
		result[i+len(sign)] = byte('0' + n%10)
		n /= 10
	}
	copy(result, sign)

	return string(result)
}
