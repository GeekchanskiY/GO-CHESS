package models

import (
	"GO-CHESS/pkg/interfaces"
)

type Move struct {
	piece  interfaces.IPiece
	oldPos int
	newPos int
}

func NewMove(piece interfaces.IPiece, newPos int, oldPos int) *Move {
	return &Move{
		oldPos: oldPos,
		piece:  piece,
		newPos: newPos,
	}
}

func (m *Move) GetData() (interfaces.IPiece, int, int) {
	return m.piece, m.oldPos, m.newPos
}
