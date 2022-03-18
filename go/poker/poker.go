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
	if len(hands) == 0 {
		return []string{}, nil
	}
	handsSplit, err := ProcessCards(hands)
	if err != nil {
		return nil, err
	}
	for _, hand := range handsSplit {
		hand.FindRank()
	}
	bestHandsIndices := []int{0}
	for i := 1; i < len(hands); i++ {
		if handsSplit[i].HandRank < handsSplit[bestHandsIndices[0]].HandRank {
			continue
		} else if handsSplit[i].HandRank > handsSplit[bestHandsIndices[0]].HandRank {
			bestHandsIndices = []int{i}
		} else {
			compareHands := CompareEqualRank(handsSplit[i], handsSplit[bestHandsIndices[0]])
			if compareHands == "first" {
				bestHandsIndices = []int{i}
			} else if compareHands == "second" {
				continue
			} else if compareHands == "equal" {
				bestHandsIndices = append(bestHandsIndices, i)
			}
		}
	}
	bestHands := []string{}
	for _, index := range bestHandsIndices {
		bestHands = append(bestHands, hands[index])
	}
	return bestHands, nil
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
		newHand := Hand{Cards: handSplit}
		newHand.FindRank()
		handsSplit = append(handsSplit, newHand)
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
	return h.Cards[0].Number == h.Cards[1].Number ||
		h.Cards[1].Number == h.Cards[2].Number ||
		h.Cards[2].Number == h.Cards[3].Number ||
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
		h.Cards[1].Number == h.Cards[2].Number &&
		h.Cards[2].Number == h.Cards[3].Number) ||
		(h.Cards[1].Number == h.Cards[2].Number &&
			h.Cards[2].Number == h.Cards[3].Number &&
			h.Cards[3].Number == h.Cards[4].Number)
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

func CompareEqualRank(first, second Hand) string {
	if first.HandRank == HighCard ||
		first.HandRank == Flush {
		return CompareHighCards(first, second)
	} else if first.HandRank == Straight ||
		first.HandRank == StraightFlush {
		if (first.Cards[0].Number == 2 &&
			first.Cards[4].Number == 14) &&
			(second.Cards[0].Number == 2 &&
			second.Cards[4].Number != 14) {
			return "second"
		} else if (first.Cards[0].Number == 2 &&
			first.Cards[4].Number != 14) && 
			(second.Cards[0].Number == 2 && 
			second.Cards[4].Number == 14) {
			return "first"
		} else {
			return CompareHighCards(first, second)
		}
	} else {
		firstEqualCards := first.FindEqualCards()
		secondEqualCards := second.FindEqualCards()
		for index, number := range firstEqualCards {
			if number > secondEqualCards[index] {
				return "first"
			} else if number < secondEqualCards[index] {
				return "second"
			} else {
				return CompareHighCards(first, second)
			}
		}
	}
	return ""
}

func CompareHighCards(first, second Hand) string {
	for i := 4; i >= 0; i-- {
		if first.Cards[i].Number > second.Cards[i].Number {
			return "first"
		} else if first.Cards[i].Number < second.Cards[i].Number {
			return "second"
		}
	}
	return "equal"
}

// Finds which sets of cards within a hand are equal
// With a full house, it returns the three of a kind at the front
// With two pair, it returns the largest pair at the front
func (h *Hand) FindEqualCards() []int {
	if h.HandRank == FourOfAKind || h.HandRank == ThreeOfAKind {
		return []int{h.Cards[2].Number}
	} else if h.HandRank == FullHouse {
		if h.Cards[0].Number == h.Cards[1].Number &&
			h.Cards[1].Number == h.Cards[2].Number {
			return []int{h.Cards[0].Number, h.Cards[4].Number}
		} else {
			return []int{h.Cards[4].Number, h.Cards[0].Number}
		}
	} else if h.HandRank == TwoPair {
		return []int{h.Cards[3].Number, h.Cards[1].Number}
	} else if h.HandRank == OnePair {
		if h.Cards[0].Number == h.Cards[1].Number ||
			h.Cards[1].Number == h.Cards[2].Number {
			return []int{h.Cards[1].Number}
		} else {
			return []int{h.Cards[3].Number}
		}
	}
	return []int{}
}
