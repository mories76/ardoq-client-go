package ardoq

// Model is the struct representation of the model JSON
type Model struct {
	Ardoq struct {
		EntityType             string      `mapstructure:"entity-type"`
		IncomingReferenceCount int         `mapstructure:"incomingReferenceCount"`
		OutgoingReferenceCount int         `mapstructure:"outgoingReferenceCount"`
		Persistent             interface{} `mapstructure:"persistent"`
	} `mapstructure:"ardoq"`
	ArdoqPersistent     []interface{}  `mapstructure:"ardoq-persistent"`
	BlankTemplate       bool           `mapstructure:"blankTemplate"`
	Category            string         `mapstructure:"category"`
	Common              bool           `mapstructure:"common"`
	Created             string         `mapstructure:"created"`
	CreatedBy           string         `mapstructure:"created-by"`
	CreatedByEmail      string         `mapstructure:"createdByEmail"`
	CreatedByName       string         `mapstructure:"createdByName"`
	DefaultViews        []string       `mapstructure:"defaultViews"`
	Description         string         `mapstructure:"description"`
	Flexible            bool           `mapstructure:"flexible"`
	ID                  string         `mapstructure:"_id"`
	LastUpdated2        string         `mapstructure:"last-updated"`
	LastModifiedBy      string         `mapstructure:"last-modified-by"`
	LastModifiedByEmail string         `mapstructure:"lastModifiedByEmail"`
	LastUpdated         string         `mapstructure:"lastupdated"`
	MaxReferenceTypeKey int            `mapstructure:"maxReferenceTypeKey"`
	Name                string         `mapstructure:"name"`
	ReferenceTypes      ReferenceTypes `mapstructure:"referenceTypes"`
	Root                ComponentTypes `mapstructure:"root"`
	StartView           string         `mapstructure:"startView"`
	UseAsTemplate       bool           `mapstructure:"useAsTemplate"`
	Version             int            `mapstructure:"_version"`
	Workspaces          struct {
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
	Origin struct {
		ID      string `mapstructure:"id"`
		Version int    `mapstructure:"_version"`
	}

	Fields map[string]interface{} `mapstructure:",remain"`
}

// func (c Model) GetFields() map[string]interface{} {
// 	if len(c.Fields) > 0 {
// 		return c.Fields
// 	}
// 	return nil
// }

// ReferenceTypes child of Model struct
type ReferenceTypes map[string]struct {
	Name string `mapstructure:"name"`
	ID   string `mapstructure:"id"`
}

// ComponentType child of Model struct
type ComponentType struct {
	Children     ComponentTypes `mapstructure:"children"`
	Color        string         `mapstructure:"color"`
	Icon         string         `mapstructure:"icon"`
	ID           string         `mapstructure:"id"`
	Image        string         `mapstructure:"image"`
	Index        string         `mapstructure:"index"`
	Level        string         `mapstructure:"level"`
	Name         string         `mapstructure:"name"`
	ReturnsValue string         `mapstructure:"returnsValue"`
	Shape        string         `mapstructure:"shape"`
	Standard     string         `mapstructure:"standard"`
}

// ComponentTypes child of Model struct
// TODO: Check if this can be a slice of ComponentType
type ComponentTypes map[string]struct {
	Children     ComponentTypes `mapstructure:"children"`
	Color        string         `mapstructure:"color"`
	Icon         string         `mapstructure:"icon"`
	ID           string         `mapstructure:"id"`
	Image        string         `mapstructure:"image"`
	Index        string         `mapstructure:"index"`
	Level        string         `mapstructure:"level"`
	Name         string         `mapstructure:"name"`
	ReturnsValue string         `mapstructure:"returnsValue"`
	Shape        string         `mapstructure:"shape"`
	Standard     string         `mapstructure:"standard"`
}

// GetComponentTypes this function is not being used yet.
// TODO: figure out the correct terraform provider schema
func (m Model) GetComponentTypes() map[string]ComponentType {
	result := make(map[string]ComponentType)

	for _, v := range m.Root {
		result[v.Name] = v
	}
	return result
}

// FIX this doesn't seem very efficient
func componentTypeGetChildren(root ComponentTypes) map[string]string {
	ComponentTypes := make(map[string]string)
	for _, r := range root {
		ComponentTypes[r.Name] = r.ID
		if len(r.Children) > 0 {
			for k, v := range componentTypeGetChildren(r.Children) {
				//
				ComponentTypes[k] = v
			}
		}
	}
	return ComponentTypes
}

// GetComponentTypeID returns a flattend map[string]string of name and ID for all the componentTypes
func (m Model) GetComponentTypeID() map[string]string {
	// result := make(map[string]string)

	// TODO make this function traverse the root until v.Name is found
	// for example model 19ea590239001b064dbc878d and component p1575623714660
	// for _, v := range m.Root {
	// 	result[v.Name] = v.ID
	// }

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