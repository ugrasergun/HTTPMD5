package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func getURIMD5(uri string) (fullURI string, MD5 string) {

	fullURI = getFullURI(uri)

	body, err := getMessageBody(fullURI)

	var MD5bytes [16]byte

	if err != nil {
		MD5bytes = (md5.Sum([]byte(err.Error())))
	} else {
		MD5bytes = md5.Sum(body)
	}

	MD5 = hex.EncodeToString(MD5bytes[:])

	return
}

func getFullURI(uri string) string {

	parsedURI, err := url.Parse(uri)

	if err != nil {
		log.Fatal(err)
	}

	if !parsedURI.IsAbs() {
		parsedURI.Scheme = "http"
	}

	fullURI, err := url.ParseRequestURI(parsedURI.String())

	if err != nil {
		log.Fatal(err)
	}

	return fullURI.String()
}

func getMessageBody(uri string) (body []byte, err error) {

	resp, err := http.Get(uri)

	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	return
}
