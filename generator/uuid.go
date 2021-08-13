package generator

import (
	"strings"

	"github.com/google/uuid"
)

// UUID method
func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
