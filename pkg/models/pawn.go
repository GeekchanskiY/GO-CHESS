package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	PawnValue = 1
)

type Pawn struct {
}

var _ interfaces.IPiece = (*Pawn)(nil)

func (p *Pawn) GetValue() int {
	return PawnValue
}
