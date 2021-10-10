package main

import (
	"net/url"
	"log"
)

type TranslationResponse struct {
	Contents ContentsResource
}

type ContentsResource struct {
	Translated string
}

func getTranslatedDescription(appCtx AppCtx, description string, habitat string, isLegendary bool) (string, error) {

	translationUrl := getTranslationURL(appCtx, habitat, isLegendary)
	resp, err := appCtx.client.R().
		SetResult(&TranslationResponse{}).
		Get(translationUrl + url.PathEscape(description))

	if err != nil {
		log.Printf("GET TRANSLATED DESCRIPTION ERROR: %s for description %s and obtained response %s", err.Error(), description, string(resp.Body()))
		return description, err
	}
	if resp.StatusCode() > 399 {
		log.Printf("GET TRANSLATED DESCRIPTION ERRO: for description %s and returned status code %d with response %s", description, resp.StatusCode(), string(resp.Body()))
		return description, &RequestError{StatusCode: resp.StatusCode(), Err: err,}
	}

    return resp.Result().(*TranslationResponse).Contents.Translated, err
}

func getTranslationURL(appCtx AppCtx, habitat string, isLegendary bool) string {
	if habitat == "cave" || isLegendary {
		return appCtx.translationUrl + "yoda.json?text="
	}
	return appCtx.translationUrl + "shakespeare.json?text="
}