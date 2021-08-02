package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

type ModelsClient interface {
	GetAll(ctx context.Context) (*[]Model, error)
	Read(ctx context.Context, id string) (*Model, error)
}

// RESTModelClient implements the ModelClient interface
type RESTModelsClient struct {
	client *APIClient
}

var _ ModelsClient = &RESTModelsClient{}

func (c *RESTModelsClient) restClient() *sling.Sling {
	return c.client.client()
}

// List retrieves a list of models from the Ardoq API
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

// Get retrieves a model from the Ardoq API
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
