package envi

import "os"

// Load an INI file.
func Load(fn string) (*INI, error) {
	data, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	ini := Unmarshal(string(data))
	return ini, nil
}
