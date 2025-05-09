package env

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Load environment file and set runtime variables
func LoadFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, file.Close())
	}()

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
			err = os.Setenv(variable, value)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// Reads a variable entry from line and returns the values
func ReadLine(line string) (string, string, error) {

	// Ignore if line starts with comments
	if strings.HasPrefix(line, `#`) {
		return "", "", nil
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

	// Unquote if necessary
	if strings.ContainsAny(value, "\"'`") {
		var err error
		value, err = strconv.Unquote(value)
		if err != nil {
			return "", "", err
		}
	}

	return variable, value, nil
}
