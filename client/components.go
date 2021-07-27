package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

type ComponentsClient interface {
	//Get(ctx context.Context, id string) (*Component, error)
	Search(ctx context.Context, req *ComponentSearchQuery) (*[]Component, error)
	GetAll(ctx context.Context) (*[]Component, error)
	//
	Create(ctx context.Context, req ComponentRequest) (*Component, error)
	Read(ctx context.Context, id string) (*Component, error)
	Update(ctx context.Context, id string, req ComponentRequest) (*Component, error)
	Delete(ctx context.Context, id string) error
}

// RESTComponentsClient implements the WorkspacesClient interface
type RESTComponentsClient struct {
	client *APIClient
}

var _ ComponentsClient = &RESTComponentsClient{}

func (c *RESTComponentsClient) restClient() *sling.Sling {
	return c.client.client()
}

// List retrieves a list of components from the Ardoq API
// TODO: Check failure case
func (c *RESTComponentsClient) Search(ctx context.Context, req *ComponentSearchQuery) (*[]Component, error) {
	res := &[]Component{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("component/search").QueryStruct(req).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not find component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// List retrieves a list of components based on a component query,
// is not being used yet
func (c *RESTComponentsClient) GetAll(ctx context.Context) (*[]Component, error) {
	res := &[]Component{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("component").Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

//////

// ComponentCreate creates an component
func (c *RESTComponentsClient) Create(ctx context.Context, req ComponentRequest) (*Component, error) {
	res := &Component{}
	errResponse := new(Error)

	resp, err := c.restClient().Post("component").
		//BodyJSON(&req).
		BodyProvider(ardoqBodyProvider{request: req, fields: req.Fields}).
		Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not create component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// ComponentRead retrieves an component from the Ardoq API
func (c *RESTComponentsClient) Read(ctx context.Context, id string) (*Component, error) {
	res := &Component{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("component/"+id).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// ComopnentUpdate updates a component
func (c *RESTComponentsClient) Update(ctx context.Context, id string, req ComponentRequest) (*Component, error) {
	// cmp := &ComponentResponse{}
	res := &Component{}
	errResponse := new(Error)

	// TODO: first receive component body, then update with updated elemenets, POST body

	resp, err := c.restClient().Patch("component/"+id).
		//BodyJSON(&req).
		BodyProvider(ardoqBodyProvider{request: req, fields: req.Fields}).
		Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not update component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// ComponentDelete deletes a component
func (c *RESTComponentsClient) Delete(ctx context.Context, id string) error {
	errResponse := new(Error)

	resp, err := c.restClient().Delete("component/"+id).Receive(nil, errResponse)
	if err != nil {
		return errors.Wrap(err, "could not delete component")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return errResponse
	}

	return nil
}
