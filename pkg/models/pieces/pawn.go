package pieces

import (
	"fmt"

	"GO-CHESS/pkg/models/board"
)

type Pawn struct {
	pos   int
	color int
}

// force interface implementation
var _ IPiece = (*Pawn)(nil)

func NewPawn(pos, color int) Pawn {
	return Pawn{
		pos:   pos,
		color: color,
	}
}

func (p *Pawn) IsMoveValid(b *board.IBoard, pos int) bool {
	return true
}

func (p *Pawn) Move(b *board.IBoard, pos int) (bool, error) {
	return false, nil
}

func (p *Pawn) GetMoves(b *board.IBoard) []int {
	return nil
}

func (p *Pawn) GetColor() int {
	return p.color

}

func (p *Pawn) String() string {
	return fmt.Sprintf("Pawn at position %d", p.pos)
}
