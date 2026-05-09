package deck

import (
	"math/rand"
)

type Card struct {
	Suit string
	Rank string
}

type Deck struct {
	Cards []Card
}

var cardValues = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5,
	"6": 6, "7": 7, "8": 8, "9": 9,
	"10": 10, "Jack": 11, "Queen": 12,
	"King": 13,
}

var suitValues = map[string]int{
	"Clubs": 1, "Diamonds": 2, "Hearts": 3, "Spades": 4,
}

func NewDeck() Deck {
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	var d Deck
	for _, suit := range suits {
		for _, rank := range ranks {
			d.Cards = append(d.Cards, Card{Suit: suit, Rank: rank})
		}
	}

	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i int, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw(n int) []Card {
	drawn := d.Cards[:n]

	d.Cards = d.Cards[n:]

	return drawn
}

func (d *Deck) DrawOpening() Card {
	drawn := d.Cards[0]

	d.Cards = d.Cards[1:]

	return drawn
}

func (c *Card) GetCardValue() (int, int) {
	return cardValues[c.Rank], suitValues[c.Suit]
}

func GetTrumpCard(card1, card2 Card) Card {
	rank1, suit1 := card1.GetCardValue()
	rank2, suit2 := card2.GetCardValue()

	if rank1 > rank2 {
		return card2
	} else if rank2 > rank1 {
		return card1
	}

	if suit1 > suit2 {
		return card2
	}
	return card1
}
