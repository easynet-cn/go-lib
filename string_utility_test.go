package golib

import "testing"

func TestLowerCamelCase(t *testing.T) {
	str := "LowerCamelCase"

	if s := LowerCamelCase(str); s != "lowerCamelCase" {
		t.Error(s)
	}
}

func TestUpperCamelCase(t *testing.T) {
	str := "upperCamelCase"

	if s := UpperCamelCase(str); s != "UpperCamelCase" {
		t.Error(s)
	}
}

func TestSnakeCase(t *testing.T) {
	str := "LowerCamelCase"

	if s := SnakeCase(str); s != "lower_camel_case" {
		t.Error(s)
	}
}

func TestSpinalCase(t *testing.T) {
	str := "LowerCamelCase"

	if s := SpinalCase(str); s != "lower-camel-case" {
		t.Error(s)
	}
}
