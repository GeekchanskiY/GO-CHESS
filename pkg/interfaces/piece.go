package interfaces

import (
	"GO-CHESS/pkg/types"
)

// IPiece represents different pieces behavior
type IPiece interface {
	GetValue() int
	GetColor() types.Color
}
