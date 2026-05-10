package main

import (
	"fmt"

	"karnoffel/internal/game"
)

func main() {
	g := game.NewGame()

	fmt.Println("=== Game Start ===")
	fmt.Printf("Trump Card: %s\n", g.TrumpCard.CardToString())
	fmt.Println()

	fmt.Println("=== Opening Hands ===")
	fmt.Printf("Player 1 opening card: %s\n", g.Player1.OpeningCard.CardToString())
	fmt.Printf("Player 2 opening card: %s\n", g.Player2.OpeningCard.CardToString())
	fmt.Println()

	fmt.Println("=== Player 1 Cards ===")
	for i, card := range g.Player1.Hand {
		fmt.Printf("%d: %s\n", i+1, card.CardToString())
	}
	fmt.Println()

	fmt.Println("=== Player 2 Cards ===")
	for i, card := range g.Player2.Hand {
		fmt.Printf("%d: %s\n", i+1, card.CardToString())
	}
	fmt.Println()

	for g.Player1.GamePoint < 12 && g.Player2.GamePoint < 12 {
		fmt.Println("=== New Hand ===")
		firstPlayer, otherPlayer := g.DetermineFirstPlayer()

		for i := 0; i < 4; i++ {
			fmt.Printf("\n--- Trick %d ---\n", i+1)
			g.PlayRound(firstPlayer, false)
			g.PlayRound(otherPlayer, true)
			g.ResolveRound(firstPlayer, otherPlayer)
		}

		g.ResolveHand(firstPlayer, otherPlayer)

		if g.Player1.GamePoint >= 12 || g.Player2.GamePoint >= 12 {
			break
		}

		g.Deck = game.NewDeckOnly()
		g.Deck.Shuffle()
		g.DealNewHand()
	}

	fmt.Println("=== Game Over ===")
	if g.Player1.GamePoint > g.Player2.GamePoint {
		fmt.Printf("Player 1 wins! Final score: %d - %d\n", g.Player1.GamePoint, g.Player2.GamePoint)
	} else {
		fmt.Printf("Player 2 wins! Final score: %d - %d\n", g.Player2.GamePoint, g.Player1.GamePoint)
	}
}
