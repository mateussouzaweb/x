package validation

import (
	"net"
	"regexp"
	"strings"
	"time"
)

// PresenceOf validates presence of string field
func (v *Validate) PresenceOf(name string, value string) {
	cond := len(strings.TrimSpace(value)) > 0
	v.Validate(cond, "%s cannot be blank", name)
}

// PresenceOfInt validates presence of int field
func (v *Validate) PresenceOfInt(name string, value int64) {
	cond := value > 0
	v.Validate(cond, "%s cannot be blank or zero", name)
}

// PresenceOfNegativeInt validates presence of int field
func (v *Validate) PresenceOfNegativeInt(name string, value int64) {
	cond := value < 0
	v.Validate(cond, "%s cannot be blank or more than zero", name)
}

// PresenceOfFloat validates presence of float field
func (v *Validate) PresenceOfFloat(name string, value float64) {
	cond := value > 0
	v.Validate(cond, "%s cannot be blank or zero", name)
}

// PresenceOfNegativeFloat validates presence of float field
func (v *Validate) PresenceOfNegativeFloat(name string, value float64) {
	cond := value < 0
	v.Validate(cond, "%s cannot be blank or more than zero", name)
}

// PresenceOfTime validates presence of time field
func (v *Validate) PresenceOfTime(name string, value time.Time) {
	cond := value.IsZero() != true
	v.Validate(cond, "%s cannot be blank or have a empty datetime", name)
}

// PresenceOfSlice validates presence of slice field
func (v *Validate) PresenceOfSlice(name string, value []interface{}) {
	cond := len(value) > 0
	v.Validate(cond, "%s cannot be empty", name)
}

// MaxLength validates maximum character length of string field
func (v *Validate) MaxLength(name string, value string, max int) {
	cond := len(value) < max
	v.Validate(cond, "%s cannot be greater than %d characters", name, max)
}

// MinLength validates minimum character length of string field
func (v *Validate) MinLength(name string, value string, min int) {
	cond := len(value) > min
	v.Validate(cond, "%s must be at least %d characters", name, min)
}

// Email validates if email is valid
func (v *Validate) Email(name string, value string) {

	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	match := regex.MatchString(value)

	v.Validate(match, "%s is not in valid email address", name)
}

// Domain validates if domain is valid
func (v *Validate) Domain(name string, value string) {

	regex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-zA-Z]{1,63}| xn--[a-z0-9]{1,59})$`)
	match := regex.MatchString(value)

	v.Validate(match, "%s is not in valid domain", name)
}

// IP validates if string in a valid IPV4 or IPV6
func (v *Validate) IP(name string, value string) {

	ip := net.ParseIP(value)
	cond := ip != nil

	v.Validate(cond, "%s must be a valid IPV4 or IPV6 address", name)
}

// IPV4 validates if string in a valid IPV4
func (v *Validate) IPV4(name string, value string) {

	ip := net.ParseIP(value)
	cond := ip != nil && len(ip) == 4

	v.Validate(cond, "%s must be a valid IPV4 address", name)
}

// IPV6 validates if string in a valid IPV6
func (v *Validate) IPV6(name string, value string) {

	ip := net.ParseIP(value)
	cond := ip != nil && len(ip) == 16

	v.Validate(cond, "%s must be a valid IPV6 address", name)
}

// InList validates if string is in list of valid values
func (v *Validate) InList(name string, value string, list []string) {

	len := len(list)
	contains := false

	for index := 0; index < len; index++ {
		if list[index] == value {
			contains = true
			break
		}
	}

	v.Validate(contains, "%s must be one of the options: %s", name, strings.Join(list, ", "))
}
