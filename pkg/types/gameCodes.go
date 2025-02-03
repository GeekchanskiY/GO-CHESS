package types

type GameCode int

const (
	TestGameStart GameCode = 100
	BotGameStart  GameCode = 101
	PvPGameStart  GameCode = 102

	GameFinishWhiteWin   = 200
	GameFinishBlackWin   = 201
	GameFinishDraw       = 203
	GameFinishPlayerQuit = 204
)
