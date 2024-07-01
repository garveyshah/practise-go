package main

import (
	"fmt"
	"os"
)

 // Mpa of integer values to their corresponding Roman numeral symbols
var romanNumerals = []struct {
	Value int
	Symbol string
}{
	{1000, "M"},
	{900,"CM"},
	{500,"D"},
	{400,"CD"},
	{100,"C"},
	{90,"XC"},
	{50,"L"},
	{40,"XL"},
	{10,"X"},
	{9,"IX"},
	{5,"V"},
	{4,"IV"},
	{1,"I"},
}

// Convert an integer to a ROman numeral with step tracking
func intToroman(num int) (string, string, error) {
	if num <= 0 || num >= 4000 {
		return "", "", fmt.
	}
}