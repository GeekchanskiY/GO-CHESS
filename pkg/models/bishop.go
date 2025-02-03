package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	BishopValue = 3
)

type Bishop struct{}

var _ interfaces.IPiece = (*Bishop)(nil)

func (b *Bishop) GetValue() int {
	return BishopValue
}
