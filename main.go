package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-resty/resty/v2"
)

type AppCtx struct {
	app *fiber.App
	client *resty.Client
	pokemonApiURL string
	translationUrl string
	timeout int
}

func main() {
	
	appCtx := initAppCtx(resty.New())
	pokemonRoutes(appCtx)
	appCtx.app.Listen(":3000")
}

func initAppCtx(client *resty.Client) AppCtx {
	return AppCtx{
		app: fiber.New(),
		client: client,
		pokemonApiURL: POKEMON_API_URL,
		translationUrl: FUN_TRANSLATION_API_URL,
		timeout: TIMEOUT,
	}
}