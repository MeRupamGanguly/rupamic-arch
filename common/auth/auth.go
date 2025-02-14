package auth

import "github.com/golang-jwt/jwt/v5"

var Seckey = "007JamesBond"

type Claims struct {
	UserId string   `json:"userId"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func CreateToken(id string, roles []string) (token string, err error) {
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{UserId: id, Roles: roles})
	token, err = tokenStr.SignedString([]byte(Seckey))
	if err != nil {
		return "", err
	}
	return token, nil
}
