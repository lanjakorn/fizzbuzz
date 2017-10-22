package main

import (
	"testing"
)

func TestFizzbuzz1(t *testing.T) {
	original := 1
	expected := "1"

	s := fizzbuzz(original)

	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}

func TestFizzbuzz3(t *testing.T) {
	original := 3
	expected := "fizz"

	s := fizzbuzz(original)

	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}

func TestFizzbuzz5(t *testing.T) {
	original := 5
	expected := "buzz"

	s := fizzbuzz(original)

	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}

func TestFizzbuzz15(t *testing.T) {
	original := 15
	expected := "fizzbuzz"

	s := fizzbuzz(original)

	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}
