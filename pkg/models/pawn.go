package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	PawnValue = 1
)

type Pawn struct {
	c   types.Color
	pos int
}

var _ interfaces.IPiece = (*Pawn)(nil)

func NewPawn(color types.Color, pos int) *Pawn {
	return &Pawn{
		c:   color,
		pos: pos,
	}
}

func (p *Pawn) GetValue() int {
	return PawnValue
}

func (p *Pawn) GetColor() types.Color {
	return p.c
}

func (p *Pawn) GetPos() int {
	return p.pos
}
