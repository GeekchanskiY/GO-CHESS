package board

import "GO-CHESS/pkg/models/pieces"

type IBoard interface {
	Draw()

	GetPieces() []pieces.IPiece
}
