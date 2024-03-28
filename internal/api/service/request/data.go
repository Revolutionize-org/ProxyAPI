package request

import "net/http"

type header map[string]string

type requestInformation struct {
	URL    string `json:"Url" validate:"required"`
	Method string `json:"Method" validate:"required"`
	Header header `json:"Header"`
	Body   string `json:"Body"`
}

type requestResponse struct {
	Status        string      `json:"Status"`
	StatusCode    int         `json:"StatusCode"`
	Proto         string      `json:"Proto"`
	ProtoMajor    int         `json:"ProtoMajor"`
	ProtoMinor    int         `json:"ProtoMinor"`
	Header        http.Header `json:"Header"`
	Body          interface{} `json:"Body"`
	ContentLength int64       `json:"ContentLength"`
}
