package config

import (
	"testing"
)

func Test_New(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatal(err)
	}
}
