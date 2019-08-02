package serialisation

import (
	"testing"
	"ratio_technical_test/internal/model"
	"bytes"
	"strings"
)

func TestMapWriter_WriteMap_ShouldWriteEmptyMap(t *testing.T) {
	// Given
	modelMap := model.NewMap()
	buffer := new(bytes.Buffer)
	mapWriter := NewMapWriter(buffer)

	// When
	err := mapWriter.WriteMap(modelMap)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	writtenMap := buffer.String()
	if len(writtenMap) > 0 {
		t.Errorf("Should have written nothing, but wrote '%v'", writtenMap)
	}
}

func TestMapWriter_WriteMap_ShouldWriteMapWithOneCityAndNoRoads(t *testing.T) {
	// Given
	modelMap := model.NewMap()
	modelMap.AddCity("Midgard")
	buffer := new(bytes.Buffer)
	mapWriter := NewMapWriter(buffer)

	// When
	err := mapWriter.WriteMap(modelMap)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	writtenMap := buffer.String()
	assertHasNumberOfLines(t, writtenMap, 1)
	assertContainsLine(t, writtenMap, "Midgard")
}

func TestMapWriter_WriteMap_ShouldWriteMapWithTwoCitiesAndOneRoad(t *testing.T) {
	// Given
	modelMap := model.NewMap()
	modelMap.AddCity("Midgard")
	modelMap.AddCity("Utai")
	modelMap.AddRoad("Midgard", "Utai", model.East)
	buffer := new(bytes.Buffer)
	mapWriter := NewMapWriter(buffer)

	// When
	err := mapWriter.WriteMap(modelMap)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	writtenMap := buffer.String()
	assertHasNumberOfLines(t, writtenMap, 2)
	assertContainsLine(t, writtenMap, "Midgard", "east=Utai")
	assertContainsLine(t, writtenMap, "Utai")
}

func TestMapWriter_WriteMap_ShouldWriteMapWithThreeCitiesAndManyRoads(t *testing.T) {
	// Given
	modelMap := model.NewMap()
	modelMap.AddCity("Midgard")
	modelMap.AddCity("Utai")
	modelMap.AddCity("Nibelheim")
	modelMap.AddRoad("Midgard", "Utai", model.East)
	modelMap.AddRoad("Midgard", "Nibelheim", model.West)
	modelMap.AddRoad("Utai", "Midgard", model.South)
	buffer := new(bytes.Buffer)
	mapWriter := NewMapWriter(buffer)

	// When
	err := mapWriter.WriteMap(modelMap)

	// Then
	if err != nil {
		t.Errorf("Returned an error: %v", err)
	}
	writtenMap := buffer.String()
	assertHasNumberOfLines(t, writtenMap, 3)
	assertContainsLine(t, writtenMap, "Midgard", "east=Utai", "west=Nibelheim")
	assertContainsLine(t, writtenMap, "Utai", "south=Midgard")
	assertContainsLine(t, writtenMap, "Nibelheim")
}

func assertHasNumberOfLines(t *testing.T, output string, expected int) {
	actual := strings.Count(output, "\n")
	if actual != expected {
		t.Errorf("Should have written %v lines but wrote %v. Output: %v", expected, actual, output)
	}
}

func assertContainsLine(t *testing.T, output string, city string, roadTokens... string) {
	lines := strings.Split(output, "\n")
	for _, currentLine := range lines {
		tokens := strings.Split(currentLine, " ")
		if len(tokens) == 1 + len(roadTokens) && tokens[0] == city && containsAllTokens(currentLine, tokens[1:]) {
			return
		}
	}

	t.Errorf("Shoulde have written line for city %v with roads '%v' but didn't. Output: %v", city, strings.Join(roadTokens, " "), output)
}

func containsAllTokens(str string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(str, token) {
			return false
		}
	}
	return true
}
