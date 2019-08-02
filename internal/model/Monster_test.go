package model

import "testing"

func TestMonster_Move_ShouldReturnNilAndNotMoveIfNoRoads(t *testing.T) {

	// Given
	city := NewCity("Midgard")
	monster := NewMonster("Cthulhu", &city)

	// When
	newPosition := monster.Move()

	// Then
	assertMonsterDidntMove(t, monster, newPosition, &city)
}

func TestMonster_Move_ShouldMoveAwayFromCityWithOneRoad(t *testing.T) {
	// Given
	city := NewCity("Midgard")
	city2 := NewCity("Asgard")
	city.AddRoadTo(West, &city2)
	monster := NewMonster("Cthulhu", &city)

	// When
	newPosition := monster.Move()

	// Then
	assertMonsterMoved(t, monster, newPosition, &city)
	if newPosition != &city2 {
		t.Errorf("Monster %v could only go to %v but went to %v instead", monster.id, city2, newPosition)
	}
}

func TestMonster_Move_ShouldMoveAwayFromCityWithTwoRoads(t *testing.T) {
	// Given
	city := NewCity("Midgard")
	city2 := NewCity("Asgard")
	city3 := NewCity("Kalm")
	city.AddRoadTo(West, &city2)
	city.AddRoadTo(East, &city3)
	monster := NewMonster("Cthulhu", &city)

	// When
	newPosition := monster.Move()

	// Then
	assertMonsterMoved(t, monster, newPosition, &city)
	t.Logf("Moved to %v", newPosition.id)
}

func assertMonsterDidntMove(t *testing.T, m Monster, returnedPosition *City, expectedPosition *City) {
	if returnedPosition != nil {
		t.Errorf("Monster %v move should have returned nil but returned a new position: %v", m.id, returnedPosition)
	}
	if m.position != expectedPosition {
		t.Errorf("Monster %v should have moved to %v, but moved to %v instead", m.id, expectedPosition, m.position)
	}
}

func assertMonsterMoved(t *testing.T, m Monster, returnedPosition *City, previousPosition *City) {
	if returnedPosition != m.position {
		t.Errorf("Monster %v move should have returned the new monster's position %v but returned something else: %v", m.id, m.position, returnedPosition)
	}
	if m.position == previousPosition {
		t.Errorf("Monster %v should have moved away from %v, but stayed there", m.id, previousPosition)
	}
}
