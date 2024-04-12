package api2go

import (
	"encoding/json"
	"errors"
)

type HasOne[T any] struct {
	EntityId `json:"data"`
}

type HasMany[T any] struct {
	Items []EntityId `json:"data"`
}

func (h HasOne[T]) GetFrom(included []Entity) (T, error) {
	var relation T

	// TODO: refactor this piece of shit
	for _, entity := range included {
		if h.EntityId == entity.EntityId {
			data, err := json.Marshal(entity)
			if err != nil {
				return relation, err
			}

			err = json.Unmarshal(data, &relation)
			return relation, err
		}
	}

	return relation, errors.New("entity not found")
}

func (h HasMany[T]) GetFrom(included []Entity) ([]T, error) {
	var relations []T
	var entities []Entity

	// TODO: refactor this piece of shit
	for _, entity := range included {
		for _, relation := range h.Items {
			if relation == entity.EntityId {
				entities = append(entities, entity)
			}
		}
	}

	data, err := json.Marshal(entities)
	if err != nil {
		return relations, err
	}

	err = json.Unmarshal(data, &relations)

	return relations, err
}
