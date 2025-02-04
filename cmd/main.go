package main

import (
	"fmt"

	"GO-CHESS/pkg/models"
)

func main() {
	game := models.NewGame("", "")
	fmt.Println(game.GetPlayers()[0].Name())
}
