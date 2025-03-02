package common

import (
	"fmt"

	"strings"
	"time"
)

func GenerateApiKey(role string) string {
	now := time.Now().UnixNano()
	date := time.Now().Format("DDMMYYYY")
	hashed := Encrypt(strings.ToUpper(role))
	apikey := fmt.Sprintf("%v-%v-%v", now, date, hashed)
	return apikey
}

func CreateAPIKey(roles []string) string {
	isAdmin := false
	isSuper := false
	isUser := false
	isGuest := false
	for _, role := range roles {
		if role == ADMINROLE {
			isAdmin = true
		}
		if role == SUPERADMINROLE {
			isSuper = true
		}
		if role == USERROLE {
			isUser = true
		}
		if role == GUESTROLE {
			isGuest = true
		}
	}
	if isSuper {
		key := GenerateApiKey(SUPERADMINROLE)
		APIKeyAddedInDB(key, "SUPERADMINROLE")
		return key
	}
	if isAdmin {
		key := GenerateApiKey(ADMINROLE)
		APIKeyAddedInDB(key, "ADMINROLE")
		return key
	}
	if isGuest {
		key := GenerateApiKey(GUESTROLE)
		APIKeyAddedInDB(key, "GUESTROLE")
		return key
	}
	if isUser {
		key := GenerateApiKey(USERROLE)
		APIKeyAddedInDB(key, "USERROLE")
		return key
	}
	return ""
}
