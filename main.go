package main

import (
	// "fmt"
	"log"

	"github.com/nsf/termbox-go"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Pellet
	PowerUp
)

var gameMap = []string{
	"####################",
	"#........#.........#",
	"#........#.........#",
	"#........#.........#",
	"#........###########",
	"#..................#",
	"####################",
}

type Vec2 struct {
	x, y float64
}

type Tile struct {
	Pos  Vec2
	Type TileType
}

type Rect struct {
	Pos Vec2
}

type Player struct {
	Pos Vec2
}

type GameState struct {
	score int64
}

func isWall(x, y float64) bool {
	ix := int(x)
	iy := int(y)

	if iy < 0 || iy >= len(tileMap) || ix < 0 || ix > len(tileMap[iy]) {
		return true
	}

	return tileMap[iy][ix].Type == Wall
}

func loadMap(mapStrings []string) [][]Tile {
	height := len(mapStrings)
	width := len(mapStrings[0])
	tiles := make([][]Tile, height)

	for y := 0; y < height; y++ {
		tiles[y] = make([]Tile, width)
		for x, ch := range mapStrings[y] {
			tileType := Empty

			switch ch {
			case '#':
				tileType = Wall
			case '.':
				tileType = Pellet
			}

			tiles[y][x] = Tile{
				Pos:  Vec2{x: float64(x), y: float64(y)},
				Type: tileType,
			}
		}

	}

	return tiles
}

var tileMap [][]Tile

func main() {
	gameState := GameState{score: 0}
	tileMap = loadMap(gameMap)

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	player := Player{Vec2{x: 1, y: 2}}

	for {
		draw(player)
		// termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		// termbox.SetCell(int(player.x), int(player.y), 'P', termbox.ColorYellow, termbox.ColorBlack)
		// termbox.Flush()

		var dy, dx float64 = 0, 0

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				dy -= 1
			case termbox.KeyArrowDown:
				dy += 1
			case termbox.KeyArrowLeft:
				dx -= 1
			case termbox.KeyArrowRight:
				dx += 1
			case termbox.KeyEsc:
				return
			}
		}

		newX := player.Pos.x + dx
		newY := player.Pos.y + dy

		if !isWall(newX, newY) {
			player.Pos.x = newX
			player.Pos.y = newY
		}
	}
}

func draw(player Player) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw map
	for y, row := range tileMap {
		for x, cell := range row {
			var fg termbox.Attribute
			tile := cell

			ch := ' '
			fg = termbox.ColorWhite

			switch tile.Type {
			case Wall:
				ch = '#'
				fg = termbox.ColorRed
			case Pellet:
				ch = '.'
				fg = termbox.ColorWhite
			}

			termbox.SetCell(x, y, ch, fg, termbox.ColorBlack)
		}
	}

	termbox.SetCell(int(player.Pos.x), int(player.Pos.y), 'P', termbox.ColorYellow|termbox.AttrBold, termbox.ColorBlack)

	termbox.Flush()
}
