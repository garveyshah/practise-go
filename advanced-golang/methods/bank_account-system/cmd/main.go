package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("		=============== Godwin's simple Bank Management System =============== 	")
	fmt.Println("        	    =========== Where your finacial issues are sorted  ============		")
	fmt.Println()
	fmt.Println("				Welcome: 				 ")
	fmt.Println("					[1]. Login					")
	fmt.Println("					[2]. Register					")
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input format")
		return
	}

	switch option {
	case "1": 
	fmt.Println("Enter Name: ")
	case "2":
		fmt.Println("Enter Name: ")
	}
	reader
	fmt.Println(option)
}

func ()
