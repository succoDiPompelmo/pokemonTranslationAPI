package main

import (
	"net/http"
	"github.com/jarcoal/httpmock"
	"time"
)

func initPokemonResponder(appCtx AppCtx) {

	httpmock.RegisterResponder("GET", appCtx.pokemonApiURL + "oddish", nonCaveNonLegendaryPokemonResponder)
	httpmock.RegisterResponder("GET", appCtx.pokemonApiURL + "zxcvb", invalidPokemonResponder)
	httpmock.RegisterResponder("GET", appCtx.pokemonApiURL + "internalServerError", internalServerError)
	httpmock.RegisterResponder("GET", appCtx.pokemonApiURL + "mewtwo", legendaryPokemonResponder)
	httpmock.RegisterResponder("GET", appCtx.pokemonApiURL + "diglett", cavePokemonResponder)

}

func initTranslationResponder(appCtx AppCtx) {

	testCaseText := "How%20are%20you%20doing%20young%20man"
	httpmock.RegisterResponder("GET", appCtx.translationUrl + "yoda.json?text=" + testCaseText, yodaResponder)
	httpmock.RegisterResponder("GET", appCtx.translationUrl + "shakespeare.json?text=" + testCaseText, shakespeareResponder)

}

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

func nonCaveNonLegendaryPokemonResponder(req *http.Request) (*http.Response, error) {
	resp := httpmock.NewStringResponse(200, nonLegendaryNonCavePokemonData)
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

func timeoutResponder(req *http.Request) (*http.Response, error) {
	time.Sleep(10 * time.Second)
	return nil, nil
}

// func f(req *http.Request) (*http.Response, error) {
// 	resp := httpmock.NewStringResponse(429, `{"error":{"code": 429, "message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 58 seconds."}}`)
// 	resp.Header.Add("Content-Type", "application/json")
// 	return resp, nil
// }