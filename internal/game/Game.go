package game

import (
	"ratio_technical_test/internal/model"
	"errors"
	"math/rand"
	"io"
	"strings"
	"strconv"
	"fmt"
)

type Game struct {
	world       *model.Map
	monsters    map[model.MonsterId]*model.Monster
	currentTurn uint
	maxTurns    uint

	writer io.Writer
}

/*
Create a new Game with a writer to output the game progression, a given world map, a max number of turns, and a number of monsters to spawn.
Returns an error if the Game couldn't be created (because the given map is empty).
 */
func NewGame(writer io.Writer, world *model.Map, maxTurns uint, monstersCount uint) (Game, error) {
	cityIds := getCityIds(world)
	if len(cityIds) == 0 {
		return Game{}, errors.New("the map has no cities")
	}
	monsters := make(map[model.MonsterId]*model.Monster)
	for i := uint(0); i < monstersCount; i++ {
		monsterId := model.MonsterId("monster " + strconv.Itoa(int(i+1)))
		cityIndex := 0
		if len(cityIds) > 1 {
			cityIndex = rand.Intn(len(cityIds))
		}
		cityId := cityIds[cityIndex]
		newMonster := model.NewMonster(monsterId, world.GetCities()[cityId])
		monsters[newMonster.GetId()] = &newMonster
	}
	return Game{writer: writer, world: world, maxTurns: maxTurns, currentTurn: 0, monsters: monsters}, nil
}

/*
Returns true if the world has been completely destroyed (no cities left).
 */
func (g Game) WorldIsDestroyed() bool {
	return len(g.world.GetCities()) == 0
}

/*
Plays a new turn (= moves all monsters and make them fight), and return true if the game is over.
Monsters don't fight on first turn, they move before fighting.
Game can be over in four ways:
	1) The maximum number of turns has been reached
	2) The world is destroyed
	3) There is no monster left
	4) No monster has been able to move this turn (because they're all stuck in cities with no roads)
 */
func (g *Game) PlayTurn() bool {
	g.currentTurn++
	citiesWithMonsters, aMonsterMoved := g.moveMonsters()

	for _, city := range citiesWithMonsters {
		g.makeFightIfNecessary(city)
	}

	return !aMonsterMoved || g.currentTurn >= g.maxTurns || len(g.monsters) == 0 || g.WorldIsDestroyed()
}

func (g *Game) moveMonsters() (map[model.CityId]*model.City, bool) {
	citiesWithMonsters := make(map[model.CityId]*model.City, len(g.world.GetCities()))
	aMonsterMoved := false
	for _, monster := range g.monsters {
		previousPosition := monster.GetPosition()
		newPosition := monster.Move()
		if newPosition != previousPosition {
			aMonsterMoved = true
		}
		citiesWithMonsters[newPosition.GetId()] = newPosition
	}
	return citiesWithMonsters, aMonsterMoved
}

func (g *Game) makeFightIfNecessary(city *model.City) {
	monsters := city.GetMonsters()
	if len(monsters) >= 2 {
		g.world.DestroyCity(city.GetId())

		monsterNames := make([]string, len(monsters))
		i := 0
		for monsterId := range monsters {
			delete(g.monsters, monsterId)
			monsterNames[i] = string(monsterId)
			i++
		}

		var stringBuilder strings.Builder
		stringBuilder.WriteString(string(city.GetId()) + " has been destroyed by " + monsterNames[0])
		if len(monsterNames) >= 3 {
			stringBuilder.WriteString(", ")
			stringBuilder.WriteString(strings.Join(monsterNames[1:len(monsterNames)-1], ", "))
		}
		stringBuilder.WriteString(" and " + monsterNames[len(monsterNames)-1] + "!")
		fmt.Fprintln(g.writer, stringBuilder.String())
	}
}

func getCityIds(world *model.Map) []model.CityId {
	cities := world.GetCities()
	cityIds := make([]model.CityId, len(cities))
	i := 0
	for cityId := range cities {
		cityIds[i] = cityId
		i++
	}
	return cityIds
}
