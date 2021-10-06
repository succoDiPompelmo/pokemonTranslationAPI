package main

import (
	"github.com/gofiber/fiber/v2"
)

type AppCtx struct {
	app *fiber.App
}

func main() {
	appCtx := AppCtx{
		app: fiber.New(),
	}
	pokemonRoutes(appCtx)
	appCtx.app.Listen(":3000")
}