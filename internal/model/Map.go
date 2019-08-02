package model

import (
	"errors"
	"fmt"
)

type Map struct {
	cities map[CityId]*City
}

func NewMap() Map {
	return Map{cities: make(map[CityId]*City)}
}

/*
Adds a new City with given id.
Returns an error if the id already exists.
 */
func (m *Map) AddCity(id CityId) error {
	if _, present := m.cities[id]; present {
		return errors.New(fmt.Sprintf("City %v already exists", id))
	}
	newCity := NewCity(id)
	m.cities[id] = &newCity
	return nil
}

/*
Adds a road from a city to another one, in given direction.
Returns an error if one of the cities doesn't exist, or if the road couldn't be created.
 */
func (m *Map) AddRoad(fromId CityId, toId CityId, direction Direction) error {
	from, present := m.cities[fromId]
	if !present {
		return errors.New(fmt.Sprintf("City %v not found in map", fromId))
	}
	to, present := m.cities[toId]
	if !present {
		return errors.New(fmt.Sprintf("City %v not found in map", toId))
	}
	err := from.AddRoadTo(direction, to)
	return err
}

/*
Destroys a city.
Returns an error if the city doesn't exist.
 */
func (m *Map) DestroyCity(id CityId) error {
	city, present := m.cities[id]
	if !present {
		return errors.New(fmt.Sprintf("City %v doesn't exist", id))
	}
	city.Destroy()
	delete(m.cities, id)
	return nil
}
