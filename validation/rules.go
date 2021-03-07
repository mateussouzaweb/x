package validation

import (
	"net"
	"regexp"
	"strings"
	"time"
)

// String validates presence of string field
func (v *Validate) String(key string, value string, message string) {
	condition := len(strings.TrimSpace(value)) > 0
	v.Validate(key, condition, message, key, value)
}

// Int validates presence of int field
func (v *Validate) Int(key string, value int64, message string) {
	condition := value > 0
	v.Validate(key, condition, message, key, value)
}

// NegativeInt validates presence of int field
func (v *Validate) NegativeInt(key string, value int64, message string) {
	condition := value < 0
	v.Validate(key, condition, message, key, value)
}

// Float validates presence of float field
func (v *Validate) Float(key string, value float64, message string) {
	condition := value > 0
	v.Validate(key, condition, message, key, value)
}

// NegativeFloat validates presence of float field
func (v *Validate) NegativeFloat(key string, value float64, message string) {
	condition := value < 0
	v.Validate(key, condition, message, key, value)
}

// Time validates presence of time field
func (v *Validate) Time(key string, value time.Time, message string) {
	condition := value.IsZero() != true
	v.Validate(key, condition, message, key, value)
}

// Slice validates presence of slice field
func (v *Validate) Slice(key string, value []interface{}, message string) {
	condition := len(value) > 0
	v.Validate(key, condition, message, key, value)
}

// Map validates presence of map field
func (v *Validate) Map(key string, value map[interface{}]interface{}, message string) {
	condition := len(value) > 0
	v.Validate(key, condition, message, key, value)
}

// Length validates character length of string field
func (v *Validate) Length(key string, value string, exact int, message string) {
	condition := len(value) == exact
	v.Validate(key, condition, message, key, exact, value)
}

// MaxLength validates maximum character length of string field
func (v *Validate) MaxLength(key string, value string, max int, message string) {
	condition := len(value) <= max
	v.Validate(key, condition, message, key, max, value)
}

// MinLength validates minimum character length of string field
func (v *Validate) MinLength(key string, value string, min int, message string) {
	condition := len(value) >= min
	v.Validate(key, condition, message, key, min, value)
}

// Email validates if email is valid
func (v *Validate) Email(key string, value string, message string) {

	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	condition := regex.MatchString(value)

	v.Validate(key, condition, message, key, value)
}

// Domain validates if domain is valid
func (v *Validate) Domain(key string, value string, message string) {

	regex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-zA-Z]{1,63}| xn--[a-z0-9]{1,59})$`)
	condition := regex.MatchString(value)

	v.Validate(key, condition, message, key, value)
}

// IP validates if string in a valid IPV4 or IPV6
func (v *Validate) IP(key string, value string, message string) {

	ip := net.ParseIP(value)
	condition := ip != nil

	v.Validate(key, condition, message, key, value)
}

// IPV4 validates if string in a valid IPV4
func (v *Validate) IPV4(key string, value string, message string) {

	ip := net.ParseIP(value)
	condition := ip != nil && len(ip) == 4

	v.Validate(key, condition, message, key, value)
}

// IPV6 validates if string in a valid IPV6
func (v *Validate) IPV6(key string, value string, message string) {

	ip := net.ParseIP(value)
	condition := ip != nil && len(ip) == 16

	v.Validate(key, condition, message, key, value)
}

// Pattern validates if string follow the desired pattern
func (v *Validate) Pattern(key string, value string, pattern string, message string) {

	regex := regexp.MustCompile(pattern)
	condition := regex.MatchString(value)

	v.Validate(key, condition, message, key, value, pattern)
}

// InList validates if string is in list of valid values
func (v *Validate) InList(key string, value string, list []string, message string) {

	len := len(list)
	contains := false

	for index := 0; index < len; index++ {
		if list[index] == value {
			contains = true
			break
		}
	}

	v.Validate(key, contains, message, key, strings.Join(list, ", "), value)
}
