package main

import (

)

type Response struct {
	responseBody []byte
	statusCode int
}

func doGet(appCtx AppCtx, url string) (*Response, error) {
	resp, err := appCtx.client.R().Get(url)
	response := &Response{
		responseBody: resp.Body(),
		statusCode: resp.StatusCode(),
	}
	return response, err
}