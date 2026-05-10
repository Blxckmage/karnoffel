package main

import (
	"fmt"

	"karnoffel/internal/game"
)

func main() {
	g := game.NewGame()

	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║         KARNOFFEL CARD GAME            ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("\nTrump Card: %s\n", g.TrumpCard.CardToString())

	fmt.Println("\n─── Opening Cards ───")
	fmt.Printf("P1: %s\n", g.Player1.OpeningCard.CardToString())
	fmt.Printf("P2: %s\n", g.Player2.OpeningCard.CardToString())

	fmt.Println("\n─── Your Hand (P1) ───")
	for i, card := range g.Player1.Hand {
		fmt.Printf("  %d: %s\n", i+1, card.CardToString())
	}

	fmt.Println("\n─── P2 Hand (hidden) ───")
	fmt.Printf("  P2 has %d cards\n", len(g.Player2.Hand))

	handNumber := 1
	for g.Player1.GamePoint < 12 && g.Player2.GamePoint < 12 {
		fmt.Printf("\n\n┌─ HAND %d ─\n", handNumber)
		firstPlayer, otherPlayer := g.DetermineFirstPlayer()

		for i := 0; i < 4; i++ {
			fmt.Printf("\nTrick %d:\n", i+1)
			g.PlayRound(firstPlayer, false)
			g.PlayRound(otherPlayer, true)
			g.ResolveRound(firstPlayer, otherPlayer)
		}

		g.ResolveHand(firstPlayer, otherPlayer)

		if g.Player1.GamePoint >= 12 || g.Player2.GamePoint >= 12 {
			break
		}

		fmt.Println("\n─── Next Hand ───")
		g.Deck = game.NewDeckOnly()
		g.Deck.Shuffle()
		g.DealNewHand()
		fmt.Printf("Trump Card: %s\n", g.TrumpCard.CardToString())

		fmt.Println("\n─── Your New Hand ───")
		for i, card := range g.Player1.Hand {
			fmt.Printf("  %d: %s\n", i+1, card.CardToString())
		}

		handNumber++
	}

	fmt.Println("\n\n╔════════════════════════════════════════╗")
	fmt.Println("║          GAME OVER - RESULTS           ║")
	fmt.Println("╚════════════════════════════════════════╝")
	if g.Player1.GamePoint > g.Player2.GamePoint {
		fmt.Printf("\n🎉 Player 1 WINS!\n")
		fmt.Printf("Final Score: %d - %d\n\n", g.Player1.GamePoint, g.Player2.GamePoint)
	} else {
		fmt.Printf("\n🎉 Player 2 WINS!\n")
		fmt.Printf("Final Score: %d - %d\n\n", g.Player2.GamePoint, g.Player1.GamePoint)
	}
}
