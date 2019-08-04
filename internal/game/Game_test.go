package game

import (
	"testing"
	"ratio_technical_test/internal/model"
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func TestNewGame_ShouldFail_OnEmptyMap(t *testing.T) {
	// Given
	world := model.NewMap()
	buffer := new(bytes.Buffer)

	// When
	_, err := NewGame(buffer, &world, 10, 10)

	// Then
	if err == nil {
		t.Error("Should have failed")
	}
}

func TestNewGame_ShouldSpawn10MonstersInOneCity(t *testing.T) {
	// Given
	world := model.NewMap()
	world.AddCity("Lindblum")
	buffer := new(bytes.Buffer)

	// When
	_, err := NewGame(buffer, &world, 10, 10)

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Returned an error: %v", err))
	}
	monstersCount := len(world.GetCities()["Lindblum"].GetMonsters())
	if monstersCount != 10 {
		t.Errorf(fmt.Sprintf("Should have spawned 10 monsters in Lindblum but actually spawned %v", monstersCount))
	}
}

func TestNewGame_ShouldSpawn15MonstersInTwoCities(t *testing.T) {
	// Given
	world := make2Cities1RoadMap()
	buffer := new(bytes.Buffer)

	// When
	_, err := NewGame(buffer, world, 10, 15)

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Returned an error: %v", err))
	}
	monstersCount1 := len(world.GetCities()["Midgard"].GetMonsters())
	monstersCount2 := len(world.GetCities()["Lindblum"].GetMonsters())
	t.Logf("Midgard: %v monsters, Lindblum: %v monsters", monstersCount1, monstersCount2)
	if monstersCount1 + monstersCount2 != 15 {
		t.Errorf(fmt.Sprintf("Should have spawned 15 monsters in total, but actually spawned %v", monstersCount1 + monstersCount2))
	}
}

func TestPlayTurn_With1CityMap_ShouldFinishAfterOneTurn(t *testing.T) {
	// Given
	world := model.NewMap()
	world.AddCity("Midgard")
	buffer := new(bytes.Buffer)
	game, err := NewGame(buffer, &world, 10, 2)

	// When
	isOver := game.PlayTurn()

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Creation of game returned an error: %v", err))
	}
	if !isOver {
		t.Error("Game should have been over")
	}
	if !game.WorldIsDestroyed() {
		t.Error("Wourld should be destroyed")
	}
	if len(game.monsters) != 0 {
		t.Errorf("All monsters should be dead, but %v are remaining", len(game.monsters))
	}
	assertOutput(t, buffer.String(), "Midgard", []string{"monster 1", "monster 2"})
}

func TestPlayTurn_With2Cities1RoadMap_ShouldFinishAfterOneTurn(t *testing.T) {
	// Given
	world := make2Cities1RoadMap()
	buffer := new(bytes.Buffer)
	game, err := NewGame(buffer, world, 10, 100)

	// When
	isOver := game.PlayTurn()

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Creation of game returned an error: %v", err))
	}
	if !isOver {
		t.Error("Game should have been over")
	}
	if game.WorldIsDestroyed() {
		t.Error("Wourld should not be destroyed")
	}
	if _, present := game.world.GetCities()["Midgard"]; !present {
		t.Error("Midgard shouldn't be destroyed")
	}
	if _, present := game.world.GetCities()["Lindblum"]; present {
		t.Error("Lindblum should be destroyed")
	}
	if len(game.monsters) != 0 {
		t.Errorf("All monsters should be dead, but %v are remaining", len(game.monsters))
	}
	monsterNames := make([]string, 0, 100)
	for i := range monsterNames {
		monsterNames[i] = "monster " + strconv.Itoa(i+1)
	}
	assertOutput(t, buffer.String(), "Lindblum", monsterNames)
}

func TestPlayTurn_With0Monsters_ShouldGameOverOnFirstTurn(t *testing.T) {
	// Given
	world := model.NewMap()
	world.AddCity("Lindblum")
	buffer := new(bytes.Buffer)
	game, err := NewGame(buffer, &world, 10, 0)

	// When
	isOver := game.PlayTurn()

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Creation of game returned an error: %v", err))
	}
	if !isOver {
		t.Error("Game should have been over")
	}
	if game.WorldIsDestroyed() {
		t.Error("Wourld should not be destroyed")
	}
}

func TestPlayTurn_With1MaxTurns_ShouldGameOverOnFirstTurn(t *testing.T) {
	// Given
	world := model.NewMap()
	world.AddCity("Lindblum")
	buffer := new(bytes.Buffer)
	game, err := NewGame(buffer, &world, 1, 1)

	// When
	isOver := game.PlayTurn()

	// Then
	if err != nil {
		t.Errorf(fmt.Sprintf("Creation of game returned an error: %v", err))
	}
	if !isOver {
		t.Error("Game should have been over")
	}
	if game.WorldIsDestroyed() {
		t.Error("Wourld should not be destroyed")
	}
}

// A map with two cities and one road
func make2Cities1RoadMap() *model.Map {
	world := model.NewMap()
	world.AddCity("Midgard")
	world.AddCity("Lindblum")
	world.AddRoad("Midgard", "Lindblum", model.North)
	return &world
}

func assertOutput(t *testing.T, output string, city string, monsters []string) {
	if !strings.HasPrefix(output, city + " ") {
		t.Errorf("'%v' should begin with '%v '", output, city)
	}
	for _, monster := range monsters {
		if !strings.Contains(output, monster) {
			t.Errorf("'%v' should contain '%v'", output, monster)
		}
	}
}