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