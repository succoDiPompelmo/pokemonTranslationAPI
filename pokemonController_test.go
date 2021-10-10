package main

import (
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-cmp/cmp"
)

func TestGetPokemonSpeciesDataValidPokemon(t *testing.T) {

	restyClient := resty.New()
	appCtx := initAppCtx(restyClient)

	httpmock.ActivateNonDefault(restyClient.GetClient())
  	defer httpmock.DeactivateAndReset()
	initPokemonResponder(appCtx)

	tests := map[string]struct {
        input string
        want PokemonSpecies
    }{
        "Get pokemon species data for legendary pokemon": {input: "mewtwo", want: PokemonSpecies{
			Name: "mewtwo",
			Flavor_text_entries: []FlavorTextEntry{
				FlavorTextEntry{Flavor_text: "How are you doing young man", Language: LanguageResource{Name: "en"}},
			},
			Is_legendary: true,
			Habitat: HabitatResource{Name: "rare"},
		}},
		"Get pokemon species data for non-legendary pokemon": {input: "diglett", want: PokemonSpecies{
			Name: "diglett",
			Flavor_text_entries: []FlavorTextEntry{
				FlavorTextEntry{Flavor_text: "How are you doing young man", Language: LanguageResource{Name: "en"}},
			},
			Is_legendary: false,
			Habitat: HabitatResource{Name: "cave"},
		}},
	}

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			got, err := getPokemonSpeciesData(appCtx, tc.input)
			if err != nil {
				t.Fatalf("Error executing the request: %s", err.Error())
			}
			diff := cmp.Diff(tc.want, *got)
			if diff != "" {
				t.Fatalf(diff)
			}
        })
    }
}

func TestGetPokemonSpeciesDataInvalidPokemon(t *testing.T) {

	restyClient := resty.New()
	appCtx := initAppCtx(restyClient)

	httpmock.ActivateNonDefault(restyClient.GetClient())
  	defer httpmock.DeactivateAndReset()
	initPokemonResponder(appCtx)

	tests := map[string]struct {
        input string
        want *RequestError
    }{
        "Get pokemon species data for invalid pokemon": {input: "zxcvb", want: &RequestError{
			Err: nil,
			StatusCode: 404,
		}},
		"Get pokemon species data with internal server error": {input: "internalServerError", want: &RequestError{
			Err: nil,
			StatusCode: 500,
		}},
	}

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
			got, err := getPokemonSpeciesData(appCtx, tc.input)
			if got != nil {
				t.Fatalf("Expected error but got pokemon data")
			}
			diff := cmp.Diff(tc.want, err)
			if diff != "" {
				t.Fatalf(diff)
			}
        })
    }
}

func TestPokemonSpeciesGetHabitat(t *testing.T) {
	tests := map[string]struct {
        input *PokemonSpecies
        want  string
    }{
        "Standard Habitat Name": {input: &PokemonSpecies{Habitat: HabitatResource{Name: "cave"}}, want: "cave"},
        "Empty Habitat Name": {input: &PokemonSpecies{Habitat: HabitatResource{Name: ""}}, want: ""},
        "Default Habitat Name": {input: &PokemonSpecies{}, want: ""},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := tc.input.getHabitat()
            if got != tc.want {
                t.Fatalf("Expected Habitat %s and found habitat %s", tc.want, got)
            }
        })
    }
}

func TestPokemonSpeciesIsLegendary(t *testing.T) {
	tests := map[string]struct {
        input *PokemonSpecies
        want  bool
    }{
        "Legendary Pokemon": {input: &PokemonSpecies{Is_legendary: true}, want: true},
        "Non Legendary Pokemon": {input: &PokemonSpecies{Is_legendary: false}, want: false},
        "Default Pokemon": {input: &PokemonSpecies{}, want: false},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := tc.input.isLegendary()
            if got != tc.want {
				if tc.want {
					t.Fatalf("Expected Pokemon to be legendary but it's not")
				} else {
					t.Fatalf("Expected Pokemon to not be legendary but it was")
				}
                
            }
        })
    }
}

func TestPokemonSpeciesGetDescription(t *testing.T) {
	tests := map[string]struct {
        input *PokemonSpecies
        want  string
    }{
        "Standard English Description": {input: &PokemonSpecies{Flavor_text_entries: []FlavorTextEntry{
			FlavorTextEntry{Flavor_text: "A", Language: LanguageResource{Name: "en"}},
		}}, want: "A"},
        "English and Non english Description": {input: &PokemonSpecies{Flavor_text_entries: []FlavorTextEntry{
			FlavorTextEntry{Flavor_text: "A", Language: LanguageResource{Name: "jp"}},
			FlavorTextEntry{Flavor_text: "B", Language: LanguageResource{Name: "en"}},
		}}, want: "B"},
        "No English Description": {input: &PokemonSpecies{Flavor_text_entries: []FlavorTextEntry{
			FlavorTextEntry{Flavor_text: "A", Language: LanguageResource{Name: "jp"}},
		}}, want: ""},
		"Empty Description": {input: &PokemonSpecies{Flavor_text_entries: []FlavorTextEntry{
			FlavorTextEntry{Flavor_text: "", Language: LanguageResource{Name: "en"}},
		}}, want: ""},
		"No Descriptions": {input: &PokemonSpecies{}, want: ""},
		"Description with escaped sequences": {input: &PokemonSpecies{Flavor_text_entries: []FlavorTextEntry{
			FlavorTextEntry{Flavor_text: "A\nB\f\n\fC", Language: LanguageResource{Name: "en"}},
		}}, want: "A B   C"},
    }

	for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := tc.input.getDescription()
            if got != tc.want {
				t.Fatalf("Expected Description %s and found description %s", tc.want, got)
            }
        })
    }
}