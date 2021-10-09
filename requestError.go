package main

import (
	"fmt"
)

type RequestError struct {
    StatusCode int
    Err error
}

func (r *RequestError) Error() string {
    return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func (r *RequestError) getErrorSatusCode() int {
    return r.StatusCode
}