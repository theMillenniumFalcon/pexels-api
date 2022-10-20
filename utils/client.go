package utils

import "net/http"

type Client struct {
	Token          string
	HttpClient     http.Client
	RemainingTimes int32
}
