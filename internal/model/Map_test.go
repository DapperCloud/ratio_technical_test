package model

import "testing"

func TestMap_AddCity_ShouldAddOneCity(t *testing.T) {
	// Given
	m := NewMap()

	// When
	err := m.AddCity("Midgard")

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	if _, present := m.cities["Midgard"]; !present {
		t.Error("Should have added the city")
	}
}

func TestMap_AddCity_ShouldAddTwoCities(t *testing.T) {
	// Given
	m := NewMap()

	// When
	m.AddCity("Balamb")
	err := m.AddCity("Esther")

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	if _, present := m.cities["Balamb"]; !present {
		t.Error("Should have added the first city")
	}
	if _, present := m.cities["Esther"]; !present {
		t.Error("Should have added the second city")
	}
}

func TestMap_AddCity_ShouldFail_OnCityAlreadyExists(t *testing.T) {
	// Given
	m := NewMap()
	m.AddCity("Midgard")

	// When
	err := m.AddCity("Midgard")

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMap_AddRoad_ShouldAddRoadBetweenExistingCities(t *testing.T) {
	// Given
	m := NewMap()
	m.AddCity("Besaid")
	m.AddCity("Zanarkand")

	// When
	err := m.AddRoad("Besaid", "Zanarkand", East)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	assertRoadExists(t, m, "Besaid", "Zanarkand", East)
}

func TestMap_AddRoad_ShouldFail_OnRoadAlreadyExists(t *testing.T) {
	// Given
	m := NewMap()
	m.AddCity("Besaid")
	m.AddCity("Zanarkand")

	// When
	m.AddRoad("Besaid", "Zanarkand", East)
	err := m.AddRoad("Besaid", "Zanarkand", East)

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
	assertRoadExists(t, m, "Besaid", "Zanarkand", East)
}

func TestMap_AddRoad_ShouldFail_OnNonExistingCityTo(t *testing.T) {
	// Given
	m := NewMap()
	m.AddCity("Besaid")

	// When
	err := m.AddRoad("Besaid", "Zanarkand", East)

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMap_AddRoad_ShouldFail_OnNonExistingCityFrom(t *testing.T) {
	// Given
	m := NewMap()
	m.AddCity("Zanarkand")

	// When
	err := m.AddRoad("Besaid", "Zanarkand", East)

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func assertRoadExists(t *testing.T, m Map, fromId CityId, toId CityId, direction Direction) {
	from, present := m.cities[fromId]
	if !present {
		t.Errorf("Map should contain city %v", fromId)
	}
	to, present := m.cities[toId]
	if !present {
		t.Errorf("Map should contain city %v", toId)
	}
	road, present := from.outgoingRoads[direction]
	if !present {
		t.Errorf("City %v should have a road in direction %v", fromId, direction)
	}
	if road != to {
		t.Errorf("City %v should have a road in direction %v to city %v, but instances don't match with actual city %v",
			fromId, direction, toId, road.id)
	}
}
