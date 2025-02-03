package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	KnightValue = 3
)

type Knight struct {
	c   types.Color
	pos int
}

var _ interfaces.IPiece = (*Knight)(nil)

func NewKnight(color types.Color, pos int) *Knight {
	return &Knight{
		c:   color,
		pos: pos,
	}
}

func (k *Knight) GetValue() int {
	return KnightValue
}

func (k *Knight) GetColor() types.Color {
	return k.c
}

func (k *Knight) GetPos() int {
	return k.pos
}
