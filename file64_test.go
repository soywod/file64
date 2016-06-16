package file64

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestEncode(t *testing.T) {
	var tests = []struct {
		path string
		code string
	}{
		{"tests/file.txt", "Q09OVEVOVA=="},
		{"tests/noexists.txt", ""},
	}

	for _, test := range tests {
		code, _ := Encode(test.path)

		if test.code != code {
			t.Errorf("Bad encoding for %q, expected %q, got %q", test.path, test.code, code)
		}
	}
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		code    string
		dest    string
		content string
	}{
		{"Q09OVEVOVA==", "tests/decode/decode.txt", "CONTENT"},
	}

	for _, test := range tests {
		if err := Decode(test.code, test.dest); err != nil {
			t.Errorf(err.Error())
		}

		buff, err := ioutil.ReadFile(test.dest)
		if err != nil {
			t.Errorf(err.Error())
		}

		content := string(buff)

		if test.content != content {
			t.Errorf("Bad decoding for %q, expected %q, got %q", test.dest, test.content, content)
		}

		os.Remove(test.dest)
	}
}
