package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
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
