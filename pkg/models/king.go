package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	KingValue = 0
)

type King struct {
	C types.Color
}

var _ interfaces.IPiece = (*King)(nil)

func (k *King) GetValue() int {
	return KingValue
}

func (k *King) GetColor() types.Color {
	return k.C
}
