package pieces

import "errors"

var (
	ErrorInvalidMove = errors.New("move invalid")
	ErrorNoMoves     = errors.New("no moves available")
)
