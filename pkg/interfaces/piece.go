package interfaces

import (
	"GO-CHESS/pkg/types"
)

type IPiece interface {
	GetValue() int
	GetColor() types.Color
}
