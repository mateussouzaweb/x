package env

import "os"

// Get returns the value of a environment variable or a default value is case of not defined value
func Get(variable string, defaultValue string) string {

	value, exists := os.LookupEnv(variable)

	if !exists {
		value = defaultValue
	}

	return value
}

// Set a environment variable
func Set(variable string, value string) error {
	return os.Setenv(variable, value)
}

// Unset a environment variable
func Unset(variable string) error {
	return os.Unsetenv(variable)
}

// Expand a string containing environment variables in the form of $var or ${var}
func Expand(template string) string {
	return os.ExpandEnv(template)
}
