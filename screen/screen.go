package screen

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	offsetVer = 4
	offsetHor = 40
	cellWidth = 2
)

var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorBlue,
	termbox.ColorCyan,
	termbox.ColorWhite,
	termbox.ColorYellow,
	termbox.ColorRed,
}

type gameScreen struct{}

func (g *gameScreen) RenderAscii(board [][]int) {
	fmt.Println("\n ====================================================================")
	for _, r := range board {
		for _, num := range r {
			if num > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (g *gameScreen) Render(board [][]int) {
	termbox.Clear(termbox.ColorGreen, termbox.ColorGreen)

	for y, v := range board {
		for x, h := range v {
			color := colors[h]
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(cellWidth*x+offsetHor+i, y+offsetVer, ' ', color, color)
			}
		}
	}

	termbox.Flush()
}

func New() *gameScreen {
	return &gameScreen{}
}
