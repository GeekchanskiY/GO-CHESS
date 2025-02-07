package models

import (
	"GO-CHESS/pkg/interfaces"
)

type Game struct {
	board   interfaces.IBoard
	players []*Player
}

func NewGame(pName1, pName2 string) *Game {

	players := make([]*Player, 2)
	players[0] = NewPlayer(pName1)
	players[1] = NewPlayer(pName2)

	board := NewBoard()
	board.InitBoard()

	return &Game{
		players: players,
		board:   board,
	}
}

func (g *Game) GetBoard() interfaces.IBoard {
	return g.board
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}
