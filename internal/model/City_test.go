package model

import "testing"

func TestCity_AddNeighbour_ShouldFail_OnRoadToSelf(t *testing.T) {
	// Given
	c := NewCity("Midgard")

	// When
	err := c.AddRoadTo(East, &c)

	// Then
	if err == nil {
		t.Error("Should have returned an error")
	}
}

func TestCity_AddNeighbour_ShouldFail_OnSameNeighbourInTwoDirections(t *testing.T) {
	// Given
	c := NewCity("Midgard")
	c2 := NewCity("Kalm")
	c.AddRoadTo(South, &c2)

	// When
	err := c.AddRoadTo(East, &c2)

	// Then
	if err == nil {
		t.Error("Should have returned an error")
	}
}

func TestCity_AddNeighbour_ShouldFail_OnRoadAlreadyExists(t *testing.T) {
	// Given
	c := NewCity("Tristram")
	c2 := NewCity("Kurath")
	c3 := NewCity("Lut Gholein")
	c.AddRoadTo(West, &c2)

	// When
	err := c.AddRoadTo(West, &c3)

	// Then
	if err == nil {
		t.Error("Should have returned an error")
	}
}

func TestCity_AddNeighbour_ShouldAddRoad(t *testing.T) {
	// Given
	c := NewCity("Ankh-Morpork")
	c2 := NewCity("Chirm")

	// When
	err := c.AddRoadTo(East, &c2)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	if c.outgoingRoads[East] != &c2 {
		t.Error("Should have added new outgoing road")
	}
	if c2.incomingRoads[&c] != East {
		t.Error("Should have added new incoming road")
	}
}

func TestCity_Destroy_ShouldDestroyAllRoads(t *testing.T) {
	// Given
	c := NewCity("Besaid")
	c2 := NewCity("Luca")
	c3 := NewCity("Zanarkand")
	c.AddRoadTo(East, &c2)
	c2.AddRoadTo(North, &c)
	c.AddRoadTo(West, &c3)
	c3.AddRoadTo(South, &c)
	c2.AddRoadTo(South, &c3)
	c3.AddRoadTo(East, &c2)

	// When
	c.Destroy()

	// Then
	if len(c2.outgoingRoads) != 1 {
		t.Errorf("%v should have 1 outgoing road but instead had %v", c2.id, len(c2.outgoingRoads))
	}
	if len(c3.outgoingRoads) != 1 {
		t.Errorf("%v should have 1 outgoing road but instead had %v", c3.id, len(c3.outgoingRoads))
	}
	if _, present := c2.outgoingRoads[North]; present {
		t.Errorf("%v road to %v should have been destroyed", c2.id, c.id)
	}
	if _, present := c3.outgoingRoads[South]; present {
		t.Errorf("%v road to %v should have been destroyed", c3.id, c.id)
	}
	if _, present := c2.incomingRoads[&c]; present {
		t.Errorf("%v incoming road from %v should have been destroyed", c2.id, c.id)
	}
	if _, present := c3.incomingRoads[&c]; present {
		t.Errorf("%v incoming road from %v should have been destroyed", c3.id, c.id)
	}
}
