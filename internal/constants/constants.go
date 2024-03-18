package constants

import "time"

type ctxKey string

const (
	KeyRequestInfo ctxKey = "request_info"
	CookieExpire          = 30 * 24 * time.Hour
	Host                  = "http://localhost:8000"
	UserRole              = 0
	AdminRole             = 1
)
