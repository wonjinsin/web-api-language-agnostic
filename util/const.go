package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Const ...
const (
	DBCharsetOption string = "DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci"
)

// CustomStr ...
type CustomStr string

// CustomStrs ...
const (
	TRID  CustomStr = "trid"
	TxKey CustomStr = "tx"
)

// CustomTimes ...
const (
	CtxTimeOut = time.Second * 10
)

// TokenTypes ...
const (
	TokenTypeBearer string = "Bearer"
)

// TokenAudiences ...
const (
	TokenAudienceAccount string = "account"
)

// ConfigKeys ...
const (
	ConfigPubTokenKey string = "pubTokenKey"
	ConfigPrvTokenKey string = "prvTokenKey"
)

// GetTRID ...
func GetTRID() string {
	t := time.Now()
	randInt := strconv.Itoa(rand.Intn(8999) + 1000)
	trid := strings.Replace(t.Format("20060102150405.00"), ".", "", -1) + randInt

	return trid
}

type contextKey string

// ContextKey ...
const (
	LoginKey contextKey = "login"
	UUID     contextKey = "uuid"
)

// CountryCode is for country code
type CountryCode string

// CountryCodeConst is for country code constant
const (
	CountryCodeJP CountryCode = "JP"
)
