package envi

import (
	"strconv"
	"strings"

	"github.com/Urethramancer/signor/stringer"
)

// Property holds a key string and value of any type.
type Property struct {
	// section for calculating envvar name.
	section *Section
	// Key is the name of the property/variable.
	Key string
	// Value is a string. Use Get*() to convert.
	Value string
	// Env variable name.
	Env string

	s     stringer.Stringer
	dirty bool
}

// NewProp creates a new property.
// Keys are set to lowercase.
func NewProp(sec *Section, k, v string) *Property {
	p := &Property{
		section: sec,
		Key:     strings.ToLower(k),
		Env:     k,
		Value:   v,
		dirty:   true,
	}
	return p
}

// Set a string.
func (p *Property) Set(v string) {
	p.Value = v
	p.dirty = true
	p.section.SetDirty()
}

// String output of the property, with newline.
func (p *Property) String() string {
	if p.dirty {
		p.s.Reset()
		_, _ = p.s.WriteStrings(p.Key, " = ", p.Value, "\n")
	}
	return p.s.String()
}

// GetBool counts anything but "true", "on" or "enabled" as false.
func (p *Property) GetBool() bool {
	return BoolValue(p.Value)
}

// GetInt returns an int64 conversion of the value.
func (p *Property) GetInt() int64 {
	x, _ := strconv.ParseInt(p.Value, 10, 64)
	return x
}

// GetFloat returns a float64 conversion of the value.
func (p *Property) GetFloat() float64 {
	x, _ := strconv.ParseFloat(p.Value, 64)
	return x
}
