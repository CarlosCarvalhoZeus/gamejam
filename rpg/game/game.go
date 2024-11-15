package game

import (
	"bufio"
	"os"
)

type GameUI interface {
	Draw(*Level)
}
type Tile rune

const (
	StoneWall  Tile = '#'
	DirtFloor  Tile = '.'
	Door       Tile = '|'
	BlankSpace Tile = ' '
)

type Level struct {
	Map [][]Tile
}

func loadLevelFromFile(filename string) *Level {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	levelLines := make([]string, 0)
	longestRow := 0
	index := 0
	for scanner.Scan() {
		levelLines = append(levelLines, scanner.Text())
		if len(levelLines[index]) > longestRow {
			longestRow = len(levelLines[index])
		}
		index++
	}
	level := &Level{}
	level.Map = make([][]Tile, len(levelLines))
	for i := range level.Map {
		level.Map[i] = make([]Tile, longestRow)
	}
	// for y := 0; y < len(level.Map); y++ {
	// 	line := levelLines[y]
	// 	for x := 0; x < longestRow; x++ {
	// 		for _, c := range line {
	// 			if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
	// 				level.Map[y][x] = BlankSpace
	// 			} else {
	// 				level.Map[y][x] = Tile(c)
	// 			}
	// 		}
	// 	}
	//
	// }
	return level
}

func Run(ui GameUI) {
	level := loadLevelFromFile("/home/zeus/Área de Trabalho/gamejam/rpg/game/maps/level1.map")
	ui.Draw(level)
}
