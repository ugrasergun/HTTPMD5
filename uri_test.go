package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

const responseBody = "Response body"
const responseError = "NotFound"

func TestFullURIs(t *testing.T) {
	t.Run("google.com", testFullURI("google.com", "http://google.com"))
	t.Run("www.amazon.com", testFullURI("www.amazon.com", "http://www.amazon.com"))
	t.Run("ftp://ftp.google.com", testFullURI("ftp://ftp.google.com", "ftp://ftp.google.com"))
}

func testFullURI(uri string, expected string) func(*testing.T) {

	return func(t *testing.T) {
		actual := getFullURI(uri)

		if actual != expected {
			t.Errorf("URI Mismatch: excpected uri: %s actual uri: %s", expected, actual)
		}
	}
}

func TestGetMessageBodySuccess(t *testing.T) {
	testableHTTPGet = mockHTTPGet

	resp, err := getMessageBody("http://www.google.com")

	if err != nil {
		t.Errorf("No error expected. Actual Error: %s", err)
	}

	response := string(resp)

	if response != responseBody {
		t.Errorf("Error expected response: %s, actual response: %s", responseBody, response)
	}

	testableHTTPGet = http.Get

}

func TestGetMessageBodyFail(t *testing.T) {
	testableHTTPGet = mockHTTPGet

	_, err := getMessageBody("http://www.yahoo.com")

	if err == nil {
		t.Errorf("Error expected but not found!")
	}

	if err.Error() != responseError {
		t.Errorf("Error Mismatch! Expected Error:%s Actual Error:%s", responseError, err.Error())
	}

	testableHTTPGet = http.Get
}

func TestGetURIMD5(t *testing.T) {
	testableHTTPGet = mockHTTPGet

	uri, md5 := getURIMD5("www.google.com")

	if uri != "http://www.google.com" {
		t.Errorf("%s is not full URI", uri)
	}

	if md5 != "38eee230c1573327fd210b5c3dda7e19" {
		t.Errorf("MD5 mismatch! expected: 38eee230c1573327fd210b5c3dda7e19 actual:%s", md5)
	}

	testableHTTPGet = http.Get
}

func mockHTTPGet(uri string) (resp *http.Response, err error) {

	res := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(responseBody)),
	}

	if uri == "http://www.google.com" {
		return &res, nil
	}

	er := errors.New(responseError)

	return nil, er

}
