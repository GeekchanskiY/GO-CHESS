package main

import (
	"fmt"

	"GO-CHESS/pkg/models"
)

func main() {
	game := models.NewGame("", "")
	fmt.Println("Game players:")
	for _, player := range game.GetPlayers() {
		fmt.Println(player.Name(), player.IsBot())
	}
}
