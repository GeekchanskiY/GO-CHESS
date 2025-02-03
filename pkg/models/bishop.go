package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	BishopValue = 3
)

type Bishop struct {
	C types.Color
}

var _ interfaces.IPiece = (*Bishop)(nil)

func NewBishop(color types.Color) *Bishop {
	return &Bishop{
		C: color,
	}
}

func (b *Bishop) GetValue() int {
	return BishopValue
}

func (b *Bishop) GetColor() types.Color {
	return b.C
}
