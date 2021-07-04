package crypt

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	const (
		key     = ThingBin
		encoded = "\x2c\xc3\x70\x31\x5e\xda\x12\x3c"
		decoded = "ROLF\x01\x00\x00\x00"
	)

	buf := []byte(encoded)
	err := Decode(buf, key)
	if err != nil {
		t.Fatal(err)
	} else if string(buf) != decoded {
		t.Fatalf("unexpected data: %q", buf)
	}

	err = Encode(buf, key)
	if err != nil {
		t.Fatal(err)
	} else if string(buf) != encoded {
		t.Fatalf("unexpected data: %q", buf)
	}
}

func TestReader(t *testing.T) {
	const (
		key     = ThingBin
		encoded = "\x2c\xc3\x70\x31\x5e\xda\x12\x3c"
		decoded = "ROLF\x01\x00\x00\x00"
	)

	r, err := NewReader(strings.NewReader(encoded), key)
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	} else if string(out) != decoded {
		t.Fatalf("unexpected data: %q", out)
	}
}
