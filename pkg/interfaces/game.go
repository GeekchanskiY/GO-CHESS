package interfaces

// IGame represents game, player connection & bot management
// TODO: add timer
type IGame interface {
	GetBoard() IBoard

	StartGame(code int) error
	FinishGame(code int) error

	IsGameStarted() bool
	IsGameAlive() bool
	IsGameOver() bool
}
