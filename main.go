package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Pellet
	PowerUp
)

type Vec2 struct {
	x, y float64
}

type Tile struct {
	Pos            Vec2
	Type           TileType
	displayContent rune
}

type Player struct {
	Pos Vec2
}

func isWall(x, y float64) bool {
	ix := int(x)
	iy := int(y)

	if iy < 0 || iy >= len(tileMap) || ix < 0 || ix >= len(tileMap[iy]) {
		return true
	}

	return tileMap[iy][ix].Type == Wall
}

var (
	DX = map[string]int{"N": 0, "S": 0, "W": -1, "E": 1}
	DY = map[string]int{"N": -1, "S": 1, "W": 0, "E": 0}
)

func carvePath(x, y int, grid [][]Tile) {
	directions := []string{"N", "S", "E", "W"}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		nx, ny := x+DX[dir]*2, y+DY[dir]*2

		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) && grid[ny][nx].Type == Wall {
			grid[y+DY[dir]][x+DX[dir]].Type = Empty
			grid[y+DY[dir]][x+DX[dir]].displayContent = ' '
			grid[ny][nx].Type = Empty
			grid[ny][nx].displayContent = ' '
			carvePath(nx, ny, grid)
		}
	}
}

func loadMap(width int, height int) [][]Tile {
	tiles := make([][]Tile, height)
	for y := 0; y < height; y++ {
		tiles[y] = make([]Tile, width)
		for x := 0; x < width; x++ {
			tiles[y][x] = Tile{
				Pos:            Vec2{x: float64(x), y: float64(y)},
				Type:           Wall,
				displayContent: '#',
			}
		}
	}

	// Start carving from a random odd position
	startX := rand.Intn(width/2)*2 + 1
	startY := rand.Intn(height/2)*2 + 1
	tiles[startY][startX].Type = Empty
	tiles[startY][startX].displayContent = ' '

	carvePath(startX, startY, tiles)

	// Place pellets
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if tiles[y][x].Type == Empty {
				tiles[y][x].Type = Pellet
				tiles[y][x].displayContent = '.'
			}
		}
	}

	return tiles
}

var tileMap [][]Tile

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	tileMap = loadMap(25, 25)

	player := Player{Vec2{x: 1, y: 1}}
	// Find a valid starting position for the player
	for {
		px := rand.Intn(len(tileMap[0]))
		py := rand.Intn(len(tileMap))
		if !isWall(float64(px), float64(py)) {
			player.Pos.x = float64(px)
			player.Pos.y = float64(py)
			break
		}
	}

	for {
		draw(player)

		var dx, dy float64 = 0, 0

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
		case termbox.EventError:
			log.Fatal(ev.Err)
		}

		newX := player.Pos.x + dx
		newY := player.Pos.y + dy

		if !isWall(newX, newY) {
			player.Pos.x = newX
			player.Pos.y = newY

			// Eat pellet
			ix := int(newX)
			iy := int(newY)
			if tileMap[iy][ix].Type == Pellet {
				tileMap[iy][ix].Type = Empty
				tileMap[iy][ix].displayContent = ' '
			}
		}
	}
}

func draw(player Player) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw map
	for y, row := range tileMap {
		for x, cell := range row {
			var fg termbox.Attribute
			ch := cell.displayContent
			fg = termbox.ColorWhite

			switch cell.Type {
			case Wall:
				fg = termbox.ColorRed
			case Pellet:
				fg = termbox.ColorWhite
			}

			termbox.SetCell(x, y, ch, fg, termbox.ColorBlack)
		}
	}

	termbox.SetCell(int(player.Pos.x), int(player.Pos.y), 'P', termbox.ColorYellow|termbox.AttrBold, termbox.ColorBlack)

	termbox.Flush()
}

