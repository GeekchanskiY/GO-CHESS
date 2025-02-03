package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	KnightValue = 3
)

type Knight struct {
}

var _ interfaces.IPiece = (*Knight)(nil)

func (k *Knight) GetValue() int {
	return KnightValue
}
