package main

import (

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
	Url string
}

func getPokemonSpeciesData(appCtx *AppCtx, pokemonName string) (*PokemonSpecies, error) {
    return nil, nil
}