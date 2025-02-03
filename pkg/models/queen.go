package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	QueenValue = 9
)

type Queen struct {
	C types.Color
}

var _ interfaces.IPiece = (*Queen)(nil)

func (q *Queen) GetValue() int {
	return QueenValue
}

func (q *Queen) GetColor() types.Color {
	return q.C
}
