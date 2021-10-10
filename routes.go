package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"log"
)

func pokemonRoutes(appCtx AppCtx) {

	appCtx.app.Use(cache.New())
	
	appCtx.app.Get("/pokemon/:pokemonName", func(c *fiber.Ctx) error {
		pokemonName := c.Params("pokemonName")
		pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, pokemonName)

		if err != nil {
			requestError, ok := err.(*RequestError)
			if ok && requestError.getErrorSatusCode() == 404 {
				log.Printf("GET POKEMON ERROR with the following description %s and pokemon name %s", err.Error(), pokemonName)
				return fiber.ErrNotFound
			} else {
				log.Printf("GET POKEMON ERROR with the following description %s and pokemon name %s", err.Error(), pokemonName)
				return fiber.ErrInternalServerError
			}
		}

		return c.JSON(fiber.Map{
			"name": pokemonName,
			"description": pokemonSpeciesData.getDescription(),
			"habitat": pokemonSpeciesData.getHabitat(),
			"Is_legendary": pokemonSpeciesData.isLegendary(),
		})
	})

	appCtx.app.Get("/pokemon/translated/:pokemonName", func(c *fiber.Ctx) error {
		
		pokemonName := c.Params("pokemonName")
		pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, pokemonName)

		if err != nil {
			requestError, ok := err.(*RequestError)
			if ok && requestError.getErrorSatusCode() == 404{
				log.Printf("GET POKEMON ERROR with the following description %s and pokemon name %s", err.Error(), pokemonName)
				return fiber.ErrNotFound
			} else {
				log.Printf("GET POKEMON ERROR with the following description %s and pokemon name %s", err.Error(), pokemonName)
				return fiber.ErrInternalServerError
			}
		}

		trasnlatedDescription, err := getTranslatedDescription(
			appCtx, 
			pokemonSpeciesData.getDescription(),
			pokemonSpeciesData.getHabitat(),
			pokemonSpeciesData.isLegendary(),
		)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(fiber.Map{
			"name": pokemonName,
			"description": trasnlatedDescription,
			"habitat": pokemonSpeciesData.getHabitat(),
			"Is_legendary": pokemonSpeciesData.isLegendary(),
		})
	})

}