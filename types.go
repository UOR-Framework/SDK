package sdk

type Element struct {
	// LocatorType is the address of type information for the location.
	LocatorType string `json:"locatorType,omitempty"`
	// ResourceType is the address of type information for the resource.
	ResourceType string `json:"resourceType,omitempty"`
	// Location is the location of the resource.
	Location map[string]interface{} `json:"location,omitempty"`
	// Resource is the resource.
	Resource map[string]interface{} `json:"resource,omitempty"`
}

// Statement provides `application/vnd.uor.statement.v1+json` mediatype structure when marshalled to JSON.
type Statement struct {
	SchemaVersion int `json:"schemaVersion,omitempty"`

	// MediaType specifies the type of this document data structure e.g. `application/vnd.uor.statement.v1json`.
	ResType string `json:"resType,omitempty"`

	// Subject is the subject of the statement.
	Subject *Element `json:"subject,omitempty"`

	// Predicate is the predicate of the statement.
	Predicate *Element `json:"predicate,omitempty"`

	// Object is the object of the statement.
	Object *Element `json:"object,omitempty"`
}
