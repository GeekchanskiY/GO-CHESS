package models

import (
	"GO-CHESS/pkg/models/interfaces"
)

const (
	KingValue = 0
)

type King struct{}

var _ interfaces.IPiece = (*King)(nil)

func (k *King) GetValue() int {
	return KingValue
}
