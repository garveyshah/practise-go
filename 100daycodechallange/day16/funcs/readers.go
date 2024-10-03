package funcs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ToDo : A reader for reading file content and returning it as a float64
type SliceReader interface {
	Reader() ([]int, error)
}

type InputSlice []int

func (in InputSlice) Reader() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	data = strings.TrimSpace(data)
	slice := strings.Fields(data)
	var slice1 []int

	for _, digit := range slice {
		num, err := strconv.Atoi(digit)
		if err != nil {
			return nil, err
		}
		slice1 = append(slice1, num)
	}
	return slice1, nil
}

type OptionReader interface {
	CReader() (int, float64)
}