package tetris

import (
	"math"
	"math/rand"
)

type vector struct {
	y, x int
}

type tetromino struct {
	shape     []vector
	color     int
	canRotate bool
}

var tetrominos = []tetromino{
	{
		// L shape
		shape:     []vector{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
		color:     0,
		canRotate: true,
	},
	{
		// J shape
		shape:     []vector{{1, -1}, {0, 0}, {0, 1}, {0, -1}},
		color:     1,
		canRotate: true,
	},
	{
		// T shape
		shape:     []vector{{0, -1}, {0, 0}, {0, 1}, {1, 0}},
		color:     2,
		canRotate: true,
	},
	{
		// O shape
		shape:     []vector{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		color:     3,
		canRotate: false,
	},
	{
		// S shape
		shape:     []vector{{0, 0}, {0, 1}, {1, -1}, {1, 0}},
		color:     4,
		canRotate: true,
	},
	{
		// Z shape
		shape:     []vector{{0, -1}, {0, 0}, {1, 0}, {1, 1}},
		color:     5,
		canRotate: true,
	},
}

func randomTetromino() tetromino {
	idx := rand.Intn(len(tetrominos)-1) + 1
	return tetrominos[idx]
}
func (t *tetromino) rotateWithAngle(ang float64) {
	cos := int(math.Round(math.Cos(ang)))
	sin := int(math.Round(math.Sin(ang)))
	for i, e := range t.shape {
		ny := e.y*cos - e.x*sin
		nx := e.y*sin - e.x*cos

		t.shape[i] = vector{ny, nx}
	}
}

func (t *tetromino) rotateBack() {
	ang := math.Pi / 2 * 3
	t.rotateWithAngle(ang)
}

func (t *tetromino) rotate() {
	ang := math.Pi / 2
	if !t.canRotate {
		return
	}
	t.rotateWithAngle(ang)
}
