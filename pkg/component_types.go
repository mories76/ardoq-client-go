package ardoq

// Component is the struct representation of the component JSON
type Component struct {
	Ardoq struct {
		EntityType             string      `mapstructure:"entity-type"`
		Persistent             interface{} `mapstructure:"persistent"`
		IncomingReferenceCount int         `mapstructure:"incomingReferenceCount"`
		OutgoingReferenceCount int         `mapstructure:"outgoingReferenceCount"`
	} `mapstructure:"ardoq"`
	Children            []string    `mapstructure:"children"`
	ComponentKey        string      `mapstructure:"component-key"`
	Created             string      `mapstructure:"created"`
	CreatedBy           string      `mapstructure:"created-by"`
	CreatedByEmail      string      `mapstructure:"createdByEmail"`
	CreatedByName       string      `mapstructure:"createdByName"`
	Description         string      `mapstructure:"description"`
	ID                  string      `mapstructure:"_id"`
	LastUpdated2        string      `mapstructure:"last-updated"`
	LastModifiedBy      string      `mapstructure:"last-modified-by"`
	LastModifiedByEmail string      `mapstructure:"lastModifiedByEmail"`
	LastModifiedByName  string      `mapstructure:"lastModifiedByName"`
	LastUpdated         string      `mapstructure:"lastupdated"`
	Model               string      `mapstructure:"model"`
	Name                string      `mapstructure:"name"`
	Order               float64     `mapstructure:"_order"`
	Parent              interface{} `mapstructure:"parent"`
	RootWorkspace       string      `mapstructure:"rootWorkspace"`
	Type                string      `mapstructure:"type"`
	TypeID              string      `mapstructure:"typeId"`
	Version             int         `mapstructure:"_version"`

	// Color              interface{} `mapstructure:"color,omitempty"`
	// LifecyclePhase     interface{} `mapstructure:"lifecycle_phase,omitempty"`
	// LiveStartDate      interface{} `mapstructure:"live_start_date,omitempty"`
	// HostingType        string      `mapstructure:"hosting_type,omitempty"`
	// LiveEndDate         interface{} `mapstructure:"live_end_date,omitempty"

	// Icon  interface{} `mapstructure:"icon,omitempty"`
	// Shape interface{} `mapstructure:"shape,omitempty"`
	// Image interface{} `mapstructure:"image,omitempty"`

	Fields map[string]interface{} `mapstructure:",remain"`
}

// GetFields returns fields if there are any, and removes empty fields
// TODO check if removeNull is nececary for other types like models, references or workspaces
func (c Component) GetFields() map[string]interface{} {
	if len(c.Fields) > 0 {
		removeNulls(c.Fields)
		return c.Fields
	}
	return nil
}

// ComponentSearchQuery defines the query parameters for a component search
type ComponentSearchQuery struct {
	Workspace string `url:"workspace"`
	// Field     string `url:"field,omitempty"`
	Name  string `url:"name,omitempty"`
	Value string `url:"value,omitempty"`
	// Org       string `url:"org,omitempty"`
}

// ComponentRequest is the payload for creating and updating a component
// Fields map has json tag "-" so that it doesn't get marshalled into JSON
// the fields are being handled by the ardoqBodyProvider
// URL: PATCH/POST /api/component
type ComponentRequest struct {
	RootWorkspace string                 `json:"rootWorkspace,omitempty"`
	Name          interface{}            `json:"name,omitempty"`
	Description   interface{}            `json:"description,omitempty"`
	Parent        interface{}            `json:"parent,omitempty"` //TODO: interface{} was string, but string default to "" if empty, interface return null in json is empty, other option might be *string
	TypeID        interface{}            `json:"typeId,omitempty"`
	Fields        map[string]interface{} `json:"-"`
}
