package serialisation

import (
	"ratio_technical_test/internal/model"
	"bufio"
	"io"
	"strings"
	"errors"
	"fmt"
)

type MapReader struct {
	reader io.Reader
}

func NewMapReader(reader io.Reader) MapReader {
	return MapReader{reader: reader}
}

func (m MapReader) GetMap() (model.Map, error) {
	resultMap := model.NewMap()
	cities := make([]City, 1000)

	scanner := bufio.NewScanner(m.reader)
	scanner.Split(bufio.ScanLines)
	lineNumber := 1

	for scanner.Scan() {
		cityString := scanner.Text()
		city, err := cityFromString(cityString)
		if err != nil {
			return model.Map{}, errors.New(fmt.Sprintf("Error reading line %v: %v", lineNumber, err.Error()))
		}
		cities = append(cities, city)
		resultMap.AddCity(city.id)
		lineNumber++
	}

	for index, city := range cities {
		for _, road := range city.roads {
			err := resultMap.AddRoad(city.id, road.destination, road.direction)
			if err != nil {
				return model.Map{}, errors.New(fmt.Sprintf("Error on line %v: %v", index, err.Error()))
			}
		}
	}
	return resultMap, nil
}

func cityFromString(cityString string) (City, error) {
	tokens := strings.Split(cityString, " ")
	if len(tokens) < 2 {
		return City{}, errors.New("A city should have at least a name and one road")
	}
	cityId := tokens[0]
	tokens = tokens[1:]
	roads := make([]Road, len(tokens))
	for i, roadString := range tokens {
		roadTokens := strings.Split(roadString, "=")
		if len(roadTokens) != 2 {
			return City{}, errors.New(fmt.Sprintf("A road should follow the format direction=city, cannot parse '%v'", roadString))
		}
		direction, err := directionFromString(roadTokens[0])
		if err != nil {
			return City{}, err
		}
		destination := roadTokens[1]
		roads[i] = Road{direction: direction, destination: model.CityId(destination)}
	}
	return NewCity(model.CityId(cityId), roads), nil
}

func directionFromString(direction string) (model.Direction, error) {
	switch direction {
	case "west":
		return model.West, nil
	case "east":
		return model.East, nil
	case "south":
		return model.South, nil
	case "north":
		return model.North, nil
	default:
		return "", errors.New(fmt.Sprintf("Cannot parse direction '%v', must be one of: west,east,south,north", direction))
	}
}
