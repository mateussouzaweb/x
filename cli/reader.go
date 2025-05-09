package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Print question and read the reply
func Read(question string, defaultValue string) string {

	value := ""

	if defaultValue != "" {
		fmt.Printf("%s [%s] ", question, defaultValue)
	} else {
		fmt.Printf("%s ", question)
	}

	reader := bufio.NewReader(os.Stdin)
	value, _ = reader.ReadString('\n')
	value = strings.Replace(value, "\n", "", -1)

	if value == "" {
		value = defaultValue
	}

	return strings.Trim(value, " ")
}

// Print question and keep reading until result is yes (true) or no (false)
func YesOrNo(question string, defaultValue bool) bool {
	if defaultValue {
		question = question + " (Y/n)"
	} else {
		question = question + " (y/N)"
	}

	for {
		value := Read(question, "")
		value = strings.ToUpper(value)

		if value == "" && defaultValue {
			return true
		} else if value == "" && !defaultValue {
			return false
		} else if value == "Y" || value == "YES" {
			return true
		} else if value == "N" || value == "NO" {
			return false
		}
	}
}
