package board

import "GO-CHESS/pkg/models/pieces"

type IBoard interface {
	InitBoard() (bool, error)
	Draw()

	GetPieces() []pieces.IPiece

	GetMoves() []pieces.IMovement
}

type Board struct {
}
