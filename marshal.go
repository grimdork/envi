package envi

import (
	"sort"

	"github.com/Urethramancer/signor/stringer"
)

// Marshal INI-structure into a string.
func (ini *INI) Marshal() string {
	proplist := []*Property{}
	for _, p := range ini.Properties {
		proplist = append(proplist, p)
	}
	s := stringer.New()
	if len(proplist) > 0 {
		sort.Sort(byPropName(proplist))
		for _, p := range proplist {
			s.WriteString(p.String())
		}
		s.WriteString("\n")
	}

	seclist := []*Section{}
	for _, sec := range ini.Sections {
		seclist = append(seclist, sec)
	}
	if len(seclist) > 0 {
		sort.Sort(bySecName(seclist))
		for _, sec := range seclist {
			s.WriteString(sec.String())
		}
	}

	return s.String()
}

//
// Section sorting.
//

type bySecName []*Section

func (b bySecName) Len() int {
	return len(b)
}

func (b bySecName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b bySecName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//
// Property sorting.
//

type byPropName []*Property

func (b byPropName) Len() int {
	return len(b)
}

func (b byPropName) Less(i, j int) bool {
	return b[i].Key < b[j].Key
}

func (b byPropName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
