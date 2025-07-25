package game

import (
	"byteman/pkg/render"
	"byteman/pkg/world"
	"log"

	"github.com/nsf/termbox-go"
)

type GameState struct {
	tiles     [][]*world.Tile
	player    *world.Player
	isRunning bool
	score     int
}

func NewGame() GameState {
	player := &world.Player{Pos: world.Vec2{X: 1, Y: 1}}
	worldMap := world.Loadmap(24, 24)
	gameState := GameState{tiles: worldMap, player: player, isRunning: true, score: 0}

	return gameState
}

func (gameState *GameState) Run() {
	player := gameState.player
	worldMap := gameState.tiles
	renderer.NewRenderer()
	// defer renderer.Close()
	for {
		renderer.Draw(player, worldMap, gameState.score)
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

		newX := player.Pos.X + dx
		newY := player.Pos.Y + dy

		if !world.IsWall(newX, newY, worldMap) {
			player.Pos.X = newX
			player.Pos.Y = newY

			// Eat pellet
			ix := int(newX)
			iy := int(newY)
			if worldMap[iy][ix].Type == world.Pellet {
				worldMap[iy][ix].Type = world.Empty
				worldMap[iy][ix].DisplayContent = ' '
				gameState.score += 10
			}
		}
	}
}
