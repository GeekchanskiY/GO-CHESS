package pieces

import "GO-CHESS/pkg/models/board"

const (
	Black = iota
	White
)

type IPiece interface {
	Move(b *board.IBoard, pos int) (bool, error)
	GetMoves(b *board.IBoard) ([]int, error)
	IsMoveValid(b *board.IBoard, pos int) bool

	// GetColor returns piece color (Black or White)
	GetColor() int

	String() string
}
