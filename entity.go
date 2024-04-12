package api2go

type EntityId struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Entity struct {
	EntityId
	Attributes    map[string]any            `json:"attributes"`
	Relationships map[string]map[string]any `json:"relationships"`
}
