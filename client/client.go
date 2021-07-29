package ardoq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/mitchellh/mapstructure"
)

const (
	// UserAgentPrefix is the prefix of the User-Agent header that all terraform REST calls perform
	UserAgentPrefix = "terraform-provider-ardoq"
)

// Custom ardoq error response handler
type Error struct {
	Code    int
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (err Error) Error() string {
	if err.Code == 0 {
		return ""
		// return nil
	} else {
		return fmt.Sprintf("ardoq error: \nstatuscode :%d\n%s %v", err.Code, err.Message, err.Data)
	}
}

func (err Error) Ok() bool {
	return err.Code == 0
}

func (err Error) NotOk() bool {
	return err.Code != 0
}

// APIClient is the client that accesses all of the ardoq resources
type APIClient struct {
	baseURI string
	apiKey  string
	org     string
	version string
}

var _ Client = &APIClient{}

type Client interface {
	Components() ComponentsClient
	Models() ModelsClient
	References() ReferencesClient
	Workspaces() WorkspacesClient
}

// OptFunc is a function that sets a setting on a client
type OptFunc func(c *APIClient) error

// WithBaseURL modifies the base URL for all requests
// func WithBaseURL(baseURL string) OptFunc {
// 	return func(c *APIClient) error {
// 		c.baseURI = baseURI
// 		return nil
// 	}
// }

// NewRestClient initializes a new API client for Ardoq
func NewRestClient(baseURI, apiKey, org, version string, opts ...OptFunc) (*APIClient, error) {
	c := &APIClient{
		baseURI: baseURI,
		apiKey:  apiKey,
		org:     org,
		version: version,
	}

	for _, f := range opts {
		if err := f(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *APIClient) client() *sling.Sling {
	type OrgSearchQuery struct {
		Org string `url:"org,omitempty"`
	}

	return sling.New().Base(c.baseURI).
		Set("User-Agent", fmt.Sprintf("%s (%s)", UserAgentPrefix, c.version)).
		Set("Authorization", fmt.Sprintf("Token token=%s", c.apiKey)).ResponseDecoder(ardoqDecoder{}).
		QueryStruct(&OrgSearchQuery{Org: c.org})
	// Doer(ardoqDoer{})
}

// Components returns a ComponentsClient interface for interacting with Components in Ardoq
func (c *APIClient) Components() ComponentsClient {
	return &RESTComponentsClient{client: c}
}

// Models returns a ModelsClient interface for interacting with Models in Ardoq
func (c *APIClient) Models() ModelsClient {
	return &RESTModelsClient{client: c}
}

// References returns a ReferencesClient interface for interacting with References in Ardoq
func (c *APIClient) References() ReferencesClient {
	return &RESTReferencesClient{client: c}
}

// Workspaces returns a WorkspacesClient interface for interacting with Workspaces in Ardoq
func (c *APIClient) Workspaces() WorkspacesClient {
	return &RESTWorkspacesClient{client: c}
}

// custom decoder toegevoegd, het doel is om alle "unknown" fields uit de json op te staan in een struct
// daarvoor wordt mapstructure gebruikt
type ardoqDecoder struct {
}

func (a ardoqDecoder) Decode(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}

	/*err :=*/
	json.Unmarshal(body, &data)

	// check if StatusCode is OK, if not add StatusCode to the Error
	// so that ErrorReponse.Errors() return true, and the actual decoded respone get shown in terraform
	if code := resp.StatusCode; 200 <= code && code <= 299 {
	} else {
		// apparently there's an error,
		var errResponse *Error = v.(*Error)
		errResponse.Code = resp.StatusCode
	}

	return mapstructure.WeakDecode(data, v)
}

// jsonBodyProvider encodes a JSON tagged struct value as a Body for requests.
// See https://golang.org/pkg/encoding/json/#MarshalIndent for details.
type ardoqBodyProvider struct {
	request interface{}
	fields  interface{}
}

func (a ardoqBodyProvider) ContentType() string {
	// return jsonContentType
	return "application/json"
}

func (a ardoqBodyProvider) Body() (io.Reader, error) {
	// request := a.request.(ComponentRequest)

	// marshal component
	requestJson, _ := json.Marshal(a.request)

	// create new map as destination for both Unmarshal methods to combine the data
	flatRequest := make(map[string]string)
	json.Unmarshal(requestJson, &flatRequest)

	if len(a.fields.(map[string]interface{})) > 0 {
		// marshal component.Fields
		fieldsJson, _ := json.Marshal(a.fields)
		json.Unmarshal(fieldsJson, &flatRequest)
	}

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(flatRequest)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
