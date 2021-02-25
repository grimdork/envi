package envi_test

import (
	"os"
	"testing"

	"github.com/Urethramancer/envi"
)

func TestMarshalling(t *testing.T) {
	ini := envi.New()
	ini.ForceUpper()
	ini.Set("main", "host", "localhost")
	ini.Set("main", "port", "2080")
	ini.Set("", "loose", "Not in a section.")

	out1 := ini.Marshal()
	t.Logf("\n%s", out1)
	ini = envi.Unmarshal(out1)
	ini.ForceUpper()
	out2 := ini.Marshal()
	if out1 != out2 {
		t.Error("Re-marshal mismatch!")
		t.FailNow()
	}

	t.Log("Marshalled and re-marshalled files are the same.")

	h := ini.GetEnvString("main", "host")
	p := ini.GetEnvString("main", "port")
	t.Logf("Host and port before setting envvars: %s:%s", h, p)
	os.Setenv("HOST", "0.0.0.0")
	os.Setenv("PORT", "8000")
	h2 := ini.GetEnvString("main", "host")
	p2 := ini.GetEnvString("main", "port")
	t.Logf("Host and port after setting envvars: %s:%s", h2, p2)
	if h == h2 || p == p2 {
		t.Error("Variables shouldn't be the same after reading from environment!")
		t.Logf("\nh = %s\nh2 = %s\np = %s\np2 = %s", h, h2, p, p2)
		t.FailNow()
	}

	l := ini.GetString("", "loose")
	if l == "" {
		t.Error("Loose var not found.")
		t.FailNow()
	}

	t.Logf("Loose variable = '%s'", l)
	t.Log("Basic sanity check passed.")
}

func TestSetGet(t *testing.T) {
	data := `host = 0.0.0.0
port = 8000
`
	ini := envi.Unmarshal(data)
	h := ini.GetString("", "host")
	if h != "0.0.0.0" {
		t.Logf("Host mismatch:\nExpected '0.0.0.0'\t\nGot '%s'\n", h)
		t.FailNow()
	}
	p := ini.GetString("", "port")
	if p != "8000" {
		t.Logf("Port mismatch:\nExpected '8000'\t\nGot '%s'\n", p)
		t.FailNow()
	}
}
