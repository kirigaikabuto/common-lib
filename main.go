package main

import (
	"fmt"
	acs "github.com/kirigaikabuto/common-lib/access-token-middleware"
	"log"
	"time"
)

func main() {
	cfg := acs.RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}
	accessTokenStore, err := acs.NewAccessTokenStore(cfg)
	if err != nil {
		log.Fatal("connection error", err)
	}
	err = accessTokenStore.Save("123", "456", time.Duration(10000000))
	if err != nil {
		log.Fatal(err)
	}
	userID, err := accessTokenStore.Get("456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userID)
}
