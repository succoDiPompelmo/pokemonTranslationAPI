package main

import (
	"github.com/gofiber/fiber/v2"
)

func pokemonRoutes(appCtx AppCtx) {
	
	appCtx.app.Get("/pokemon/:pokemonName", func(c *fiber.Ctx) error {
		return nil
	})

	appCtx.app.Get("/pokemon/translated/:pokemonName", func(c *fiber.Ctx) error {
		return nil
	})

}