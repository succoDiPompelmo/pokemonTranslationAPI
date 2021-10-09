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

func emptyTranslationResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(429, `{"error":{"code": 429, "message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 58 seconds."}}`)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func legendaryPokemonResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(200, legendaryPokemonData)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func cavePokemonResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(200, cavePokemonData)
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func invalidPokemonResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(404, "Not Found")
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

func internalServerError(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(500, "")
	resp.Header.Add("Content-Type", "application/json")
	return resp, nil
}

// func f(req *http.Request) (*http.Response, error) {
// 	resp := httpmock.NewStringResponse(429, `{"error":{"code": 429, "message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 58 seconds."}}`)
// 	resp.Header.Add("Content-Type", "application/json")
// 	return resp, nil
// }