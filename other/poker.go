package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit rune

const (
	club    Suit = 'c'
	diamond Suit = 'd'
	heart   Suit = 'h'
	spade   Suit = 's'
	none    Suit = '0'
)

type Card struct {
	s   Suit
	pip int
}

func (c Card) Print() {
	fmt.Printf("%d%s, ", c.pip, string(c.s))
}

func Print(l []Card) {
	for _, card := range l {
		card.Print()
	}
	fmt.Println()
}

func shuffle(d *[52]Card) {
	for i := range d {
		j := rand.Intn(52)
		d[i], d[j] = d[j], d[i]
	}
}

func hasFlush(hand [7]Card) Suit {
	var c, d, h, s int8
	for _, card := range hand {
		switch card.s {
		case club:
			c++
		case diamond:
			d++
		case heart:
			h++
		case spade:
			s++
		}
	}

	if c >= 5 {
		return club
	} else if d >= 5 {
		return diamond
	} else if h >= 5 {
		return heart
	} else if s >= 5 {
		return spade
	} else {
		// false case
		return none
	}
}

func getFlush(hand [7]Card, flushSuit Suit) []Card {
	var flush []Card
	for _, card := range hand {
		if card.s == flushSuit {
			flush = append(flush, card)
		}
	}
	if len(flush) < 5 {
		panic("Bad Flush")
	}
	return flush
}

// Sort the hand and check if there's at least 5 `Card.pip` in a row
func isStraight(hand []Card) bool {
	sort.Slice(hand[:], func(i, j int) bool {
		return hand[i].pip < hand[j].pip
	})

	for i := range len(hand) - 4 {
		if hand[i+4].pip-hand[i].pip == 4 {
			return true
		}
	}
	return false
}

func main() {
	var deck [52]Card
	for i := range 52 {
		deck[i].pip = i%13 + 1

		switch i % 4 {
		case 0:
			deck[i].s = club
		case 1:
			deck[i].s = heart
		case 2:
			deck[i].s = diamond
		case 3:
			deck[i].s = spade
		}
	}

	totalDeals := 1_000_000
	var found int
	for range totalDeals {
		shuffle(&deck)

		// we can deal 7 hands without reshuffling
		for i := range 7 {
			var hand [7]Card
			for j := range 7 {
				hand[j] = deck[j+i*7]
			}

			flushSuit := hasFlush(hand)
			if flushSuit != none {
				cards := getFlush(hand, flushSuit)
				if isStraight(cards) {
					found++
				}
			}
		}
	}

	// objective ~ 0.027851 %
	fmt.Println("Found", found, "in", totalDeals*7)
	fmt.Println(float64(found)/float64(totalDeals*7)*100, "%")
}
