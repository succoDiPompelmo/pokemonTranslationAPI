package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-resty/resty/v2"
	"fmt"
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
	
	var err error

	if HTTPS {
		err = appCtx.app.ListenTLS(":3000", "./auth/example.pem", "./auth/example.key")
	} else {
		err = appCtx.app.Listen(":3000")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
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