package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	KnightValue = 3
)

type Knight struct {
	C types.Color
}

var _ interfaces.IPiece = (*Knight)(nil)

func (k *Knight) GetValue() int {
	return KnightValue
}

func (k *Knight) GetColor() types.Color {
	return k.C
}
