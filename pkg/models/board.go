package models

import (
	"GO-CHESS/pkg/interfaces"
	"GO-CHESS/pkg/types"
)

const (
	defaultPiecesAmount = 16
)

type Board struct {
	pieces []interfaces.IPiece
}

var _ interfaces.IBoard = (*Board)(nil)

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) InitBoard() {
	b.pieces = make([]interfaces.IPiece, 0, defaultPiecesAmount)

	// pawns init
	for _, color := range []types.Color{types.Black, types.White} {
		for i := 0; i < 8; i++ {
			newPawnPos := i

			if color == types.Black {
				newPawnPos += 8
			} else {
				newPawnPos += 48
			}

			newPawn := NewPawn(color, newPawnPos)
			b.pieces = append(b.pieces, newPawn)
		}
	}

	b.pieces = append(
		b.pieces,
		// Black pieces
		NewBishop(types.Black, 58),
		NewBishop(types.Black, 61),
		NewKnight(types.Black, 57),
		NewKnight(types.Black, 62),
		NewRook(types.Black, 56),
		NewRook(types.Black, 63),
		NewQueen(types.Black, 59),
		NewKing(types.Black, 60),

		// White pieces
		NewBishop(types.White, 2),
		NewBishop(types.White, 5),
		NewRook(types.White, 0),
		NewRook(types.White, 7),
		NewKnight(types.White, 1),
		NewKnight(types.White, 6),
		NewQueen(types.White, 3),
		NewKing(types.White, 4),
	)

}
