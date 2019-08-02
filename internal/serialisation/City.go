package serialisation

import "ratio_technical_test/internal/model"

type City struct {
	id    model.CityId
	roads []Road
}

func NewCity(id model.CityId, roads []Road) City {
	return City{id: id, roads: roads}
}

type Road struct {
	direction   model.Direction
	destination model.CityId
}
