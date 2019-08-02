package model

import (
	"errors"
	"fmt"
)

type CityId string

type City struct {
	id            CityId
	outgoingRoads map[Direction]*City
	incomingRoads map[*City]Direction
}

func NewCity(id CityId) City {
	return City{id: id, outgoingRoads: make(map[Direction]*City), incomingRoads: make(map[*City]Direction)}
}

/*
Adds a road to another city in given direction.
You cannot add two roads to the same city, and you cannot add two roads in the same direction.
 */
func (c *City) AddRoadTo(direction Direction, city *City) error {
	if c == city {
		return errors.New(fmt.Sprintf("City %v cannot add a new road to itself", c.id))
	}
	if _, present := c.outgoingRoads[direction]; present {
		return errors.New(fmt.Sprintf("City %v already has a road in direction %v, cannot add another one",
			c.id, direction))
	}
	for existingDirection, existingNeighbour := range c.outgoingRoads {
		if existingDirection != direction && existingNeighbour == city {
			return errors.New(fmt.Sprintf("City %v already has a road to %v in direction %v, so cannot add a new road to id",
				c.id, existingNeighbour, existingDirection))
		}
	}
	c.outgoingRoads[direction] = city
	city.incomingRoads[c] = direction
	return nil
}

type Road struct {
	Direction   Direction
	Destination CityId
}

/*
Returns the city id
 */
func (c City) GetId() CityId {
	return c.id
}

/*
Returns the roads going out of this city.
 */
func (c City) GetRoads() []Road {
	roads := make([]Road, len(c.outgoingRoads))
	i := 0
	for direction, destination := range c.outgoingRoads {
		roads[i] = Road{Direction: direction, Destination: destination.id}
		i++
	}
	return roads
}

/*
Returns the City in given direction, or an error if no city in that direction
 */
func (c City) GetCityInDirection(direction Direction) (*City, error) {
 	city, present := c.outgoingRoads[direction]
 	if !present {
 		return nil, errors.New(fmt.Sprintf("City %v has no road going in direction %v", c.id, direction))
	}
	return city, nil
}

/*
Destroys the city.
All roads to and from this city will be destroyed.
 */
func (c *City) Destroy() {
	for _, neighbour := range c.outgoingRoads {
		delete(neighbour.incomingRoads, c)
	}
	for city, direction := range c.incomingRoads {
		delete(city.outgoingRoads, direction)
	}
}
