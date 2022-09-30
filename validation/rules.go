package validation

import (
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// String validates presence of string field
func (v *Validate) String(key string, value string, message error) {
	condition := len(strings.TrimSpace(value)) > 0
	v.Validate(key, condition, message)
}

// Int validates presence of int field
func (v *Validate) Int(key string, value int64, message error) {
	condition := value > 0
	v.Validate(key, condition, message)
}

// NegativeInt validates presence of int field
func (v *Validate) NegativeInt(key string, value int64, message error) {
	condition := value < 0
	v.Validate(key, condition, message)
}

// Float validates presence of float field
func (v *Validate) Float(key string, value float64, message error) {
	condition := value > 0
	v.Validate(key, condition, message)
}

// NegativeFloat validates presence of float field
func (v *Validate) NegativeFloat(key string, value float64, message error) {
	condition := value < 0
	v.Validate(key, condition, message)
}

// Time validates presence of time field
func (v *Validate) Time(key string, value time.Time, message error) {
	condition := !value.IsZero()
	v.Validate(key, condition, message)
}

// Slice validates presence of slice field
func (v *Validate) Slice(key string, value []interface{}, message error) {
	condition := len(value) > 0
	v.Validate(key, condition, message)
}

// Map validates presence of map field
func (v *Validate) Map(key string, value map[interface{}]interface{}, message error) {
	condition := len(value) > 0
	v.Validate(key, condition, message)
}

// Length validates character length of string field
func (v *Validate) Length(key string, value string, exact int, message error) {
	condition := len(value) == exact
	v.Validate(key, condition, message)
}

// MaxLength validates maximum character length of string field
func (v *Validate) MaxLength(key string, value string, max int, message error) {
	condition := len(value) <= max
	v.Validate(key, condition, message)
}

// MinLength validates minimum character length of string field
func (v *Validate) MinLength(key string, value string, min int, message error) {
	condition := len(value) >= min
	v.Validate(key, condition, message)
}

// Email validates if email is valid
func (v *Validate) Email(key string, value string, message error) {

	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	condition := regex.MatchString(value)

	v.Validate(key, condition, message)
}

// Domain validates if domain is valid
func (v *Validate) Domain(key string, value string, message error) {

	regex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-zA-Z]{1,63}| xn--[a-z0-9]{1,59})$`)
	condition := regex.MatchString(value)

	v.Validate(key, condition, message)
}

// URL validates if URL is valid
func (v *Validate) URL(key string, value string, message error) {

	parse, err := url.Parse(value)
	condition := err == nil && parse.Scheme != "" && parse.Host != ""

	v.Validate(key, condition, message)
}

// IP validates if string in a valid IPV4 or IPV6
func (v *Validate) IP(key string, value string, message error) {

	ip := net.ParseIP(value)
	condition := ip != nil

	v.Validate(key, condition, message)
}

// IPV4 validates if string in a valid IPV4
func (v *Validate) IPV4(key string, value string, message error) {

	ip := net.ParseIP(value)
	condition := ip != nil && len(ip) == 4

	v.Validate(key, condition, message)
}

// IPV6 validates if string in a valid IPV6
func (v *Validate) IPV6(key string, value string, message error) {

	ip := net.ParseIP(value)
	condition := ip != nil && len(ip) == 16

	v.Validate(key, condition, message)
}

// Pattern validates if string follow the desired pattern
func (v *Validate) Pattern(key string, value string, pattern string, message error) {

	regex := regexp.MustCompile(pattern)
	condition := regex.MatchString(value)

	v.Validate(key, condition, message)
}

// InList validates if string is in list of valid values
func (v *Validate) InList(key string, value string, list []string, message error) {

	len := len(list)
	contains := false

	for index := 0; index < len; index++ {
		if list[index] == value {
			contains = true
			break
		}
	}

	v.Validate(key, contains, message)
}
