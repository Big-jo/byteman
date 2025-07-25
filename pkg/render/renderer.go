package renderer

import (
	"byteman/pkg/world"
	"log"

	"github.com/nsf/termbox-go"
)

func NewRenderer() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func Draw(player *world.Player, tiles [][]*world.Tile) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// draw map
	for y, row := range tiles {
		for x, cell := range row {
			var fg termbox.Attribute
			ch := cell.DisplayContent
			fg = termbox.ColorWhite

			switch cell.Type {
			case world.Wall:
				fg = termbox.ColorRed
			case world.Pellet:
				fg = termbox.ColorWhite
			}

			termbox.SetCell(x, y, ch, fg, termbox.ColorBlack)
		}
	}

	termbox.SetCell(int(player.Pos.X), int(player.Pos.Y), 'p', termbox.ColorYellow|termbox.AttrBold, termbox.ColorBlack)

	termbox.Flush()
}

func Close() {
	termbox.Close()
}
