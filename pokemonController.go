package main

import (
	"strings"
	"log"
)

type PokemonSpecies struct {
	Is_legendary bool
	Name string
	Flavor_text_entries []FlavorTextEntry
	Habitat HabitatResource
}

type HabitatResource struct {
	Name string
}

type FlavorTextEntry struct {
	Flavor_text string
	Language LanguageResource
}

type LanguageResource struct {
	Name string
}

func getPokemonSpeciesData(appCtx AppCtx, pokemonName string) (*PokemonSpecies, error) {
	resp, err := appCtx.client.R().SetResult(&PokemonSpecies{}).Get(appCtx.pokemonApiURL + pokemonName)
	if err != nil {
		log.Printf("GET POKEMON SPECIES ERROR: %s for pokemon %s and obtained response %s", err.Error(), pokemonName, string(resp.Body()))
		return nil, err
	}
	if resp.StatusCode() > 399 {
		log.Printf("GET POKEMON SPECIES ERROR: for pokemon %s and obtained status code %d and response %s", pokemonName, resp.StatusCode(), string(resp.Body()))
		return nil, &RequestError{StatusCode: resp.StatusCode(), Err: err,}
	}
    return resp.Result().(*PokemonSpecies), err
}

func (pokemonSpecies PokemonSpecies) getDescription() string {
	description := ""
	for _, flavorText := range pokemonSpecies.Flavor_text_entries {
		if flavorText.Language.Name == "en" {
			description = strings.ReplaceAll(strings.ReplaceAll(flavorText.Flavor_text, "\n", " "), "\f", " ")
			return description
		}
	}
	return description
}

func (pokemonSpecies PokemonSpecies) getHabitat() string {
	return pokemonSpecies.Habitat.Name
}

func (pokemonSpecies PokemonSpecies) isLegendary() bool {
	return pokemonSpecies.Is_legendary
}