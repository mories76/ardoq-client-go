package ardoq

import (
	"context"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

type WorkspacesClient interface {
	Get(ctx context.Context, id string) (*Workspace, error)
	Search(ctx context.Context, req *WorkspaceSearchQuery) (*Workspace, error)
	List(ctx context.Context, req *WorkspaceSearchQuery) (*[]Workspace, error)
}

// RESTWorkspacesClient implements the WorkspacesClient interface
type RESTWorkspacesClient struct {
	client *APIClient
}

var _ WorkspacesClient = &RESTWorkspacesClient{}

func (c *RESTWorkspacesClient) restClient() *sling.Sling {
	return c.client.client()
}

// Get retrieves a workspace from the Ardoq API
// TODO: Check failure case
func (c *RESTWorkspacesClient) Get(ctx context.Context, id string) (*Workspace, error) {
	res := &Workspace{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("workspaces/"+id).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get workspace")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// List retrieves a list of workspaces from the Ardoq API
func (c *RESTWorkspacesClient) Search(ctx context.Context, req *WorkspaceSearchQuery) (*Workspace, error) {
	res := &Workspace{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("workspace/search").QueryStruct(req).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get workspace")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}

// List retrieves a list of services based on a service query
func (c *RESTWorkspacesClient) List(ctx context.Context, req *WorkspaceSearchQuery) (*[]Workspace, error) {
	res := &[]Workspace{}
	errResponse := new(Error)

	resp, err := c.restClient().Get("workspace").QueryStruct(req).Receive(res, errResponse)
	if err != nil {
		return nil, errors.Wrap(err, "could not get workspace")
	}
	if errResponse.NotOk() {
		errResponse.Code = resp.StatusCode
		return nil, errResponse
	}

	return res, nil
}
