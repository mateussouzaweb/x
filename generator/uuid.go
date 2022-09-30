package generator

import (
	"strings"

	"github.com/google/uuid"
)

// UUID generates a new random UUID without dashes
func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
