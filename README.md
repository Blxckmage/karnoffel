# Karnoffel

A card game implementation in Go where you play against an AI opponent.

## History

Karnöffel is one of the oldest European card games with continuous play down to the present day. It likely originated in the upper-German region in the early 15th century, with the first recorded mention in a 1426 municipal ordinance in Nördlingen, Bavaria. The game was notably played by Landsknechts (German mercenaries) and is depicted in the Herrenberg Altarpiece.

The name "karnöffeln" historically meant "to cudgel or thrash" — a fitting name for this competitive trick-taking game. Karnöffel is considered a possible precursor to the trump suit concept found in Tarot and modern card games.

This implementation is a simplified 2-player digital version of the traditional 4-player game.

## Game Rules

### Setup

- Each hand, both players receive:
  - 1 **opening card** (determines the trump suit for that hand)
  - 4 **playable cards** (used in the 4 tricks)

### Trump Card

- The opening card with the **lowest rank** becomes trump
  - Rank order: 2 (lowest) → 3 → 4 → ... → Queen → King (highest)
- If both opening cards have the same rank, the one with the **higher suit** wins
  - Suit order: Clubs → Diamonds → Hearts → Spades (highest)

### Tricks (4 per hand)

1. Players alternate playing one card per trick
2. **Highest card wins the trick:**
   - Trump card beats any non-trump
   - Within the same suit, higher rank wins
   - If suits don't match and neither is trump, higher suit wins
3. Winner of each trick earns 1 point

### Scoring

- After all 4 tricks, whoever won more tricks gets **4 game points**
- First player to reach **12 game points** wins the game

## How to Run

```bash
go run ./cmd/main.go
```

You'll be prompted to select cards by number. AI plays automatically.
