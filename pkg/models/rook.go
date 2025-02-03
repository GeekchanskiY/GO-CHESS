package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	RookValue = 5
)

type Rook struct {
	C types.Color
}

var _ interfaces.IPiece = (*Rook)(nil)

func (r *Rook) GetValue() int {
	return RookValue
}

func (r *Rook) GetColor() types.Color {
	return r.C
}
