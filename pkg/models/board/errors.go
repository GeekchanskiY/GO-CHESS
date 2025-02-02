package board

import "errors"

var (
	GameAlreadyStarted  = errors.New("game already started")
	GameAlreadyFinished = errors.New("game already finished")

	ErrorOutsideBoard = errors.New("outside board")
)
