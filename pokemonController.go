package main

import (
	"fmt"
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

func getPokemonSpeciesData(appCtx *AppCtx, pokemonName string) (*PokemonSpecies, error) {
	resp, err := appCtx.client.R().SetResult(&PokemonSpecies{}).Get("https://pokeapi.co/api/v2/pokemon-species" + "/" + pokemonName)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("The Poke API Request returned a status code %d", resp.StatusCode())
	}
    return resp.Result().(*PokemonSpecies), err
}

func (pokemonSpecies PokemonSpecies) getDescription() string {
	return ""
}

func (pokemonSpecies PokemonSpecies) getHabitat() string {
	return ""
}

func (pokemonSpecies PokemonSpecies) isLegendary() bool {
	return false
}