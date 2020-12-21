package access_token_middleware

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var (
	cfg = RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}
	accessTokenStoreTest AccessTokenStore
	err                  error
	key                  = "key_test"
	value                = "value"
)

func TestNewAccessTokenStore(t *testing.T) {
	accessTokenStoreTest, err = NewAccessTokenStore(cfg)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s is ended successfully \n", t.Name())
	fmt.Println("access token store connected")
}

func TestAccessTokenStore_Save(t *testing.T) {
	err = accessTokenStoreTest.Save(value, key, time.Duration(500000000000))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s is ended successfully\n", t.Name())
	fmt.Printf("userID-> %s is saved by key-> %s\n", value, key)
}

func TestAccessTokenStore_Get(t *testing.T) {
	userID, err := accessTokenStoreTest.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is ended successfully\n", t.Name())
	fmt.Printf("userID-> %s is getted by key-> %s\n", userID, key)
}
