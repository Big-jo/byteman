package world

import (
	"math/rand"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Pellet
	PowerUp
)

type Vec2 struct {
	X, Y float64
}

type Tile struct {
	Pos            Vec2
	Type           TileType
	DisplayContent rune
}

type Player struct {
	Pos Vec2
}

type WorldMap [][]string

func carvePath(x, y int, grid [][]*Tile) {
	var (
		DX = map[string]int{"N": 0, "S": 0, "W": -1, "E": 1}
		DY = map[string]int{"N": -1, "S": 1, "W": 0, "E": 0}
	)

	directions := []string{"N", "S", "E", "W"}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		nx, ny := x+DX[dir]*2, y+DY[dir]*2

		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) && grid[ny][nx].Type == Wall {
			grid[y+DY[dir]][x+DX[dir]].Type = Empty
			grid[y+DY[dir]][x+DX[dir]].DisplayContent = ' '
			grid[ny][nx].Type = Empty
			grid[ny][nx].DisplayContent = ' '
			carvePath(nx, ny, grid)
		}
	}
}

func IsWall(x, y float64, tileMap [][]*Tile) bool {
	ix := int(x)
	iy := int(y)

	if iy < 0 || iy >= len(tileMap) || ix < 0 || ix >= len(tileMap[iy]) {
		return true
	}

	return tileMap[iy][ix].Type == Wall
}

func Loadmap(width int, height int) [][]*Tile {
	tiles := make([][]*Tile, height)
	for y := 0; y < height; y++ {
		tiles[y] = make([]*Tile, width)
		for x := 0; x < width; x++ {
			tiles[y][x] = &Tile{
				Pos:            Vec2{X: float64(x), Y: float64(y)},
				Type:           Wall,
				DisplayContent: '#',
			}
		}
	}

	// start carving from a random odd position
	startx := rand.Intn(width/2)*2 + 1
	starty := rand.Intn(height/2)*2 + 1
	tiles[starty][startx].Type = Empty
	tiles[starty][startx].DisplayContent = ' '

	carvePath(startx, starty, tiles)

	// place pellets
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if tiles[y][x].Type == Empty {
				tiles[y][x].Type = Pellet
				tiles[y][x].DisplayContent = '.'
			}
		}
	}

	return tiles
}
