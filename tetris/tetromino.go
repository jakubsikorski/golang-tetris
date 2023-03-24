package tetris

import (
	"math/rand"
)

type vector struct {
	y, x int
}

type tetromino struct {
	shape []vector
	color int
}

var tetrominos = []tetromino{
	{
		// L shape
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
		color: 0,
	},
	{
		// J shape
		shape: []vector{{1, -1}, {0, 0}, {0, 1}, {0, -1}},
		color: 1,
	},
	{
		// T shape
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {1, 0}},
		color: 2,
	},
	{
		// O shape
		shape: []vector{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		color: 3,
	},
	{
		// S shape
		shape: []vector{{0, 0}, {0, 1}, {1, -1}, {1, 0}},
		color: 4,
	},
	{
		// Z shape
		shape: []vector{{0, -1}, {0, 0}, {1, 0}, {1, 1}},
		color: 5,
	},
}

func randomTetromino() tetromino {
	idx := rand.Intn(len(tetrominos) - 1)
	return tetrominos[idx]
}
