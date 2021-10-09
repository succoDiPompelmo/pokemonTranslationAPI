package main

import (
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/go-resty/resty/v2"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
)

type PokemonResponse struct {
	name string
	description string
	is_legendary bool
	habitat string
}

type wantResponse struct {
	statusCode int
	pokemon PokemonResponse
}

func TestPokemonBasicRoute(t *testing.T) {

	appCtx := AppCtx{
		app: fiber.New(),
		client: resty.New(),
	}

	pokemonRoutes(appCtx)

	tests := map[string]struct {
        input string
        want wantResponse
    }{
        "GET /pokemon with legendary pokemon": {input: "mewtwo", want: wantResponse{
			statusCode: 200,
			pokemon: PokemonResponse{
				name: "mewtwo",
				description: "",
				is_legendary: true,
				habitat: "rare",
			},
		}},
		"GET /pokemon with non-legendary pokemon": {input: "diglett", want: wantResponse{
			statusCode: 200,
			pokemon: PokemonResponse{
				name: "diglett",
				description: "",
				is_legendary: false,
				habitat: "cave",
			},
		}},
		"GET /pokemon with empty pokemon name parameter": {input: "", want: wantResponse{
			statusCode: 404,
			pokemon: PokemonResponse{
				name: "",
				description: "",
				is_legendary: false,
				habitat: "",
			},
		}},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/pokemon/" + tc.input, nil)
			resp, err := appCtx.app.Test(req, 5)

			if err != nil {
				t.Fatalf("Error executing the request: %s", err.Error())
			}
			checkResponse(t, resp, tc.want)
        })
    }
}

func checkResponse(t *testing.T, resp *http.Response, wantResp wantResponse) {

	if resp.StatusCode != wantResp.statusCode {
		t.Fatalf("Expected status code %d but found %d", resp.StatusCode, wantResp.statusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        t.Fatalf("Error in reading the response body: %s", err.Error())
    }

	var gotPokemon PokemonResponse
	json.Unmarshal(body, &gotPokemon)

	diff := cmp.Diff(wantResp.pokemon, gotPokemon)
	if diff != "" {
		t.Fatalf(diff)
	}
}