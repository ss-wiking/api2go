package api2go

type Resource[T any] struct {
	Data     T              `json:"data"`
	Included []Entity       `json:"included"`
	Meta     map[string]any `json:"meta"`
}

type ResourceCollection[T any] struct {
	Data     []T      `json:"data"`
	Included []Entity `json:"included"`
	Links    struct {
		Self  string `json:"self"`
		Next  string `json:"next"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		CurrentPage uint   `json:"current_page"`
		From        uint   `json:"from"`
		LastPage    uint   `json:"last_page"`
		PerPage     uint   `json:"per_page"`
		To          uint   `json:"to"`
		Total       uint   `json:"total"`
		Path        string `json:"path"`
		LastId      string `json:"last_id"`
	} `json:"meta"`
}
