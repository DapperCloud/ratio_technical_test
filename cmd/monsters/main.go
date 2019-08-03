package main

import (
	"os"
	"fmt"
	"strconv"
	"ratio_technical_test/internal/serialisation"
	"ratio_technical_test/internal/game"
	"math/rand"
	"time"
)

const (
	MaxTurns = 10000
)

func main() {
	if len(os.Args) != 3 {
		showUsageAndExit()
	}
	mapFilePath := os.Args[1]
	monstersNumber, err := strconv.Atoi(os.Args[2])
	if err != nil {
		errorAndExit(fmt.Sprintf("error while reading monsters number: %v", err))
		showUsageAndExit()
	}
	if monstersNumber <= 0 {
		fmt.Println("monsters number should be a positive integer")
		showUsageAndExit()
	}

	file, err := os.Open(mapFilePath)
	if err != nil {
		errorAndExit(fmt.Sprintf("error while reading the map file: %v", err))
	}
	mapReader := serialisation.NewMapReader(file)
	worldMap, err := mapReader.GetMap()
	if err != nil {
		errorAndExit(fmt.Sprintf("error while reading the map file: %v", err))
	}

	rand.Seed(time.Now().Unix())
	gameWriter := os.Stdout
	theGame, err := game.NewGame(gameWriter, &worldMap, MaxTurns, uint(monstersNumber))
	if err != nil {
		errorAndExit(fmt.Sprintf("error while creating the game: %v", err))
	}

	for !theGame.PlayTurn() {
		// The game is playing
	}

	if theGame.WorldIsDestroyed() {
		fmt.Fprintln(gameWriter, "The world was completely destroyed! :(")
	} else {
		fmt.Fprintln(gameWriter, "The world has survived! Here is the new world map:")
		mapWriter := serialisation.NewMapWriter(gameWriter)
		err := mapWriter.WriteMap(worldMap)
		if err != nil {
			errorAndExit(fmt.Sprintf("error while writing the map: %v", err))
		}
	}
}

func showUsageAndExit() {
	fmt.Printf("\nUSAGE: monsters [map file] [monsters number]\n" +
		"map file - path to text file containing the world map\n" +
		"monsters number - number of monsters to spawn (positive integer)\n")
	os.Exit(1)
}

func errorAndExit(errorMessage string) {
	fmt.Println(errorMessage)
	os.Exit(1)
}