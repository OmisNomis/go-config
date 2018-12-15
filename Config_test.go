package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	conf := NewConfig("config.json")

	s := conf.Get("simpleString")

	if s != "one" {
		t.Errorf("Expected one, got %s", s)
	}

	n := conf.Get("simpleNumber")

	if n != 1 {
		t.Errorf("Expected 1, got %d", n)
	}

	b := conf.Get("simpleBool")

	if b != true {
		t.Errorf("Expected true, got %t", b)
	}
}
