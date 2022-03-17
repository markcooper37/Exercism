package poker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Number string
	Suit   rune
}

type Hand struct {
	Cards    []Card
	HandRank HandRank
}

type Hands []Hand

type HandRank int

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
	fmt.Println(handsSplit)
	panic("Please implement the BestHand function")
}

// Check if there is a pair - if there is, it must be one pair,
// two pair, three of a kind, four of a kind or full house
// Otherwise check for straight and/or flush. The four remaining
// hand types stem from this

func ProcessCards(hands []string) ([][]Card, error) {
	handsSplit := [][]Card{}
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
			number := string(runeCard[:len(runeCard)-1])
			if number != "J" && number != "Q" && number != "K" && number != "A" {
				numberInt, err := strconv.Atoi(number)
				if err != nil {
					return nil, err
				}
				if numberInt < 2 || numberInt > 10 {
					return nil, errors.New("card number must be between 2 and 10")
				}
			}
			handSplit = append(handSplit, Card{number, suit})
		}

		handsSplit = append(handsSplit, handSplit)
	}
	return handsSplit, nil
}
