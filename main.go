package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-resty/resty/v2"
	"time"
	"log"
)

type AppCtx struct {
	app *fiber.App
	client *resty.Client
	pokemonApiURL string
	translationUrl string
}

func main() {
	
	appCtx := initAppCtx(initRestyClient())
	pokemonRoutes(appCtx)
	
	var err error

	if HTTPS {
		err = appCtx.app.ListenTLS(":3000", "./auth/example.pem", "./auth/example.key")
	} else {
		err = appCtx.app.Listen(":3000")
	}

	if err != nil {
		log.Printf("ERROR in initializing the app, the following message is returned %s", err.Error())
	}
}

func initAppCtx(client *resty.Client) AppCtx {
	return AppCtx{
		app: fiber.New(),
		client: client,
		pokemonApiURL: POKEMON_API_URL,
		translationUrl: FUN_TRANSLATION_API_URL,
	}
}

func initRestyClient() *resty.Client {
	restyClient := resty.New()
	restyClient.SetTimeout(TIMEOUT * time.Second)
	return restyClient
}