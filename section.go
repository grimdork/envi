package envi

import (
	"sort"
	"strings"

	"github.com/Urethramancer/signor/stringer"
)

// Section with properties.
type Section struct {
	// Name (title) of the section.
	Name string
	// Properties in this section.
	Properties map[string]*Property

	dirty bool
	s     stringer.Stringer
}

// GetString from property in this section, or empty if it doesn't exist.
func (sec *Section) GetString(k string) string {
	p, ok := sec.Properties[k]
	if !ok {
		return ""
	}

	return p.Value
}

// GetBool from property in this section, or false if it doesn't exist.
func (sec *Section) GetBool(k string) bool {
	p, ok := sec.Properties[k]
	if !ok {
		return false
	}

	return p.GetBool()
}

// GetInt from property in this section, or 0 if it doesn't exist.
func (sec *Section) GetInt(k string) int64 {
	p, ok := sec.Properties[k]
	if !ok {
		return 0
	}

	return p.GetInt()
}

// GetFloat from property in this section, or 0.0 if it doesn't exist.
func (sec *Section) GetFloat(k string) float64 {
	p, ok := sec.Properties[k]
	if !ok {
		return 0.0
	}

	return p.GetFloat()
}

// Set or replace a property in the section.
func (sec *Section) Set(k, v string) {
	k = strings.ToLower(k)
	p, ok := sec.Properties[k]
	if ok {
		p.Value = v
	} else {
		sec.Properties[k] = NewProp(sec, k, v)
	}
}

// SetDirty marks string output for recalculation.
func (sec *Section) SetDirty() {
	sec.dirty = true
}

// String output of the section, with a trailing newline.
func (sec *Section) String() string {
	if sec.dirty {
		sec.s.Reset()
		_, _ = sec.s.WriteStrings("[", sec.Name, "]\n")
		proplist := []*Property{}
		for _, p := range sec.Properties {
			proplist = append(proplist, p)
		}
		if len(proplist) > 0 {
			sort.Sort(byPropName(proplist))
			for _, p := range proplist {
				sec.s.WriteString(p.String())
			}
		}
		sec.s.WriteString("\n")
	}
	return sec.s.String()
}
