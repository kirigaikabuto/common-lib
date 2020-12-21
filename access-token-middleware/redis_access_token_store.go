package access_token_middleware

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type accessTokenStore struct {
	cfg RedisConfig
	clt *redis.Client
}

func NewAccessTokenStore(cfg RedisConfig) (AccessTokenStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port),
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &accessTokenStore{cfg, client}, nil
}

func (acs *accessTokenStore) Save(userId, token string, ttl time.Duration) error {
	data, err := json.Marshal(userId)
	if err != nil {
		return err
	}
	_, err = acs.clt.Set("user_id:"+token, data, ttl).Result()
	return err
}

func (acs *accessTokenStore) Get(token string) (string, error) {
	val, err := acs.clt.Get("user_id:" + token).Bytes()
	if err != nil {
		return "", err
	}
	userId := ""
	if err := json.Unmarshal(val, &userId); err != nil {
		return "", err
	}
	return userId, nil
}
