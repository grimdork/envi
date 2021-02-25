package envi

import (
	"strings"
)

// Unmarshal a string with the contents of an INI file.
func Unmarshal(s string) *INI {
	if s == "" {
		return nil
	}

	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return nil
	}

	ini := New()
	for {
		if len(lines) == 0 {
			break
		}

		lines[0] = strings.TrimSpace(lines[0])
		if lines[0] == "" {
			lines = lines[1:]
			continue
		}

		if lines[0][0] == ';' || lines[0][0] == '#' {
			lines = lines[1:]
			continue
		}

		if lines[0][0] == '[' {
			lines = ini.parseSection(lines)
		} else {
			a := splitProp(lines[0])
			ini.Set("", a[0], a[1])
			lines = lines[1:]
		}
	}

	return ini
}

func (ini *INI) parseSection(lines []string) []string {
	lines[0] = lines[0][1 : len(lines[0])-1]
	sec := ini.MakeSection(lines[0])
	lines = lines[1:]
	for {
		if len(lines) == 0 {
			break
		}

		if lines[0] == "" {
			lines = lines[1:]
			continue
		}

		if lines[0][0] == ';' || lines[0][0] == '#' {
			lines = lines[1:]
			continue
		}

		if lines[0][0] == '[' {
			break
		}

		a := splitProp(lines[0])
		sec.Set(a[0], a[1])
		lines = lines[1:]
	}
	return lines
}

func splitProp(s string) []string {
	a := strings.SplitN(s, "=", 2)
	if len(a) == 1 {
		return []string{a[0], ""}
	}
	a[0] = strings.TrimSpace(a[0])
	a[1] = strings.TrimSpace(a[1])
	return a
}
