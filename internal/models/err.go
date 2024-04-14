package models

import "errors"

var (
	ErrNoDocument error = errors.New("no user found with the given name")

	ErrInsufficientQueryArg error = errors.New("no required query arg")
)

type ErrResponse struct {
	Msg string `json:"msg"`
}
