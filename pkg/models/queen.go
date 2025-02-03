package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	QueenValue = 9
)

type Queen struct {
	c   types.Color
	pos int
}

var _ interfaces.IPiece = (*Queen)(nil)

func NewQueen(c types.Color, pos int) *Queen {
	return &Queen{
		c:   c,
		pos: pos,
	}
}

func (q *Queen) GetValue() int {
	return QueenValue
}

func (q *Queen) GetColor() types.Color {
	return q.c
}

func (q *Queen) GetPos() int {
	return q.pos
}
