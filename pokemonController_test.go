package main

import (
	"testing"
	"github.com/go-resty/resty/v2"
)

func TestGetPokemonSpeciesDataValidPokemon(t *testing.T) {
	appCtx := &AppCtx{
		client: resty.New(),
	}
	pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, "mewtwo")
	if err != nil {
		t.Errorf("Expected no Errors but found: %s", err.Error())
	}

	if pokemonSpeciesData == nil {
		t.Fatalf("Expected pokemon data but none found") 
	}

	if pokemonSpeciesData.Name != "mewtwo" {
		t.Errorf("Expected name mewtwo but found: %s", pokemonSpeciesData.Name) 
	}

	if pokemonSpeciesData.Habitat.Name != "rare" {
		t.Errorf("Mewtwo habitat expected to be rare, found: %s", pokemonSpeciesData.Habitat.Name) 
	}

	if !pokemonSpeciesData.Is_legendary {
		t.Errorf("Mewtwo found to be non legendary, expected legendary") 
	}
}

func TestGetPokemonSpeciesDataInvalidPokemon(t *testing.T) {
	appCtx := &AppCtx{
		client: resty.New(),
	}
	pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, "paolo")

	if pokemonSpeciesData != nil {
		t.Errorf("Expected no data to be returned with invalid pokemon name")
	}

	if err == nil {
		t.Errorf("Expected Errors with invalid pokemon name but none found")
	}
}

func TestGetPokemonSpeciesDataConnectionError(t *testing.T) {

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
        "Standard Habitat Name": {input: &PokemonSpecies{Is_legendary: true}, want: true},
        "Empty Habitat Name": {input: &PokemonSpecies{Is_legendary: false}, want: false},
        "No Habitat Resource": {input: &PokemonSpecies{}, want: false},
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
			FlavorTextEntry{Flavor_text: "", Language: LanguageResource{Name: "jp"}},
		}}, want: ""},
		"No Descriptions": {input: &PokemonSpecies{}, want: ""},
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