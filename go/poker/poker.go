package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Number int
	Suit   rune
}

type Hand struct {
	Cards    []Card
	HandRank HandRank
}

type HandRank int

var PictureToNumber = map[string]int{"J": 11, "Q": 12, "K": 13, "A": 14}

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func BestHand(hands []string) ([]string, error) {
	handsSplit, err := ProcessCards(hands)
	if err != nil {
		return nil, err
	}
	for _, hand := range handsSplit {
		hand.FindRank()
	}

	panic("Please implement the BestHand function")
}

// Check if there is a pair - if there is, it must be one pair,
// two pair, three of a kind, four of a kind or full house
// Otherwise check for straight and/or flush. The four remaining
// hand types stem from this

func ProcessCards(hands []string) ([]Hand, error) {
	handsSplit := []Hand{}
	for _, hand := range hands {
		separateCards := strings.Split(hand, " ")
		if len(separateCards) != 5 {
			return nil, errors.New("hand contains wrong number of cards")
		}
		handSplit := []Card{}
		for _, card := range separateCards {
			runeCard := []rune(card)
			suit := runeCard[len(runeCard)-1]
			if suit != '♡' && suit != '♢' && suit != '♤' && suit != '♧' {
				return nil, errors.New("no appropriate suit")
			}
			numberString := string(runeCard[:len(runeCard)-1])
			number := 0
			if numberString == "J" || numberString == "Q" || numberString == "K" || numberString == "A" {
				number = PictureToNumber[numberString]
			} else {
				numberInt, err := strconv.Atoi(numberString)
				if err != nil {
					return nil, err
				}
				if numberInt < 2 || numberInt > 10 {
					return nil, errors.New("card number must be between 2 and 10")
				}
				number = numberInt
			}
			handSplit = append(handSplit, Card{number, suit})
		}
		sort.Slice(handSplit, func(i, j int) bool { return handSplit[i].Number < handSplit[j].Number })
		handsSplit = append(handsSplit, Hand{Cards: handSplit})
	}
	return handsSplit, nil
}

func (h *Hand) FindRank() {
	if h.PairExists() {
		if h.IsFourOfAKind() {
			h.HandRank = FourOfAKind
		} else if h.IsFullHouse() {
			h.HandRank = FullHouse
		} else if h.IsThreeOfAKind() {
			h.HandRank = ThreeOfAKind
		} else if h.IsTwoPair() {
			h.HandRank = TwoPair
		} else {
			h.HandRank = OnePair
		}
	} else {
		if h.IsFlush() && h.IsStraight() {
			h.HandRank = StraightFlush
		} else if h.IsFlush() {
			h.HandRank = Flush
		} else if h.IsStraight() {
			h.HandRank = Straight
		} else {
			h.HandRank = HighCard
		}
	}
}

func (h *Hand) PairExists() bool {
	return h.Cards[0].Number == h.Cards[1].Number &&
		h.Cards[1].Number == h.Cards[2].Number &&
		h.Cards[2].Number == h.Cards[3].Number &&
		h.Cards[3].Number == h.Cards[4].Number
}

func (h *Hand) IsStraight() bool {
	return (h.Cards[0].Number+1 == h.Cards[1].Number &&
		h.Cards[1].Number+1 == h.Cards[2].Number &&
		h.Cards[2].Number+1 == h.Cards[3].Number &&
		h.Cards[3].Number+1 == h.Cards[4].Number) ||
		(h.Cards[0].Number+1 == h.Cards[1].Number &&
			h.Cards[1].Number+1 == h.Cards[2].Number &&
			h.Cards[2].Number+1 == h.Cards[3].Number &&
			(h.Cards[4].Number%13)+1 == h.Cards[0].Number)
}

func (h *Hand) IsFlush() bool {
	return h.Cards[0].Suit == h.Cards[1].Suit &&
		h.Cards[1].Suit == h.Cards[2].Suit &&
		h.Cards[2].Suit == h.Cards[3].Suit &&
		h.Cards[3].Suit == h.Cards[4].Suit
}

func (h *Hand) IsFourOfAKind() bool {
	return (h.Cards[0].Number == h.Cards[1].Number &&
		h.Cards[1].Number+1 == h.Cards[2].Number &&
		h.Cards[2].Number+1 == h.Cards[3].Number) ||
		(h.Cards[1].Number+1 == h.Cards[2].Number &&
			h.Cards[2].Number+1 == h.Cards[3].Number &&
			h.Cards[3].Number+1 == h.Cards[4].Number)
}

func (h *Hand) IsFullHouse() bool {
	return (h.Cards[0].Number == h.Cards[1].Number &&
		h.Cards[2].Number == h.Cards[3].Number &&
		h.Cards[3].Number == h.Cards[4].Number) ||
		(h.Cards[0].Number == h.Cards[1].Number &&
			h.Cards[1].Number == h.Cards[2].Number &&
			h.Cards[3].Number == h.Cards[4].Number)
}

func (h *Hand) IsThreeOfAKind() bool {
	return (h.Cards[0].Number == h.Cards[1].Number &&
		h.Cards[1].Number == h.Cards[2].Number) ||
		(h.Cards[1].Number == h.Cards[2].Number &&
			h.Cards[2].Number == h.Cards[3].Number) ||
		(h.Cards[2].Number == h.Cards[3].Number &&
			h.Cards[3].Number == h.Cards[4].Number)
}

func (h *Hand) IsTwoPair() bool {
	return (h.Cards[0].Number == h.Cards[1].Number &&
		h.Cards[2].Number == h.Cards[3].Number) ||
		(h.Cards[0].Number == h.Cards[1].Number &&
			h.Cards[3].Number == h.Cards[4].Number) ||
		(h.Cards[1].Number == h.Cards[2].Number &&
			h.Cards[3].Number == h.Cards[4].Number)
}
