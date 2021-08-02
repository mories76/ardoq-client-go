package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

// ModelsClient is an interface to the defined methods for components
type ModelsClient interface {
	GetAll(ctx context.Context) (*[]Model, error)
	Read(ctx context.Context, id string) (*Model, error)
}

// RESTModelsClient implements the ModelsClient interface
type RESTModelsClient struct {
	client *APIClient
}

var _ ModelsClient = &RESTModelsClient{}

func (c *RESTModelsClient) restClient() *sling.Sling {
	return c.client.client()
}

// GetAll retrieves a list of all models
func (c *RESTModelsClient) GetAll(ctx context.Context) (*[]Model, error) {
	res := &[]Model{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("model").Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get model")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Read retrieves a model by its ID
func (c *RESTModelsClient) Read(ctx context.Context, id string) (*Model, error) {
	res := &Model{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("model/"+id).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get model")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}
