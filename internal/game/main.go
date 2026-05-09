package game

import (
	"karnoffel/internal/deck"
)

type Player struct {
	Hand  []deck.Card
	Point int
}

type Game struct {
	Deck      deck.Deck
	TrumpCard deck.Card
	Player1   Player
	Player2   Player
}

func NewGame() *Game {
	d := deck.NewDeck()
	d.Shuffle()
	var p1_hand []deck.Card
	var p2_hand []deck.Card

	player1 := Player{Hand: p1_hand, Point: 0}
	player2 := Player{Hand: p2_hand, Point: 0}

	player1.Hand = append(player1.Hand, d.DrawOpening())
	player2.Hand = append(player2.Hand, d.DrawOpening())

	player1.Hand = append(player1.Hand, d.Draw(4)...)
	player2.Hand = append(player2.Hand, d.Draw(4)...)

	return &Game{
		Deck:      d,
		TrumpCard: deck.GetTrumpCard(player1.Hand[0], player2.Hand[0]),
		Player1:   player1,
		Player2:   player2,
	}
}
