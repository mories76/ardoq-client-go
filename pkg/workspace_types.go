package ardoq

// Workspace is the struct representation of the workspace JSON
type Workspace struct {
	Ardoq struct {
		EntityType string `mapstructure:"entity-type"`
	} `mapstructure:"ardoq"`
	ArdoqPersistent     []interface{}          `mapstructure:"ardoq-persistent"`
	CompCounter         int                    `mapstructure:"comp-counter"`
	ComponentModel      string                 `mapstructure:"componentModel"`
	ComponentTemplate   string                 `mapstructure:"componentTemplate"`
	Created             string                 `mapstructure:"created"`
	CreatedBy           string                 `mapstructure:"created-by"`
	CreatedByEmail      string                 `mapstructure:"createdByEmail"`
	CreatedByName       string                 `mapstructure:"createdByName"`
	Description         string                 `mapstructure:"description"`
	DefaultPerspective  string                 `mapstructure:"defaultPerspective"`
	Fields              map[string]interface{} `mapstructure:",remain"`
	ID                  string                 `mapstructure:"_id"`
	LastUpdated2        string                 `mapstructure:"last-updated"`
	LastModifiedBy      string                 `mapstructure:"last-modified-by"`
	LastModifiedByEmail string                 `mapstructure:"lastModifiedByEmail"`
	LastModifiedByName  string                 `mapstructure:"lastModifiedByName"`
	LastUpdated         string                 `mapstructure:"lastupdated"`
	LinkedWorkspaces    struct {
		Linked     []string `mapstructure:"linked"`
		BackLinked []string `mapstructure:"backlinked"`
	} `mapstructure:"linked-workspaces"`
	Name   string `mapstructure:"name"`
	Origin struct {
		EntityType string `mapstructure:"entity-type"`
	} `mapstructure:"origin"`
	StartView    string   `mapstructure:"startView"`
	Type         string   `mapstructure:"type"`
	Version      int      `mapstructure:"_version"`
	Views        []string `mapstructure:"views"`
	WorkspaceKey string   `mapstructure:"workspace-key"`
}

// type Workspaces []Workspace

// WorkspaceSearchQuery is the query used to search for workspaces
type WorkspaceSearchQuery struct {
	Name string `url:"name,omitempty"`
}
