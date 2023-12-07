package main

import (
    "os"
    "fmt"
	"slices"
	"strconv"
    "strings"
)

var CardStrengh = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var CardStrenghJoker = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type HandStrengh int

const (
	HighCard HandStrengh = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h HandStrengh) String() string {
	switch h {
	case FiveOfAKind:
		return "FiveOfAKind"
	case FourOfAKind:
		return "FourOfAKind"
	case FullHouse:
		return "FullHouse"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case TwoPair:
		return "TwoPair"
	case OnePair:
		return "OnePair"
	default:
		return "HighCard"
	}
}

type Hand struct {
	cards []string
	bid int
	strengh HandStrengh	
}

func ScoreHand(cards []string) HandStrengh {
	seen := ""
	var counts []int
	for _, card := range cards {
		index := strings.Index(seen, card)
		if index > -1 {
			counts[index] += 1
		} else {
			seen += card
			counts = append(counts, 1)
		}
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	length := len(counts)
	var val int
	if length != 1 {
		val = counts[0] * 10 + counts[1]
	} else {
		val = 50
	}

	switch val {
	case 50:
		return FiveOfAKind
	case 41:
		return FourOfAKind
	case 32:
		return FullHouse
	case 31:
		return ThreeOfAKind
	case 22:
		return TwoPair
	case 21:
		return OnePair
	default:
		return HighCard
	}
}

func JokerScore(cards string) HandStrengh {
	othercards := "23456789TAQK"
	scores := make([]HandStrengh, len(othercards))

	for i, card := range othercards {
		testCards := strings.Split(strings.Replace(cards,"J",string(card),-1),"")
		scores[i] = ScoreHand(testCards)
	}
	
	return slices.Max(scores)
}

func SortHands(a, b Hand) int {
	if a.strengh > b.strengh { return 1 }
	if a.strengh < b.strengh { return -1 }
	for i, _ := range a.cards {
		aStrengh := CardStrengh[a.cards[i]]
		bStrengh := CardStrengh[b.cards[i]]
		if aStrengh > bStrengh { return 1 }
		if aStrengh < bStrengh { return -1 }
	}

	return 1
}

func SortHandsJoker(a, b Hand) int {
	if a.strengh > b.strengh { return 1 }
	if a.strengh < b.strengh { return -1 }
	for i, _ := range a.cards {
		aStrengh := CardStrenghJoker[a.cards[i]]
		bStrengh := CardStrenghJoker[b.cards[i]]
		if aStrengh > bStrengh { return 1 }
		if aStrengh < bStrengh { return -1 }
	}

	return 1
}

func PartOne(arr *[]string) int  {
    input := *arr
    result := 0
	hands := make([]Hand, len(input))
	
	for i, hand := range input {
		roundData := strings.Fields(hand)
		cards := roundData[0]
		bid, _ := strconv.Atoi(roundData[1])
		cardArr := strings.Split(cards, "")

		newHand := Hand{
			bid: bid,
			cards: cardArr,
			strengh: ScoreHand(cardArr),
		}
		hands[i] = newHand
	}

	slices.SortFunc(hands, SortHands)	
	
	for rank, hand := range hands {
		result += hand.bid * (rank + 1)
	}

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    result := 0
	hands := make([]Hand, len(input))
	
	for i, hand := range input {
		roundData := strings.Fields(hand)
		cards := roundData[0]
		bid, _ := strconv.Atoi(roundData[1])
		cardArr := strings.Split(cards, "")

		newHand := Hand{
			bid: bid,
			cards: cardArr,
			strengh: ScoreHand(cardArr),
		}

		if strings.Contains(cards, "J") {
			newHand.strengh = JokerScore(cards)
		}
		
		hands[i] = newHand
	}

	slices.SortFunc(hands, SortHandsJoker)	
	
	for rank, hand := range hands {
		result += hand.bid * (rank + 1)
	}

    return result
}

func main () {
    inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-1]), "\n")

    partOneResult := PartOne(&input)
    fmt.Printf("Part One Result: %d\n", partOneResult)

    partTwoResult := PartTwo(&input)
    fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
