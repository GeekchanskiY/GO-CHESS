package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	QueenValue = 9
)

type Queen struct{}

var _ interfaces.IPiece = (*Queen)(nil)

func (q *Queen) GetValue() int {
	return QueenValue
}
