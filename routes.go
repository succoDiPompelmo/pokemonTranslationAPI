package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"fmt"
)

func pokemonRoutes(appCtx AppCtx) {

	appCtx.app.Use(cache.New())
	
	appCtx.app.Get("/pokemon/:pokemonName", func(c *fiber.Ctx) error {
		pokemonName := c.Params("pokemonName")
		pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, pokemonName)

		if err != nil {
			requestError, ok := err.(*RequestError)
			if ok && requestError.getErrorSatusCode() == 404 {
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
		
		pokemonName := c.Params("pokemonName")
		pokemonSpeciesData, err := getPokemonSpeciesData(appCtx, pokemonName)

		if err != nil {
			requestError, ok := err.(*RequestError)
			if ok && requestError.getErrorSatusCode() == 404{
				return fiber.ErrNotFound
			} else {
				fmt.Println(err.Error())
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