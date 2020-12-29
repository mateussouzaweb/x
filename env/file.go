package env

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Load method
func Load(filename string) error {

	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	for _, line := range lines {

		variable, value, err := ReadLine(line)

		if err != nil {
			return err
		}

		if variable != "" {
			err = Set(variable, value)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

// ReadLine reads a variable entry line and returns the values
func ReadLine(line string) (string, string, error) {

	// Remove comments
	if strings.Contains(line, "#") {
		split := strings.SplitN(line, "#", 2)
		line = split[0]
	}

	line = strings.TrimSpace(line)

	// Ignore if line is empty
	if len(line) == 0 {
		return "", "", nil
	}

	// Split by =
	pair := strings.SplitN(line, "=", 2)
	variable := pair[0]
	value := pair[1]

	value, err := strconv.Unquote(value)

	if err != nil {
		return "", "", err
	}

	return variable, value, nil
}
