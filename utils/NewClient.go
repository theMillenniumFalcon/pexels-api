package utils

import (
	"net/http"
)

func NewClient(token string) Client {
	c := http.Client{}
	return Client{Token: token, HttpClient: c}
}
