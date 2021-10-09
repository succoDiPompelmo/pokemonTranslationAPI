package main

import (
	"net/url"
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
		return description, err
	}
	if resp.StatusCode() > 399 {
		return description, &RequestError{statusCode: resp.StatusCode(), err: err,}
	}

    return resp.Result().(*TranslationResponse).Contents.Translated, err
}

func getTranslationURL(appCtx AppCtx, habitat string, isLegendary bool) string {
	if habitat == "cave" || isLegendary {
		return "https://api.funtranslations.com/translate/yoda.json?text="
	}
	return "https://api.funtranslations.com/translate/shakespeare.json?text="
}