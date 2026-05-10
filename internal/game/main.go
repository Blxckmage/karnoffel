package game

import (
	"fmt"
	"math/rand"

	"karnoffel/internal/deck"
)

type Player struct {
	Hand        []deck.Card
	GamePoint   int
	RoundPoint  int
	PlayedCards []deck.Card
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

	var p1Hand []deck.Card
	var p2Hand []deck.Card

	player1 := Player{Hand: p1Hand, GamePoint: 0, RoundPoint: 0}
	player2 := Player{Hand: p2Hand, GamePoint: 0, RoundPoint: 0}

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

func (g *Game) DealNewHand() {
	g.Player1.PlayedCards = []deck.Card{}
	g.Player2.PlayedCards = []deck.Card{}
	g.Player1.Hand = []deck.Card{}
	g.Player2.Hand = []deck.Card{}

	if len(g.Deck.Cards) < 10 {
		fmt.Println("Reshuffling deck...")
		g.Deck = deck.NewDeck()
		g.Deck.Shuffle()
	}

	g.Player1.Hand = append(g.Player1.Hand, g.Deck.DrawOpening())
	g.Player2.Hand = append(g.Player2.Hand, g.Deck.DrawOpening())

	g.Player1.Hand = append(g.Player1.Hand, g.Deck.Draw(4)...)
	g.Player2.Hand = append(g.Player2.Hand, g.Deck.Draw(4)...)

	g.TrumpCard = deck.GetTrumpCard(g.Player1.Hand[0], g.Player2.Hand[0])
	fmt.Printf("Trump Card: %s\n", g.TrumpCard.CardToString())
}

func (g *Game) DetermineFirstPlayer() (*Player, *Player) {
	rank1, suit1 := g.Player1.Hand[0].GetCardValue()
	rank2, suit2 := g.Player2.Hand[0].GetCardValue()

	if rank1 > rank2 {
		fmt.Println("Player 1 starts")
		return &g.Player1, &g.Player2
	} else if rank2 > rank1 {
		fmt.Println("Player 2 starts")
		return &g.Player2, &g.Player1
	} else if suit1 > suit2 {
		fmt.Println("Player 1 starts")
		return &g.Player1, &g.Player2
	} else {
		fmt.Println("Player 2 starts")
		return &g.Player2, &g.Player1
	}
}

func (g *Game) PlayRound(p *Player, isAI bool) {
	var choices int
	if isAI {
		randomIndex := rand.Intn(len(p.Hand))
		playedCard := p.Hand[randomIndex]
		p.PlayedCards = append(p.PlayedCards, playedCard)
		p.Hand = append(p.Hand[:randomIndex], p.Hand[randomIndex+1:]...)
		fmt.Printf("AI plays: %s\n", playedCard.CardToString())
	} else {
		for i, card := range p.Hand {
			fmt.Printf("%d: %s\n", i+1, card.CardToString())
		}

		for {
			fmt.Print("Type a number: ")
			_, err := fmt.Scan(&choices)
			if err != nil {
				fmt.Println("Invalid Input")
				continue
			}

			if choices < 1 || choices > len(p.Hand) {
				fmt.Printf("Number must be between %d to %d\n", 1, len(p.Hand))
				continue
			}

			cardIndex := choices - 1
			playedCard := p.Hand[cardIndex]
			p.PlayedCards = append(p.PlayedCards, playedCard)
			p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
			fmt.Println("Card Choices: ", playedCard)

			break
		}
	}
}

func (g *Game) ResolveRound(p1, p2 *Player) {
	player1Card := p1.PlayedCards[len(p1.PlayedCards)-1]
	player2Card := p2.PlayedCards[len(p2.PlayedCards)-1]
	winningCard := deck.CompareCards(player1Card, player2Card, g.TrumpCard)

	fmt.Printf("Player 1 played: %s\n", player1Card.CardToString())
	fmt.Printf("Player 2 played: %s\n", player2Card.CardToString())
	fmt.Printf("Winning card: %s\n", winningCard.CardToString())

	if winningCard == p1.PlayedCards[len(p1.PlayedCards)-1] {
		p1.RoundPoint += 1
		fmt.Println("Player 1 wins this trick!")
	} else {
		p2.RoundPoint += 1
		fmt.Println("Player 2 wins this trick!")
	}
	fmt.Println()
}

func (g *Game) ResolveHand(p1, p2 *Player) {
	fmt.Println("=== Hand Results ===")
	fmt.Printf("Player 1 tricks won: %d\n", p1.RoundPoint)
	fmt.Printf("Player 2 tricks won: %d\n", p2.RoundPoint)

	if p1.RoundPoint > p2.RoundPoint {
		p1.GamePoint += 4
		fmt.Println("Player 1 wins the hand! +4 points")
	} else if p2.RoundPoint > p1.RoundPoint {
		p2.GamePoint += 4
		fmt.Println("Player 2 wins the hand! +4 points")
	} else {
		fmt.Println("Hand is tied!")
	}

	fmt.Printf("Game score - Player 1: %d, Player 2: %d\n", p1.GamePoint, p2.GamePoint)
	fmt.Println()

	p1.RoundPoint = 0
	p2.RoundPoint = 0
}
