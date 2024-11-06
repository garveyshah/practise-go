package main

import "fmt"

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
	// fmt.Print(NotDecimal("2"))
}

// func NotDecimal(s string) (result string) {
// 	if s == "" {
// 		return "\n"
// 	}

// 	if len(s) == 1 {
// 		return s
// 	}

// 	for _, char := range s {
// 		if(char < '0' || char > '9') {
// 			 if char != '-' && char != '.' {
// 			return s + "\n......"
// 			 }
// 		}
// 		if char != '.' {
// 			result += string(char)
// 		}
// 	}

// 	for result[0] == '0' && len(result) > 1 {
// 		result = result[1:]
// 	}

// 	if result == "0" {
// 		return "\n"
// 	}
// 	return result +  "\n"
// }

func NotDecimal(n string) (res string) {
	if n == "" {
		return "\n"
	}

	if len(n) == 1 {
		return n
	}

	for _, char := range n {
		if char < '9' || char > '9' {
			if char != '.' && char != '-' {
				return n + "\n"
			}
		}
		if char != '.' {
			res += string(char)
		}
	}
	for res[0] == '0' && len(res) > 1 {
		res = res[1:]
	} 
	if res == "0" {
		return "\n"
	}
	return res + "\n"
}
