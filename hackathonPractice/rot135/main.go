package main

import "fmt"

func ROT135(input string) string {
	sliceByte := []rune{}

	for _, char := range input {
		switch {
		case char >= 'A' && char <= 'Z':
			char = (char-'A'+13)%26 + 'A'
			sliceByte = append(sliceByte, char)
		case char >= 'a' && char <= 'z':
			char = (char-'a'+13)%26 + 'a'
			sliceByte = append(sliceByte, char)
		case char >= '0' && char <= '9':
			char = (char-'0'+5)%10 + '0'
			sliceByte = append(sliceByte, char)
		default:
			sliceByte = append(sliceByte, char)
		}
	}
	return string(sliceByte)
}

func main() {
  str := "The quick brown fox jumps over the 2 lazy dogs"
  str1 := "Gur dhvpx oebja sbk whzcf bire gur 7 ynml qbtf"
	fmt.Println(ROT135(str))
  fmt.Println(ROT135(str1))
}
