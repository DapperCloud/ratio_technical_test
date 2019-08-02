package model

type Direction string

const (
	West  Direction = "west"
	East  Direction = "east"
	North Direction = "north"
	South Direction = "south"
)

func (d Direction) opposite() Direction {
	switch d {
	case West:
		return East
	case East:
		return West
	case North:
		return South
	case South:
		return North
	default:
		panic(d + "is not a valid direction")
	}
}
