package crawler

import "testing"

var urlCSR = "https://cmlabs.co"
var urlSSR = "https://thoriqadillah.github.io/cat-n-code/"

func TestIsCSRWork(t *testing.T) {
	expected := true
	result := isCSR(urlCSR, 20)

	if expected != result {
		t.Fail()
	}
}

func TestIsSSR(t *testing.T) {
	expected := false
	result := isCSR(urlSSR, 20)

	if expected == result {
		t.Fail()
	}
}
