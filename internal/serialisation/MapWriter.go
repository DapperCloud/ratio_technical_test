package serialisation

import (
	"io"
	"ratio_technical_test/internal/model"
	"strings"
	"fmt"
	"errors"
)

type MapWriter struct {
	writer io.Writer
}

func NewMapWriter(writer io.Writer) MapWriter {
	return MapWriter{writer: writer}
}

func (w *MapWriter) WriteMap(modelMap model.Map) error {
	modelCities := modelMap.GetCities()
	for _, city := range modelCities {
		n, err := fmt.Fprintln(w.writer, cityToString(*city))
		if n <= 0 {
			return errors.New(fmt.Sprintf("Nothing written for city %v", city.GetId()))
		}
		if err != nil {
			return errors.New(fmt.Sprintf("Error while writing city %v: %v", city, err))
		}
	}
	return nil
}

func cityToString(city model.City) string {
	var builder strings.Builder
	builder.WriteString(string(city.GetId()))
	roads := city.GetRoads()
	for _, road := range roads {
		builder.WriteString(fmt.Sprintf(" %v=%v", road.Direction, road.Destination))
	}
	return builder.String()
}