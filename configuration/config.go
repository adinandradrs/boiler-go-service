package configuration

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/adinandradrs/codefun-go-service/util"
	"github.com/go-redis/redis"
)

func ConfigRestClient(timeout int) *http.Client {
	if os.Getenv("REST_SKIP_SSL") == util.VALUE_YES {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	httpClient := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &httpClient
}

func ConfigCache(host string, passwd string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     passwd,
		DB:           0,
		PoolSize:     50,
		MinIdleConns: 10,
	})
	return rdb
}
