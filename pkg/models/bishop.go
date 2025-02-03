package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	BishopValue = 3
)

type Bishop struct {
	C   types.Color
	Pos int
}

var _ interfaces.IPiece = (*Bishop)(nil)

func NewBishop(color types.Color, pos int) *Bishop {
	return &Bishop{
		C:   color,
		Pos: pos,
	}
}

func (b *Bishop) GetValue() int {
	return BishopValue
}

func (b *Bishop) GetColor() types.Color {
	return b.C
}

func (b *Bishop) GetPos() int {
	return b.Pos
}
