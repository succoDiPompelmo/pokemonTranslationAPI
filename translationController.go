package main

import (
	"net/url"
	"log"
	"encoding/json"
)

type TranslationResponse struct {
	Contents ContentsResource
}

type ContentsResource struct {
	Translated string
}

func getTranslatedDescription(appCtx AppCtx, description string, habitat string, isLegendary bool) (string, error) {

	translationUrl := getTranslationURL(appCtx, habitat, isLegendary)
	resp, err := doGet(appCtx, translationUrl + url.PathEscape(description))

	if err != nil {
		log.Printf("GET TRANSLATED DESCRIPTION ERROR: %s for description %s and obtained response %s", err.Error(), description, string(resp.responseBody))
		return description, err
	}
	if resp.statusCode > 399 {
		log.Printf("GET TRANSLATED DESCRIPTION ERRO: for description %s and returned status code %d with response %s", description, resp.statusCode, string(resp.responseBody))
		return description, &RequestError{StatusCode: resp.statusCode, Err: err,}
	}

	var result TranslationResponse
	json.Unmarshal(resp.responseBody, &result)
    return result.Contents.Translated, err
}

func getTranslationURL(appCtx AppCtx, habitat string, isLegendary bool) string {
	if habitat == "cave" || isLegendary {
		return appCtx.translationUrl + "yoda.json?text="
	}
	return appCtx.translationUrl + "shakespeare.json?text="
}