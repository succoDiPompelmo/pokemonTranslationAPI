package main

import (
	"net/http"
	"github.com/jarcoal/httpmock"
)

func yodaResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(200, `{"contents":{"translated": "You doing young man,  how are"}}`)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func shakespeareResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(200, `{"contents":{"translated": "How art thee doing young sir"}}`)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func emptyResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(429, `{"error":{"code": 429, "message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 58 seconds."}}`)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}