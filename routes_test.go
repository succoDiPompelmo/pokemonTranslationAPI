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
	"github.com/google/go-cmp/cmp/cmpopts"
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
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "mewtwo",
				Description: "",
				Is_legendary: true,
				Habitat: "rare",
			},
		}},
		"GET /pokemon with non-legendary pokemon": {input: "diglett", want: wantResponse{
			StatusCode: 200,
			Pokemon: PokemonResponse{
				Name: "diglett",
				Description: "",
				Is_legendary: false,
				Habitat: "cave",
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

func checkResponse(t *testing.T, resp *http.Response, wantResp wantResponse) {

	if resp.StatusCode != wantResp.StatusCode {
		t.Fatalf("Expected status code %d but found %d", resp.StatusCode, wantResp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        t.Fatalf("Error in reading the response body: %s", err.Error())
    }

	var gotPokemon PokemonResponse
	json.Unmarshal(body, &gotPokemon)

	opts := []cmp.Option{
		cmpopts.IgnoreFields(PokemonResponse{}, "Description"),
	}
	diff := cmp.Diff(wantResp.Pokemon, gotPokemon, opts...)
	if diff != "" {
		t.Fatalf(diff)
	}
}