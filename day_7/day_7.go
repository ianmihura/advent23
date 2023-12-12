package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ORDER []string = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

type Hands struct {
	hands []string
	bids  []int
}

// TODO pointer receiver error
type Playable interface {
	get_score() int
	sort() Hands
	concat(Hands) Hands
	append(string, int)
}

// Result of battle between two Hands
// -1: first one wins (left input)
// 0: tie
// 1: second one wins (right input)
func get_highest_order(hand_left, hand_right rune) int {
	for o := range ORDER {
		if ORDER[o] == string(hand_left) && ORDER[o] == string(hand_right) {
			return 0
		} else if ORDER[o] == string(hand_left) {
			return -1
		} else if ORDER[o] == string(hand_right) {
			return 1
		}
	}
	// default case should never happen
	return 0
}

func get_longest_group(groups []string) int {
	max_ids := 0
	for i := range groups {
		if len(groups[max_ids]) < len(groups[i]) {
			max_ids = i
		}
	}
	return max_ids
}

// score is trailing sum of len(group)^2
func get_hand_score(hand string) int {
	groups := []string{string("")}
	j_count := 0

	// Separate hand cards into groups of same card
	for i := 0; i < len(hand); i++ {
		if string(hand[i]) == "J" {
			j_count++
			continue
		}
		match := false
		for j := range groups {
			if len(groups[j]) > 0 && groups[j][0] == hand[i] {
				groups[j] += string(hand[i])
				match = true
			}
		}
		if !match {
			groups = append(groups, string(hand[i]))
		}
	}
	if j_count > 0 {
		max_ids := get_longest_group(groups)
		for i := 0; i < j_count; i++ {
			groups[max_ids] += "J"
		}
	}

	score := 0
	for i := range groups {
		score += len(groups[i]) * len(groups[i])
	}

	return score
}

// Result of battle between two Hands
// -1: first one wins (left input)
// 0: tie
// 1: second one wins (right input)
func get_battle_result(hand_left, hand_right string) int {
	left_score := get_hand_score(hand_left)
	right_score := get_hand_score(hand_right)

	if hand_left == hand_right {
		return 0
	} else if left_score == right_score {
		for i := range hand_left {
			order := get_highest_order(rune(hand_left[i]), rune(hand_right[i]))
			if order != 0 {
				return order
			}
		}
		// default case should never happen
		return 0
	} else if left_score > right_score {
		return -1
	} else {
		return 1
	}
}

// Calculates score for all hands (relative to their order)
// and returns their trailing addition
func (hand_list Hands) get_score() int {
	trailing := 0
	for i := 1; i <= len(hand_list.hands); i++ {
		trailing += hand_list.bids[i-1] * i
	}
	return trailing
}

func (hand_list Hands) sort() Hands {
	var low, same, high Hands

	if len(hand_list.hands) == 0 {
		return hand_list
	}

	pivot := hand_list.hands[0]

	for i, item := range hand_list.hands {
		battle_result := get_battle_result(item, pivot)

		if battle_result == 1 {
			low.append(item, hand_list.bids[i])
		} else if battle_result == 0 {
			same.append(item, hand_list.bids[i])
		} else if battle_result == -1 {
			high.append(item, hand_list.bids[i])
		}
	}

	return low.sort().concat(same.concat(high.sort()))
}

func (left_hand_list Hands) concat(right_hand_list Hands) Hands {
	for i := 0; i < len(right_hand_list.hands); i++ {
		left_hand_list.append(right_hand_list.hands[i], right_hand_list.bids[i])
	}
	return left_hand_list
}

func (hand_list *Hands) append(hand string, bid int) {
	hand_list.hands = append(hand_list.hands, hand)
	hand_list.bids = append(hand_list.bids, bid)
}

func run(hand_list Hands) int {
	hand_list = hand_list.sort()
	return hand_list.get_score()
}

func format(byte_data []byte) Hands {
	data := strings.Split(string(byte_data), "\n")
	hand_list := Hands{}

	for i := 0; i < len(data); i++ {
		d := strings.Split(data[i], " ")
		bid, _ := strconv.Atoi(d[1])

		hand_list.append(d[0], bid)
	}
	return hand_list
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	hand_list := format(data)
	answer := run(hand_list)
	fmt.Println(answer)
}
