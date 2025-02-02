package pieces

type IPiece interface {
	Move(pos int) (bool, error)
	GetMoves() ([]int, error)

	// ValidateMove check if move is valid
	ValidateMove(pos int) bool

	String() string
}
