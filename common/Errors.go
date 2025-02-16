package common

import "errors"

var (
	ErrUserCredWrong = errors.New("user credentails mismatched, please use correct username password")
)
