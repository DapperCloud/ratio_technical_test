package model

import "math/rand"

type MonsterId string

type Monster struct {
	id       MonsterId
	position *City
}

func NewMonster(id MonsterId, position *City) Monster {
	return Monster{id: id, position: position}
}

/*
The monster makes a move, and returns its new position, or nil if it couldn't move
 */
func (m *Monster) Move() *City {
	roads := m.position.GetRoads()
	roadsNumber := len(roads)
	if roadsNumber == 0 {
		return nil
	}
	roadIndex := 0
	if roadsNumber > 1 {
		roadIndex = rand.Intn(len(roads) - 1)
	}
	newPosition, err := m.position.GetCityInDirection(roads[roadIndex].Direction)
	if err != nil {
		panic(err)
	}
	m.position = newPosition
	return m.position
}

/*
Returns the monster id
 */
func (m Monster) GetId() MonsterId {
	return m.id
}
