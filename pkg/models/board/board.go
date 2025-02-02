package board

import "GO-CHESS/pkg/models/pieces"

const (
	// Size is a board size. Board positions start from 0
	Size         = 8*8 - 1
	PiecesAmount = 8 * 2
)

type IBoard interface {
	InitBoard() (bool, error)
	Draw()

	GetPieces() []pieces.IPiece

	GetMoves() []pieces.IMovement
}

type Board struct {
	pieces []pieces.IPiece
	moves  []pieces.IMovement
}

func (b *Board) NewBoard() Board {
	initialPieces := make([]pieces.IPiece, 0, PiecesAmount)

	// Set pawns
	for i := 0; i < 8; i++ {
		blackPawn := pieces.NewPawn(0, pieces.Black)
		whitePawn := pieces.NewPawn(0, pieces.White)

		initialPieces = append(initialPieces, &blackPawn)
		initialPieces = append(initialPieces, &whitePawn)
	}
	return Board{}
}

func (b *Board) GetMoves() []pieces.IMovement {
	return b.moves
}

func (b *Board) GetPieces() []pieces.IPiece {
	return b.pieces
}
