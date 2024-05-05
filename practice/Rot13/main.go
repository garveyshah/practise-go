package main

import "fmt"

func main() {
	g := "My name is Ouma Godwin"
	fmt.Println(rot13(g))
	
}

func rot13(s string) string {
	var result []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >='A'  && c <= 'Z':
			result = append(result, (c-'A'+13)%26+'A')
		case c >=  'a' && c <= 'z':
			result = append(result, (c-'a'+13)%26+'a')
		default:
			result = append(result, c)
		}
	}
	return string(result)
}

