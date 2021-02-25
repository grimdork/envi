package envi

import (
	"os"
	"strconv"
	"strings"
)

// INI main structure.
type INI struct {
	// Sections hold properties.
	Sections   map[string]*Section
	Properties map[string]*Property
	upper      bool
}

// New INI convenience function.
func New() *INI {
	return &INI{
		Sections:   make(map[string]*Section),
		Properties: make(map[string]*Property),
	}
}

// ForceUpper checks the environment for upper-case versions of the supplied INI variables.
func (ini *INI) ForceUpper() {
	ini.upper = true
}

// GetString returns a variable from specified section.
func (ini *INI) GetString(s, k string) string {
	l := strings.ToLower(k)
	if s == "" {
		p, ok := ini.Properties[l]
		if !ok {
			return ""
		}

		return p.Value
	}

	sec, ok := ini.Sections[s]
	if !ok {
		return ""
	}

	return sec.GetString(l)
}

// GetEnvString returns a variable from the specified section, or overridden by an environment variable.
func (ini *INI) GetEnvString(s, k string) string {
	if ini.upper {
		k = strings.ToUpper(k)
	}
	v := os.Getenv(k)
	if v != "" {
		return v
	}

	return ini.GetString(s, strings.ToLower(k))
}

// GetBool returns the boolean value for a variable in the INI file.
// Non-existense results in false as the return value.
func (ini *INI) GetBool(s, k string) bool {
	if s == "" {
		p, ok := ini.Properties[k]
		if !ok {
			return false
		}

		return p.GetBool()
	}

	sec, ok := ini.Sections[s]
	if !ok {
		return false
	}

	return sec.GetBool(k)
}

// GetEnvBool returns GetBool(), overriding with an environment variable if found.
func (ini *INI) GetEnvBool(s, k string) bool {
	if ini.upper {
		k = strings.ToUpper(k)
	}

	v := os.Getenv(k)
	if v != "" {
		return BoolValue(v)
	}

	return ini.GetBool(s, strings.ToLower(k))
}

// GetInt returns the integer value for a variable in the INI file.
// Non-existense results in 0 as the return value.
func (ini *INI) GetInt(s, k string) int64 {
	if s == "" {
		p, ok := ini.Properties[k]
		if !ok {
			return 0
		}

		return p.GetInt()
	}

	sec, ok := ini.Sections[s]
	if !ok {
		return 0
	}

	return sec.GetInt(k)
}

// GetEnvInt returns GetInt(), overriding with an environment variable if found.
func (ini *INI) GetEnvInt(s, k string) int64 {
	if ini.upper {
		k = strings.ToUpper(k)
	}

	v := os.Getenv(k)
	if v != "" {
		x, _ := strconv.ParseInt(v, 10, 64)
		return x
	}

	return ini.GetInt(s, strings.ToLower(k))
}

// GetFloat returns the float value for a variable in the INI file.
// Non-existense results in 0 as the return value.
func (ini *INI) GetFloat(s, k string) float64 {
	if s == "" {
		p, ok := ini.Properties[k]
		if !ok {
			return 0.0
		}

		return p.GetFloat()
	}

	sec, ok := ini.Sections[s]
	if !ok {
		return 0.0
	}

	return sec.GetFloat(k)
}

// GetEnvFloat returns GetFloat(), overriding with an environment variable if found.
func (ini *INI) GetEnvFloat(s, k string) float64 {
	if ini.upper {
		k = strings.ToUpper(k)
	}

	v := os.Getenv(k)
	if v != "" {
		x, _ := strconv.ParseFloat(v, 64)
		return x
	}

	return ini.GetFloat(s, strings.ToLower(k))
}

// Set is a convenience method to drill down into the correct section to set a property.
// The section will be created if missing.
func (ini *INI) Set(s, k, v string) {
	if s == "" {
		ini.Properties[k] = NewProp(nil, k, v)
		return
	}

	sec, ok := ini.Sections[s]
	if !ok {
		sec = ini.MakeSection(s)
	}
	sec.Properties[k] = NewProp(nil, k, v)
}

// MakeSection will create a new section, or return an existing one with the same name.
func (ini *INI) MakeSection(name string) *Section {
	sec := ini.Sections[name]
	if sec != nil {
		return sec
	}

	sec = &Section{
		Name:       name,
		Properties: make(map[string]*Property),
		dirty:      true,
	}

	ini.Sections[name] = sec
	return sec
}

// BoolValue of supplied string.
func BoolValue(s string) bool {
	switch s {
	case "true", "on", "enabled", "yes":
		return true
	}
	return false
}
