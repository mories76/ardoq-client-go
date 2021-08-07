package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

// FieldsClient is an interface to the defined methods for components
type FieldsClient interface {
	GetAll(ctx context.Context) (*[]Field, error)
	Read(ctx context.Context, id string) (*Field, error)
}

// RESTFieldsClient implements the FieldsClient interface
type RESTFieldsClient struct {
	client *APIClient
}

var _ FieldsClient = &RESTFieldsClient{}

func (c *RESTFieldsClient) restClient() *sling.Sling {
	return c.client.client()
}

// GetAll retrieves a list of all fields
func (c *RESTFieldsClient) GetAll(ctx context.Context) (*[]Field, error) {
	res := &[]Field{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("field").Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get field")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Read retrieves a field by its ID
func (c *RESTFieldsClient) Read(ctx context.Context, id string) (*Field, error) {
	res := &Field{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("field/"+id).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get field")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}
