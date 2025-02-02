package pieces

type IMovement interface {
	Validate() error
}

type Movement struct {
	initPos int
	newPos  int

	piece IPiece
}

func NewMovement(piece IPiece, new int) *Movement {
	return &Movement{}
}
