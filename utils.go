package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	catURL = "http://thecatapi.com/api/images/get?format=src&type=jpg"
)

func downloadFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return []byte{}, errors.New("Image url respose code != 200")
	}

	return ioutil.ReadAll(response.Body)
}

func downloadCat() ([]byte, error) {
	return downloadFromUrl(catURL)
}
