package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

// ReferencesClient is an interface to the defined methods for references
type ReferencesClient interface {
	GetAll(ctx context.Context) (*[]Reference, error)

	Create(ctx context.Context, req ReferenceRequest) (*Reference, error)
	Read(ctx context.Context, id string) (*Reference, error)
	Update(ctx context.Context, id string, req ReferenceRequest) (*Reference, error)
	Delete(ctx context.Context, id string) error
}

// RESTReferencesClient implements the ReferencesClient interface
type RESTReferencesClient struct {
	client *APIClient
}

var _ ReferencesClient = &RESTReferencesClient{}

func (c *RESTReferencesClient) restClient() *sling.Sling {
	return c.client.client()
}

// GetAll retrieves a list of references
func (c *RESTReferencesClient) GetAll(ctx context.Context) (*[]Reference, error) {
	res := &[]Reference{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("reference").Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get reference")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Create creates a reference
func (c *RESTReferencesClient) Create(ctx context.Context, req ReferenceRequest) (*Reference, error) {
	res := &Reference{}
	errResponse := new(Error)

	resp, err := c.restClient().Post("reference").
		// BodyJSON(&req).
		BodyProvider(ardoqBodyProvider{request: req, fields: req.Fields}).
		Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not create reference")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Read retrieves a model from the Ardoq API
func (c *RESTReferencesClient) Read(ctx context.Context, id string) (*Reference, error) {
	res := &Reference{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("reference/"+id).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get reference")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Update updates a reference
func (c *RESTReferencesClient) Update(ctx context.Context, id string, req ReferenceRequest) (*Reference, error) {
	// cmp := &ComponentResponse{}
	res := &Reference{}
	errResponse := new(Error)

	// TODO: first receive component body, then update with updated elemenets, POST body

	resp, err := c.restClient().Patch("reference/"+id).
		//BodyJSON(&req).
		BodyProvider(ardoqBodyProvider{request: req, fields: req.Fields}).
		Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not update reference")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// Delete deletes a component
func (c *RESTReferencesClient) Delete(ctx context.Context, id string) error {
	errResponse := new(Error)

	resp, err := c.restClient().Delete("reference/"+id).Receive(nil, errResponse)
	if err != nil {
		return errors.Wrap(err, "could not delete reference")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return errResponse
	}

	return nil
}
