package main

import "testing"

func testFullURIs(t *testing.T) {
	t.Run("google.com", testFullURI("google.com", "http://google.com"))
	t.Run("www.amazon.com", testFullURI("www.amazon.com", "http://www.amazon.com"))
	t.Run("ftp://ftp.google.com", testFullURI("ftp://ftp.google.com", "ftp://ftp.google.com"))
}

func testFullURI(uri string, expected string) func(*testing.T) {

	return func(t *testing.T) {
		actual := getFullURI(uri)

		if actual != expected {
			t.Errorf("Error: excpected uri: %s actual uri: %s", expected, actual)
		}
	}
}
