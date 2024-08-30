package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(PrintName())
}

func PrintName() string {
	fmt.Println("What is your Name?")

	Reader := bufio.NewReader(os.Stdin)
	name, _ := Reader.ReadString('\n')

	if name == "\n" {
		fmt.Println("Please input your name...")

		Reader = bufio.NewReader(os.Stdin)
		name, _ = Reader.ReadString('\n')
	}

	fmt.Println("How old are you?")
	Reader1 := bufio.NewReader(os.Stdin)
	age, _ := Reader1.ReadString('\n')

	if age == "\n" {
		fmt.Println("Please type in your age ...")
		Reader1 = bufio.NewReader(os.Stdin)
		age, _ = Reader1.ReadString('\n')

	}
	fmt.Println()
	return name + age
}