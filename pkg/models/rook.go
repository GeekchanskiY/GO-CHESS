package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	RookValue = 5
)

type Rook struct {
	c   types.Color
	pos int
}

var _ interfaces.IPiece = (*Rook)(nil)

func NewRook(c types.Color, pos int) *Rook {
	return &Rook{
		c:   c,
		pos: pos,
	}
}

func (r *Rook) GetValue() int {
	return RookValue
}

func (r *Rook) GetColor() types.Color {
	return r.c
}

func (r *Rook) GetPos() int {
	return r.pos
}
