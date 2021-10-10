package main

import (
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/go-resty/resty/v2"
)

type TranslationInput struct {
	habitat string
	isLegendary bool
	description string
}

func TestGetTranslationURL(t *testing.T) {
	
	appCtx := AppCtx{
		translationUrl: FUN_TRANSLATION_API_URL,
	}
	
	tests := map[string]struct {
        input TranslationInput
        want string
    }{
        "Get URL for legendary pokemon": {input: TranslationInput{
			habitat: "rare", 
			isLegendary: true,}, want: "https://api.funtranslations.com/translate/yoda.json?text=",
		},
		"Get URL for cave pokemon": {input: TranslationInput{
			habitat: "cave", 
			isLegendary: false,}, want: "https://api.funtranslations.com/translate/yoda.json?text=",
		},
		"Get URL for legendary and cave pokemon": {input: TranslationInput{
			habitat: "cave", 
			isLegendary: true,}, want: "https://api.funtranslations.com/translate/yoda.json?text=",
		},
		"Get URL for non-legendary and non-cave pokemon": {input: TranslationInput{
			habitat: "forest", 
			isLegendary: false,}, want: "https://api.funtranslations.com/translate/shakespeare.json?text=",
		},
	}

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			got := getTranslationURL(appCtx, tc.input.habitat, tc.input.isLegendary)
			if got != tc.want {
				t.Fatalf("The expected URL is %s, we got %s", tc.want, got)
			}
        })
    }
}

func TestGetTranslatedDescription(t *testing.T) {
	
	restyClient := resty.New()
	appCtx := initAppCtx(restyClient)

	httpmock.ActivateNonDefault(restyClient.GetClient())
  	defer httpmock.DeactivateAndReset()
	initTranslationResponder(appCtx)

	tests := map[string]struct {
        input TranslationInput
        want string
    }{
        "Get Translated description for legendary pokemon": {input: TranslationInput{
			habitat: "rare", 
			isLegendary: true,
			description: "How are you doing young man"}, want: "You doing young man,  how are",
		},
		"Get Translated description for cave pokemon": {input: TranslationInput{
			habitat: "cave", 
			isLegendary: false,
			description: "How are you doing young man"}, want: "You doing young man,  how are",
		},
		"Get Translated description for non-legendary and non-cave pokemon": {input: TranslationInput{
			habitat: "forest", 
			isLegendary: false,
			description: "How are you doing young man"}, want: "How art thee doing young sir",
		},
		// "Get Translated description when reached rate limit": {input: TranslationInput{
		// 	habitat: "forest", 
		// 	isLegendary: false,
		// 	description: "Rate Limit"}, want: "Rate Limit",
		// },
	}

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			got, err := getTranslatedDescription(appCtx, tc.input.description,tc.input.habitat, tc.input.isLegendary)

			if err != nil {
				t.Fatalf(err.Error())
			}

			if got != tc.want {
				t.Fatalf("The expected translation is %s, we got %s", tc.want, got)
			}
        })
    }
}