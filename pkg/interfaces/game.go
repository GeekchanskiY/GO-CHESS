package interfaces

type IGame interface {
	GetBoard() IBoard

	StartGame(code int) error
	FinishGame(code int) error

	IsGameStarted() bool
	IsGameAlive() bool
	IsGameOver() bool
}
