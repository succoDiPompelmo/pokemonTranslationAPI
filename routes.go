package main

import (
	"github.com/gofiber/fiber/v2"
)

func pokemonRoutes(appCtx AppCtx) {
	
	appCtx.app.Get("/pokemon/:pokemonName", func(c *fiber.Ctx) error {
		pokemonName := c.Params("pokemonName")
		pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, pokemonName)

		if err != nil {
			requestError, ok := err.(*RequestError)
			if ok && requestError.getErrorSatusCode() == 404{
				return fiber.ErrNotFound
			} else {
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
		return nil
	})

}