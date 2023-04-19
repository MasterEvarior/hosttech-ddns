package main

import (
	"io"
	"net/http"
)

const apiUrl = "https://api.ipify.org"

func GetPublicIp() (string, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
