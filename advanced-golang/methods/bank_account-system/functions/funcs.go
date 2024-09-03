package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name      string
	Password  string
	PhoneNo   string
	AccountNo string

	
}

var Users = make(map[string]string)

// Methods

// getData to get user input data and populate the strucs.
func getData(option string) (interface{}, error) {
	reader := bufio.NewReader(os.Stdin)

	switch strings.ToLower(option) {
	case "login":
		fmt.Println("Enter Name: ")
		return getLoginData(reader)
	case "register":
		fmt.Println("Enter Name: ")
		return getRegisterData(reader)
	default: 
	return nil, fmt.Errorf("unkown field")
	}
}

func (u User)getLoginData(reader *bufio.Reader) (interface{}, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return Users,nil
	}

	if input == "" {
		return Users, fmt.Errorf("Please enter your name")
	}
	

}