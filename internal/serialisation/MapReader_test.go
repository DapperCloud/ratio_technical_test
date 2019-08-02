package serialisation

import (
	"testing"
	"strings"
	"ratio_technical_test/internal/model"
)

func TestMapReader_GetMap_ShouldGetMapWithTwoCities(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard east=Kalm\nKalm south=Midgard")
	mapReader := NewMapReader(reader)

	// When
	newMap, err := mapReader.GetMap()

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	assertCityHasRoads(t, newMap, "Midgard", model.Road{Direction: model.East, Destination: "Kalm"})
	assertCityHasRoads(t, newMap, "Kalm", model.Road{Direction: model.South, Destination: "Midgard"})
}

func TestMapReader_GetMap_ShouldFail_OnEmptyLine(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard east=Kalm\nKalm south=Midgard\n\n")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMapReader_GetMap_ShouldFail_OnNoRoads(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard east=Kalm\nKalm")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMapReader_GetMap_ShouldFail_OnInvalidRoad(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard foo\nKalm south=Midgard")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMapReader_GetMap_ShouldFail_OnInvalidRoad2(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard east=Kalm=foo\nKalm south=Midgard")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMapReader_GetMap_ShouldFail_OnInvalidDirection(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard foo=Kalm\nKalm south=Midgard")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestMapReader_GetMap_ShouldFail_OnUnknownCity(t *testing.T) {
	// Given
	reader := strings.NewReader("Midgard east=Kalm south=Neverland\nKalm south=Midgard")
	mapReader := NewMapReader(reader)

	// When
	_, err := mapReader.GetMap()

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func assertCityHasRoads(t *testing.T, newMap model.Map, cityId model.CityId, expectedRoads ... model.Road) {
	actualRoads, err := newMap.GetRoadsFrom(cityId)
	if err != nil {
		t.Errorf("Error while checking Midgard's roads: %v", err)
	}
	if len(actualRoads) != len(expectedRoads) {
		t.Errorf("%v should have %v roads, but actually has %v", cityId, len(expectedRoads), len(actualRoads))
	}
	for _, road := range expectedRoads {
		if !contains(actualRoads, road) {
			t.Errorf("%v should have %v road, but does not. Actual roads: %v", cityId, road, actualRoads)
		}
	}
}

func contains(slice []model.Road, element model.Road) bool {
    for _, a := range slice {
        if a == element {
            return true
        }
    }
    return false
}
