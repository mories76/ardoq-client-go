package ardoq

// Field is the struct representation of the model JSON
type Field struct {
	Ardoq struct {
		EntityType string `mapstructure:"entity-type"`
	} `mapstructure:"ardoq"`
	ComponentType  []string `mapstructure:"componentType"`
	Created        string   `mapstructure:"created"`
	CreatedBy      string   `mapstructure:"created-by"`
	CreatedByEmail string   `mapstructure:"createdByEmail"`
	CreatedByName  string   `mapstructure:"createdByName"`
	DateTimeFields []struct {
		DefaultValue interface{} `mapstructure:"defaultValue"`
		Label        string      `mapstructure:"label"`
		Name         string      `mapstructure:"name"`
		Type         string      `mapstructure:"type"`
	} `mapstructure:"dateTimeFields"`
	DefaultValue        string  `mapstructure:"defaultValue"`
	Description         string  `mapstructure:"description"`
	Global              bool    `mapstructure:"global"`
	GlobalRef           bool    `mapstructure:"globalref"`
	ID                  string  `mapstructure:"_id"`
	Label               string  `mapstructure:"label"`
	LastModifiedBy      string  `mapstructure:"last-modified-by"`
	LastModifiedByEmail string  `mapstructure:"lastModifiedByEmail"`
	LastModifiedByName  string  `mapstructure:"lastModifiedByName"`
	LastUpdated         string  `mapstructure:"lastupdated"`
	LastUpdated2        string  `mapstructure:"last-updated"`
	Model               string  `mapstructure:"model"`
	Name                string  `mapstructure:"name"`
	Order               float64 `mapstructure:"_order"`
	Origin              struct {
		ID      string `mapstructure:"id"`
		Version int    `mapstructure:"_version"`
	}
	ReferenceType []string `mapstructure:"referenceType"`
	Type          string   `mapstructure:"type"`
	Version       int      `mapstructure:"_version"`

	Fields map[string]interface{} `mapstructure:",remain"`
}
