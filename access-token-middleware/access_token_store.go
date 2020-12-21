package access_token_middleware

import "time"

type AccessTokenStore interface {
	Save(userId, key string, ttl time.Duration) error
	Get(token string) (string, error)
}
