package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	KingValue = 0
)

type King struct {
	c   types.Color
	pos int
}

var _ interfaces.IPiece = (*King)(nil)

func NewKing(c types.Color, pos int) *King {
	return &King{
		c:   c,
		pos: pos,
	}
}

func (k *King) GetValue() int {
	return KingValue
}

func (k *King) GetColor() types.Color {
	return k.c
}

func (k *King) GetPos() int {
	return k.pos
}
