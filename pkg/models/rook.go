package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	RookValue = 5
)

type Rook struct{}

var _ interfaces.IPiece = (*Rook)(nil)

func (r *Rook) GetValue() int {
	return RookValue
}
