package model

import "math/rand"

type MonsterId string

type Monster struct {
	id       MonsterId
	position *City
}

func NewMonster(id MonsterId, position *City) Monster {
	newMonster := Monster{id: id, position: position}
	position.AddMonster(&newMonster)
	return newMonster
}

/*
The monster makes a move, and returns its new position
 */
func (m *Monster) Move() *City {
	roads := m.position.GetRoads()
	roadsNumber := len(roads)
	if roadsNumber == 0 {
		return m.position
	}
	roadIndex := 0
	if roadsNumber > 1 {
		roadIndex = rand.Intn(len(roads))
	}
	newPosition, err := m.position.GetCityInDirection(roads[roadIndex].Direction)
	if err != nil {
		panic(err)
	}
	m.position.RemoveMonster(m)
	m.position = newPosition
	m.position.AddMonster(m)
	return m.position
}

/*
Returns the monster id
 */
func (m Monster) GetId() MonsterId {
	return m.id
}

/*
Returns the monster position
 */
func (m Monster) GetPosition() *City {
	return m.position
}
