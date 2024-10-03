package main

// TODO:
//  1. put the input processing logic code in a separate function

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project01/100daycodechallange/day16/funcs"
)

func main() {
	fmt.Println("Hello welcome to Ouma's sorted array merger CLI>")
	fmt.Println("There's two ways to feed in data :\n")
	fmt.Println("	1. Data in a txt file")
	fmt.Println("	2. Type in input from the terminal\n")
	fmt.Println("Enter Option")

	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	data = strings.TrimSpace(data)
	num, err := strconv.Atoi(data)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	data2 := ""
	if num == 1 {
		fmt.Println("Enter filename : ")
		reader := bufio.NewReader(os.Stdin)
		data1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		data2 = data1
		fmt.Println(data2)
	} else if num == 2 {
		fmt.Println("Type in slice1 - data : ... \n")
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		data = strings.TrimSpace(data)
		slice := strings.Fields(data)
		var slice1 []int

		for _, digit := range slice {
			num, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println("Err: ", err)
				return
			}
			slice1 = append(slice1, num)
		}

		fmt.Println("Type in slice2 - data : ... \n")
		reader2 := bufio.NewReader(os.Stdin)
		data2, err := reader2.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		data2 = strings.TrimSpace(data2)
		slice0 := strings.Fields(data2)
		var slice2 []int

		for _, digit := range slice0 {
			num, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println("Err: ", err)
				return
			}
			slice2 = append(slice2, num)
		}

		fmt.Printf("Slice 1 :- %v,\n\nSlice 2 :- %v\n\n", slice1, slice2)

		merge := funcs.Merge(slice1, slice2)
		fmt.Printf("Merged Slice =- %v,\n\n", merge)

	} else {
		fmt.Println("Wrong turn")
	}
}
