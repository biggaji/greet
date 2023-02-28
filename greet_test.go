package greet

import (
	"regexp"
	"testing"
)

func TestSayHi(t *testing.T) {
	name := "Dolapo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := SayHi("Dolapo")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Dolapo") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestSayHiEmpty(t *testing.T) {
	msg, err := SayHi("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
