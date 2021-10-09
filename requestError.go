package main

import (
	"fmt"
)

type RequestError struct {
    statusCode int
    err error
}

func (r *RequestError) Error() string {
    return fmt.Sprintf("status %d: err %v", r.statusCode, r.err)
}

func (r *RequestError) getErrorSatusCode() int {
    return r.statusCode
}