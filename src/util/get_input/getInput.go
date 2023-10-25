package util

import (
	"bufio"
	"os"
	"strings"
)

func GetUserInput() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
