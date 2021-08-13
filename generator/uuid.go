package generator

import (
	"strings"

	"github.com/google/uuid"
)

// UUID generates a new ramdom UUID without dashes
func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
