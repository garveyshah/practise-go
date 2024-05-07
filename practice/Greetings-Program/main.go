package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Prompt the user to input their name
	fmt.Println("Hello, please enter your name:")

	// Read the input
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	name = strings.TrimSuffix(name, "\n")
	fmt.Printf("Nice meeting you %s. How are you doing?\n", name)

	// Read the user's response
	reader2 := bufio.NewReader(os.Stdin)
	mood, err := reader2.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	mood = strings.TrimSuffix(mood, "\n")

	var response string

	// Analyze the user's response to determine the appropriate reaction
	if strings.Contains(strings.ToLower(mood), "happy") || strings.Contains(strings.ToLower(mood), "good") || strings.Contains(strings.ToLower(mood), "fine") {
		response = "That's great to hear! What do you want to work on today?"
	} else if strings.Contains(strings.ToLower(mood), "not happy") || strings.Contains(strings.ToLower(mood), "not okay") || strings.Contains(strings.ToLower(mood), "not good") || strings.Contains(strings.ToLower(mood), "sad") || strings.Contains(strings.ToLower(mood), "bad") {
		response = "I'm sorry to hear that. Hope things get better soon.\nWhat would you like to work on today?"
	} else if strings.Contains(strings.ToLower(mood), "okay") {
		response = "Hope things get better soon.\nWhat would you like to work on?"
	} else {
		response = "I see. What would you like to work on today?"
	}
	fmt.Println(response)

	// Read the user's task
	reader3 := bufio.NewReader(os.Stdin)
	task, err := reader3.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input: ", err)
	}

	task = strings.TrimSuffix(task, "\n")
	fmt.Printf("Great!, %s, is an interesting subject. Happy Coding ahead friend!\n", task)
}
