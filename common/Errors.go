package common

import "errors"

var (
	ErrUserCredWrong = errors.New("user credentails mismatched, please use correct username password")
	ErrRateLimiting  = errors.New("you are sending too many reqyests, please wait sometime")
	ErrAPIKey        = errors.New("api key is missing")
)
