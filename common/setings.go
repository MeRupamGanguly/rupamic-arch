package common

import "time"

const (
	ApiRateLimitKey         = "rupam_rate_limiting"
	ApiRateLimitMaxRequests = 10
	ApiRateLimitDuration    = time.Minute * 5
)

// Roles
const (
	ADMINROLE      = "ADMIN"
	SUPERADMINROLE = "SUPER"
	USERROLE       = "USER"
	GUESTROLE      = "GUEST"
)

var APIKEYS map[string]string

func APIKEYS_INIT() {
	APIKEYS = make(map[string]string)
}

func APIKeyAddedInDB(key, val string) {
	APIKEYS[key] = val
}
func GetAPIKeyFromDB(key string) string {
	return APIKEYS[key]
}
