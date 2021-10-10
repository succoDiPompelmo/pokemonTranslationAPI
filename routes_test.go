package main

import (
	"testing"
	"github.com/go-resty/resty/v2"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
)

type PokemonResponse struct {
	Name string
	Description string
	Is_legendary bool
	Habitat string
}

type wantResponse struct {
	StatusCode int
	Pokemon PokemonResponse
}

func TestPokemonBasicRoute(t *testing.T) {

	restyClient := resty.New()
	appCtx := initAppCtx(restyClient)
	pokemonRoutes(appCtx)

	httpmock.ActivateNonDefault(restyClient.GetClient())
  	defer httpmock.DeactivateAndReset()
	initPokemonResponder(appCtx)
	
	tests := map[string]struct {
        input string
        want wantResponse
    }{
        "GET /pokemon with legendary pokemon": {input: "mewtwo", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "mewtwo",
				Description: "How are you doing young man",
				Is_legendary: true,
				Habitat: "rare",
			},
		}},
		"GET /pokemon with non-legendary pokemon": {input: "diglett", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "diglett",
				Description: "How are you doing young man",
				Is_legendary: false,
				Habitat: "cave",
			},
		}},
		"GET /pokemon with non-existing pokemon": {input: "zxcvb", want: wantResponse{
			StatusCode: 404,
			Pokemon: PokemonResponse{
				Name: "",
				Description: "",
				Is_legendary: false,
				Habitat: "",
			},
		}},
		"GET /pokemon with empty pokemon name parameter": {input: "", want: wantResponse{
			StatusCode: 404,
			Pokemon: PokemonResponse{
				Name: "",
				Description: "",
				Is_legendary: false,
				Habitat: "",
			},
		}},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/pokemon/" + tc.input, nil)
			resp, err := appCtx.app.Test(req, 60000) // Set Timeout

			if err != nil {
				t.Fatalf("Error executing the request: %s", err.Error())
			}
			checkResponse(t, resp, tc.want)
        })
    }
}

func TestPokemonTraslationRoute(t *testing.T) {

	restyClient := resty.New()
	appCtx := initAppCtx(restyClient)
	pokemonRoutes(appCtx)

	httpmock.ActivateNonDefault(restyClient.GetClient())
  	defer httpmock.DeactivateAndReset()
	initPokemonResponder(appCtx)
	initTranslationResponder(appCtx)

	tests := map[string]struct {
        input string
        want wantResponse
    }{
        "GET /pokemon/translated with legendary pokemon": {input: "mewtwo", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "mewtwo",
				Description: "You doing young man,  how are",
				Is_legendary: true,
				Habitat: "rare",
			},
		}},
		"GET /pokemon/translated with cave pokemon": {input: "diglett", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "diglett",
				Description: "You doing young man,  how are",
				Is_legendary: false,
				Habitat: "cave",
			},
		}},
		"GET /pokemon/translated with non-legendary non-cave pokemon": {input: "oddish", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "oddish",
				Description: "How art thee doing young sir",
				Is_legendary: false,
				Habitat: "forest",
			},
		}},
		"GET /pokemon/translated with non-existing pokemon": {input: "zxcvb", want: wantResponse{
			StatusCode: 404,
			Pokemon: PokemonResponse{
				Name: "",
				Description: "",
				Is_legendary: false,
				Habitat: "",
			},
		}},
		"GET /pokemon/translated with empty pokemon name parameter": {input: "", want: wantResponse{
			StatusCode: 500,
			Pokemon: PokemonResponse{
				Name: "",
				Description: "",
				Is_legendary: false,
				Habitat: "",
			},
		}},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/pokemon/translated/" + tc.input, nil)
			resp, err := appCtx.app.Test(req, 60000) // Set Timeout

			if err != nil {
				t.Fatalf("Error executing the request: %s", err.Error())
			}
			checkResponse(t, resp, tc.want)
        })
    }
}

func checkResponse(t *testing.T, resp *http.Response, wantResp wantResponse) {

	if resp.StatusCode != wantResp.StatusCode {
		t.Fatalf("Expected status code %d but found %d", wantResp.StatusCode, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        t.Fatalf("Error in reading the response body: %s", err.Error())
    }

	var gotPokemon PokemonResponse
	json.Unmarshal(body, &gotPokemon)

	diff := cmp.Diff(wantResp.Pokemon, gotPokemon)
	if diff != "" {
		t.Fatalf(diff)
	}
}