package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	coin_count := 0
	red := 0.0
	black := 0.0
	AceSpade := 0.0
	diceResults := [6]float64{}
	TestSize := 1000.0
	var seed int64
	fmt.Println("How many tests wouldyou like to simulate? ")
	fmt.Scanln(&TestSize)
	fmt.Println("please type an integer to use as a seed? type '0' for a random seed")
	fmt.Scanln(&seed)
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	rand.Seed(seed)
	for i := 0; i < int(TestSize); i++ {
		coin_count += flipcoin()
		DiceRoll(&diceResults)
		if CardColor() {
			red++
		} else {
			black++
		}
		if DrawSpecificCard() {
			AceSpade++
		}
	}
	headPercent := float64(coin_count)
	headPercent = headPercent / TestSize
	//Print the simulated Coin flips

	PrintHeader("Coin Flip")
	fmt.Printf(" heads: %d\n tails: %d\n percentage of heads %.2f%%\n", coin_count, int(TestSize)-coin_count, 100*headPercent)
	PrintLine()
	fmt.Printf("expected results\nheads: %.0f\ntails: %0.f\npercentage of heads 50%%\n", 0.5*TestSize, 0.5*TestSize)
	PrintLine()
	PrintSpace(2)
	//Print the simulated Dice Rolls
	PrintHeader("Dice Roll")
	fmt.Printf("1's: %.0f\n2's: %.0f\n", diceResults[0], diceResults[1])
	fmt.Printf("3's: %.0f\n4's: %.0f\n", diceResults[3], diceResults[4])
	fmt.Printf("5's: %.0f\n6's: %.0f\n", diceResults[4], diceResults[5])
	fmt.Printf("1's percentage : %.0f\n2's percentage : %.0f\n", diceResults[0]/TestSize*100.0, diceResults[1]/TestSize*100.0)
	fmt.Printf("3's percentage : %.0f\n4's percentage : %.0f\n", diceResults[3]/TestSize*100.0, diceResults[4]/TestSize*100.0)
	fmt.Printf("5's percentage : %.0f\n6's percentage : %.0f\n", diceResults[4]/TestSize*100.0, diceResults[5]/TestSize*100.0)
	PrintLine()
	fmt.Println("Expected Results")
	expected := TestSize * 1.0 / 6.0
	fmt.Printf("1's: %.3f\n2's: %.3f\n3's: %.3f\n4's: %.3f\n5's: %.3f\n6's: %.3f\n", expected, expected, expected, expected, expected, expected)
	expected = 1.0 / 6.0 * 100.00
	fmt.Printf("1's percentage : %.3f%%\n2's percentage : %.3f%%\n3's percentage : %.3f%%\n4's percentage : %.3f%%\n5's percentage : %.3f%%\n6's percentage : %.3f%%\n",
		expected, expected, expected, expected, expected, expected)
	PrintLine()
	PrintSpace(2)
	//Print the simulated Card Color Drawing
	PrintHeader("Card Color Draw")
	fmt.Printf("Reds: %.0f\nBlacks: %.0f\nPercentage Red:%.2f%%\nPercentage Black: %.2f%%\n", red, black, 100*red/TestSize, 100*black/TestSize)
	PrintLine()
	fmt.Println("Expected Results")
	fmt.Printf("Reds: %.0f\nBlacks: %.0f\nPercentage Red:%.2f%%\nPercentage Black: %.2f%%\n", 500.0, 500.0, 50.0, 50.0)
	PrintLine()
	//simulate drawing a ace of spades
	PrintSpace(2)
	PrintHeader("Draw an ace of Spades")
	fmt.Printf("Ace of Spades Drawn: %.0f\nother cards drawn %0.f\n", AceSpade, TestSize-AceSpade)
	fmt.Printf("Ace of Spades Percentage : %.3f%%\n", AceSpade/TestSize*100)
	PrintLine()
	fmt.Println("Expected Results")
	fmt.Printf("ace of spades Drawn: %.2f\nother cards Drawn %.2f\n", 1.0/52.0*TestSize, 51.0/52.0*TestSize)
	fmt.Printf("Ace of Spades Percentage : %.3f%%\n", 1.0/52.0*100)
	PrintLine()
	PrintSpace(5)

}

//random generators
func flipcoin() int {
	//a 0 is a tails, and 1 is a heads
	return rand.Intn(2)
}

func CardColor() bool {
	//a true result is red, a false is black
	if rand.Intn(53) <= 27 {
		return true
	} else {
		return false
	}

}

func DrawSpecificCard() bool {
	if rand.Intn(52) == 0 {
		return true
	}
	return false
}

func DiceRoll(storage *[6]float64) {

	storage[rand.Intn(6)]++
}

///Below Are just useful little functions to make printing easier
func PrintLine() {
	fmt.Println("-------------------------------\n")
}
func PrintHeader(header string) {
	fmt.Printf("=========%s========\n", header)
}
func PrintSpace(size int) {
	for i := 0; i < size; i++ {
		fmt.Println("")
	}
}
