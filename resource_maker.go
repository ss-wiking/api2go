package api2go

import (
	"encoding/json"
)

// MakeFromJson is a function to parse single resource to a struct
func MakeFromJson[T any](data []byte) (Resource[T], error) {
	var r Resource[T]
	err := json.Unmarshal(data, &r)

	return r, err
}

// MakeCollectionFromJson is a function to parse resource collection to a struct
func MakeCollectionFromJson[T any](data []byte) (ResourceCollection[T], error) {
	var c ResourceCollection[T]
	err := json.Unmarshal(data, &c)

	return c, err
}
