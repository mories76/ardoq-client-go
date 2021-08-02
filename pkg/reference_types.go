package ardoq

// Reference is the struct representation of the reference JSON
type Reference struct {
	Ardoq struct {
		EntityType             string      `mapstructure:"entity-type"`
		IncomingReferenceCount int         `mapstructure:"incomingReferenceCount"`
		OutgoingReferenceCount int         `mapstructure:"outgoingReferenceCount"`
		Persistent             interface{} `mapstructure:"persistent"`
	} `mapstructure:"ardoq"`
	Created             string                 `mapstructure:"created"`
	CreatedBy           string                 `mapstructure:"created-by"`
	CreatedByEmail      string                 `mapstructure:"createdByEmail"`
	CreatedByName       string                 `mapstructure:"createdByName"`
	DisplayText         string                 `mapstructure:"displayText"`
	Description         string                 `mapstructure:"description"`
	ID                  string                 `mapstructure:"_id"`
	LastUpdated2        string                 `mapstructure:"last-updated"`
	LastModifiedBy      string                 `mapstructure:"last-modified-by"`
	LastModifiedByName  string                 `mapstructure:"lastModifiedByName"`
	LastModifiedByEmail string                 `mapstructure:"lastModifiedByEmail"`
	LastUpdated         string                 `mapstructure:"lastupdated"`
	Order               int                    `mapstructure:"order"`
	RootWorkspace       string                 `mapstructure:"rootWorkspace"`
	Source              string                 `mapstructure:"source"`
	Target              string                 `mapstructure:"target"`
	TargetWorkspace     string                 `mapstructure:"targetWorkspace"`
	Type                int                    `mapstructure:"type"`
	Model               string                 `mapstructure:"model"`
	Version             int                    `mapstructure:"_version"`
	Fields              map[string]interface{} `mapstructure:",remain"`
}

// ReferenceRequest is the payload for creating and updating a reference
// Fields map has json tag "-" so that it doesn't get marshalled into JSON
// the fields are being handled by the ardoqBodyProvider
// URL: PATCH/POST /api/reference
type ReferenceRequest struct {
	Description     interface{}            `json:"description,omitempty"`
	DisplayText     interface{}            `json:"displayText,omitempty"`
	RootWorkspace   interface{}            `json:"rootWorkspace,omitempty"`
	Source          interface{}            `json:"source,omitempty"`
	Target          interface{}            `json:"target,omitempty"`
	TargetWorkspace interface{}            `json:"targetWorkspace,omitempty"`
	Type            interface{}            `json:"type,omitempty"`
	Fields          map[string]interface{} `json:"-"`
}
