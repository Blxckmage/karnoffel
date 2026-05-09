package main

import (
	"fmt"

	"karnoffel/internal/game"
)

func main() {
	game := game.NewGame()

	fmt.Printf("Game TrumpCard: %v\n", game.TrumpCard)

	fmt.Printf("Player 1 hand: %v\n", game.Player1.Hand)
	fmt.Printf("Player 2 hand: %v\n", game.Player2.Hand)

	fmt.Printf("Player 1 points: %d\n", game.Player1.Point)
	fmt.Printf("Player 2 points: %d\n", game.Player2.Point)
}
