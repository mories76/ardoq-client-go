package ardoq

// Model is the struct representation of the model JSON
type Model struct {
	Ardoq struct {
		EntityType             string      `mapstructure:"entity-type"`
		IncomingReferenceCount int         `mapstructure:"incomingReferenceCount"`
		OutgoingReferenceCount int         `mapstructure:"outgoingReferenceCount"`
		Persistent             interface{} `mapstructure:"persistent"`
	} `mapstructure:"ardoq"`
	ArdoqPersistent     []interface{} `mapstructure:"ardoq-persistent"`
	BlankTemplate       bool          `mapstructure:"blankTemplate"`
	Category            string        `mapstructure:"category"`
	Common              bool          `mapstructure:"common"`
	Created             string        `mapstructure:"created"`
	CreatedBy           string        `mapstructure:"createdBy"`
	CreatedBy2          string        `mapstructure:"created-by"`
	CreatedByEmail      string        `mapstructure:"createdByEmail"`
	CreatedByName       string        `mapstructure:"createdByName"`
	DefaultViews        []string      `mapstructure:"defaultViews"`
	Description         string        `mapstructure:"description"`
	Flexible            bool          `mapstructure:"flexible"`
	ID                  string        `mapstructure:"_id"`
	LastModifiedBy      string        `mapstructure:"lastModifiedBy"`
	LastModifiedBy2     string        `mapstructure:"last-modified-by"`
	LastModifiedByEmail string        `mapstructure:"lastModifiedByEmail"`
	LastModifiedByName  string        `mapstructure:"lastModifiedByName"`
	LastUpdated         string        `mapstructure:"lastupdated"`
	LastUpdated2        string        `mapstructure:"last-updated"`
	MaxReferenceTypeKey int           `mapstructure:"maxReferenceTypeKey"`
	Name                string        `mapstructure:"name"`
	Origin              struct {
		ID      string `mapstructure:"id"`
		Version int    `mapstructure:"_version"`
	}
	ReferenceTypes MdoelReferenceTypes `mapstructure:"referenceTypes"`
	Root           ModelComponentTypes `mapstructure:"root"`
	StartView      string              `mapstructure:"startView"`
	UseAsTemplate  bool                `mapstructure:"useAsTemplate"`
	Version        int                 `mapstructure:"_version"`
	Workspaces     struct {
		Restricted int `mapstructure:"restricted"`
		UsedBy     []struct {
			ID                  string `mapstructure:"_id"`
			Name                string `mapstructure:"name"`
			CreatedByName       string `mapstructure:"createdByName"`
			CreatedByEmail      string `mapstructure:"createdByEmail"`
			LastModifiedByName  string `mapstructure:"lastModifiedByName"`
			LastModifiedByEmail string `mapstructure:"lastModifiedByEmail"`
			Ardoq               struct {
				EntityType string `mapstructure:"entity-type"`
			} `mapstructure:"ardoq"`
		} `mapstructure:"used-by"`
	} `mapstructure:"workspaces"`
	// Fields, is the safetynet for when mapping the API response to the struct.
	// The goal is to have all fields documented in the API to have a known field in the struct.
	// For models, Fields should be null
	Fields map[string]interface{} `mapstructure:",remain"`
}

// MdoelReferenceTypes child of Model struct
type MdoelReferenceTypes map[string]struct {
	Name string `mapstructure:"name"`
	ID   string `mapstructure:"id"`
}

// ModelComponentTypes child of Model struct
type ModelComponentTypes map[string]struct {
	Children     ModelComponentTypes `mapstructure:"children"`
	Color        string              `mapstructure:"color"`
	Icon         string              `mapstructure:"icon"`
	ID           string              `mapstructure:"id"`
	Image        string              `mapstructure:"image"`
	Index        string              `mapstructure:"index"`
	Level        string              `mapstructure:"level"`
	Name         string              `mapstructure:"name"`
	ReturnsValue string              `mapstructure:"returnsValue"`
	Shape        string              `mapstructure:"shape"`
	Standard     string              `mapstructure:"standard"`
}

// componentTypeGetChildren walks the tree of ModelComponentTypes starting from the Model.Root
func componentTypeGetChildren(componentTypes ModelComponentTypes) map[string]string {
	// create result type
	result := make(map[string]string)

	// for each ComponentType (ct) in these componentTypes
	for _, ct := range componentTypes {
		// add Name and ID to result
		result[ct.Name] = ct.ID

		// check if this branch has children
		if len(ct.Children) > 0 {
			// get rescursive children and loop
			for k, v := range componentTypeGetChildren(ct.Children) {
				// add k,v => Name and ID to the result
				result[k] = v
			}
		}
	}

	return result
}

// GetComponentTypeID returns a flattend map[string]string of name and ID for all the componentTypes
// the respone from the Ardoq model reflects the structure of the metamodel
// so there is a tree of N levels deep of "ModelComponentTypes"
// this function returns something like
// returnvalue["componentName"] = "componentID"
func (m Model) GetComponentTypeID() map[string]string {
	return componentTypeGetChildren(m.Root)
}

// GetReferenceTypes rewrites the input "map[string]struct" with attributes Name and ID
// to a more usuable structure of map[string]string
// this way you can use the function GetReferenceTypes()["name of reference"] and in return get its ID
func (m Model) GetReferenceTypes() map[string]string {
	result := make(map[string]string)
	for _, referenceType := range m.ReferenceTypes {
		result[referenceType.Name] = referenceType.ID
	}

	return result
}
