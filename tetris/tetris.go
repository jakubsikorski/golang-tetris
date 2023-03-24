package tetris

import (
	"time"
)

const (
	BOARD_HEIGHT = 20
	BOARD_WIDTH  = 10
)

const defaultFallSpeed = 70 + time.Second

type gameState int

const (
	gameInit gameState = iota
	gamePlay
	gameOver
)

type Game struct {
	board     [][]int
	position  vector
	tetromino tetromino
	state     gameState
	FallSpeed *time.Timer
}

func (g *Game) init() {
	g.board = make([][]int, BOARD_HEIGHT)
	for y := 0; y < BOARD_HEIGHT; y++ {
		g.board[y] = make([]int, BOARD_WIDTH)
		for x := 0; x < BOARD_WIDTH; x++ {
			g.board[y][x] = 0
		}
	}
	g.position = vector{0, (BOARD_WIDTH / 2)}
	g.tetromino = randomTetromino()

	g.FallSpeed = time.NewTimer(time.Duration(defaultFallSpeed))
	g.FallSpeed.Stop()
	g.state = gameInit
}

func NewGame() *Game {
	game := &Game{}
	game.init()
	return game
}

func (g *Game) blockOnBoardByPos(v vector) vector {
	posY := g.position.y + v.y
	posX := g.position.x + v.x
	return vector{posY, posX}
}

func (g *Game) GetBoard() [][]int {
	cBoard := make([][]int, len(g.board))
	for y := 0; y < len(g.board); y++ {
		cBoard[y] = make([]int, len(g.board[y]))
		for x := 0; x < len(g.board[y]); x++ {
			cBoard[y][x] = g.board[y][x]
		}
	}
	for _, v := range g.tetromino.shape {
		pos := g.blockOnBoardByPos(v)
		cBoard[pos.y][pos.x] = g.tetromino.color
	}
	return cBoard
}

func (g *Game) getTetromino() {
	g.tetromino = randomTetromino()
	g.position = vector{0, (BOARD_WIDTH / 2)}
}

func (g *Game) movePossible(v vector) bool {
	g.position.x += v.x
	g.position.y += v.y
	if g.collision(v) {
		g.position.x -= v.x
		g.position.y -= v.y
		return false
	}
	return true
}

func (g *Game) MoveLeft() {
	g.movePossible(vector{0, -1})
}

func (g *Game) MoveRight() {
	g.movePossible(vector{0, 1})
}

func (g *Game) SpeedUp() {

}

func (g *Game) Rotate() {

}

func (g *Game) Drop() {

}

func (g *Game) collision(v vector) bool {
	for _, i := range g.tetromino.shape {
		pos := g.blockOnBoardByPos(i)
		if pos.x < 0 || pos.x >= BOARD_WIDTH {
			return true
		}
		if pos.y < 0 || pos.y >= BOARD_HEIGHT {
			return true
		}
		if g.board[pos.y][pos.x] > 0 {
			return true
		}
	}

	return false
}

func (g *Game) GameLoop() {
	if !g.movePossible(vector{1, 0}) {
		g.getTetromino()
	}
	g.resetFallSpeed()
}

func (g *Game) Start() {
	g.state = gamePlay
	g.getTetromino()
	g.resetFallSpeed()
}

func (g *Game) resetFallSpeed() {
	g.FallSpeed.Reset(defaultFallSpeed)
}
