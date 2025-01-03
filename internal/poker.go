package internal

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Define card ranks and suits with emojis
var RANKS = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var SUITS = []string{"♥", "♦", "♣", "♠"}
var DECKS = map[string]string{
	"2♥": "🂢", "3♥": "🂣", "4♥": "🂤", "5♥": "🂥", "6♥": "🂦", "7♥": "🂧",
	"8♥": "🂨", "9♥": "🂩", "10♥": "🂪", "J♥": "🂫", "Q♥": "🂭", "K♥": "🂮", "A♥": "🂡",
	"2♦": "🃂", "3♦": "🃃", "4♦": "🃄", "5♦": "🃅", "6♦": "🃆", "7♦": "🃇",
	"8♦": "🃈", "9♦": "🃉", "10♦": "🃊", "J♦": "🃋", "Q♦": "🃍", "K♦": "🃎", "A♦": "🃁",
	"2♣": "🃒", "3♣": "🃓", "4♣": "🃔", "5♣": "🃕", "6♣": "🃖", "7♣": "🃗",
	"8♣": "🃘", "9♣": "🃙", "10♣": "🃚", "J♣": "🃛", "Q♣": "🃝", "K♣": "🃞", "A♣": "🃑",
	"2♠": "🂢", "3♠": "🂣", "4♠": "🂤", "5♠": "🂥", "6♠": "🂦", "7♠": "🂧",
	"8♠": "🂨", "9♠": "🂩", "10♠": "🂪", "J♠": "🂫", "Q♠": "🂭", "K♠": "🂮", "A♠": "🂡",
}

type Card struct {
	Rank string
	Suit string
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	cards := []Card{}
	for _, rank := range RANKS {
		for _, suit := range SUITS {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Deal(num int) []Card {
	dealtCards := d.Cards[:num]
	d.Cards = d.Cards[num:]
	return dealtCards
}

type PokerHand struct {
	Cards []Card
}

func NewPokerHand(cards []Card) *PokerHand {
	return &PokerHand{Cards: cards}
}

func (ph *PokerHand) String() string {
	var cards []string
	for _, card := range ph.Cards {
		cards = append(cards, wrapColor(card))
	}
	return fmt.Sprintf("\033[97m\033[107m%s \033[0m\033[97m", strings.Join(cards, " "))
}

func wrapColor(card Card) string {
	cardStr := card.String()
	if card.Suit == "♦" || card.Suit == "♥" {
		return fmt.Sprintf("\033[91m%s", DECKS[cardStr])
	}
	return fmt.Sprintf("\033[30m%s", DECKS[cardStr])
}

func (ph *PokerHand) EvaluateHand() string {
	ranks := make(map[string]int)
	suits := make(map[string]int)
	var rankIndices []int

	for _, card := range ph.Cards {
		ranks[card.Rank]++
		suits[card.Suit]++
		for i, rank := range RANKS {
			if rank == card.Rank {
				rankIndices = append(rankIndices, i)
			}
		}
	}

	sort.Ints(rankIndices)

	isStraight := true
	for i := 1; i < len(rankIndices); i++ {
		if rankIndices[i] != rankIndices[i-1]+1 {
			isStraight = false
			break
		}
	}

	isFlush := len(suits) == 1

	switch {
	case isStraight && isFlush:
		return "Straight Flush"
	case containsValue(ranks, 4):
		return "Four of a Kind"
	case containsValue(ranks, 3) && containsValue(ranks, 2):
		return "Full House"
	case isFlush:
		return "Flush"
	case isStraight:
		return "Straight"
	case containsValue(ranks, 3):
		return "Three of a Kind"
	case countValue(ranks, 2) == 2:
		return "Two Pair"
	case containsValue(ranks, 2):
		return "One Pair"
	default:
		return "High Card"
	}
}

func containsValue(m map[string]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func countValue(m map[string]int, value int) int {
	count := 0
	for _, v := range m {
		if v == value {
			count++
		}
	}
	return count
}

type PokerGame struct {
	Deck       *Deck
	PlayerHand *PokerHand
}

func NewPokerGame() *PokerGame {
	return &PokerGame{Deck: NewDeck()}
}

func (pg *PokerGame) StartGame() {
	pg.Deck.Shuffle()
	pg.PlayerHand = NewPokerHand(pg.Deck.Deal(5))
}

func (pg *PokerGame) ShowScore() {
	fmt.Println(pg.PlayerHand, pg.PlayerHand.EvaluateHand()+"!")
}

func PlayPoker() {
	game := NewPokerGame()
	game.StartGame()
	game.ShowScore()
}
