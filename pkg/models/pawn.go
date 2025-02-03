package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	PawnValue = 1
)

type Pawn struct {
	C types.Color
}

var _ interfaces.IPiece = (*Pawn)(nil)

func (p *Pawn) GetValue() int {
	return PawnValue
}

func (p *Pawn) GetColor() types.Color {
	return p.C
}
