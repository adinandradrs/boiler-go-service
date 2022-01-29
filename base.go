package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"crypto/tls"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/adinandradrs/codefun-go-service/util"
	"github.com/go-redis/redis"
)

type RestAdaptorCapsule struct {
	Client  *http.Client
	ApiUrl  string
	Method  string
	Header  http.Header
	Payload interface{}
}

type RestAdaptor interface {
	Execute(inp interface{}) (out interface{}, err error)
}

func NewRestAdaptor(client *http.Client) *RestAdaptorCapsule {
	return &RestAdaptorCapsule{
		Client: client,
	}
}

func (c *RestAdaptorCapsule) GeneratePayload() (*http.Request, error) {
	var pyld *bytes.Buffer = nil
	if c.Payload != nil {
		pyld = new(bytes.Buffer)
		err := json.NewEncoder(pyld).Encode(c.Payload)
		if err != nil {
			return nil, err
		}
	}
	req, reqerr := http.NewRequest(c.Method, c.ApiUrl, pyld)
	if reqerr != nil {
		return nil, fmt.Errorf("request is unable to proceed = %s", reqerr)
	}
	if c.Header != nil {
		req.Header = c.Header
	}
	return req, nil
}

type BaseRepositoryCapsule struct {
	Database *pgxpool.Pool
}

func NewBaseRepository(database *pgxpool.Pool) *BaseRepositoryCapsule {
	return &BaseRepositoryCapsule{
		Database: database,
	}
}

type BaseRepository interface {
	Save(input interface{}) (interface{}, error)
	FindById(id uint) (interface{}, error)
	Delete(id uint) error
}

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
