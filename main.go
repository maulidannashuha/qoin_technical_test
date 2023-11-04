package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func stillPlay(players [][]int) bool {
	stillHaveDice := 0
	for i := range players {
		if len(players[i]) > 0 {
			stillHaveDice++
		}
	}

	return stillHaveDice > 1;
}

func removeItem(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func printDice(dices []int) string{
	diceString := ""

	for i := range dices {
		if i != 0 {
			diceString += " "
		}

		diceString += strconv.Itoa(dices[i])

		if(i < len(dices) - 1) {
			diceString += ","
		}
	}

	return diceString
}

func main() {
	var playerCount int
	fmt.Print("Pemain: ")
    _, err := fmt.Scanf("%d", &playerCount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var diceCount int
	fmt.Print("Dadu: ")
    _, err = fmt.Scanf("%d", &diceCount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	var players = make([][]int, playerCount)
	var playerScores = make([]int, playerCount)
	for i := range players {
		players[i] = make([]int, diceCount)
		playerScores[i] = 0
	}

	fmt.Println("==================")
	for z:=0 ; stillPlay(players); z++ {
		fmt.Println("Giliran", (z+1), "lempar dadu:")

		var appendToPlayers = make([][]int, playerCount)
		for i := range players {
			for j := range players[i] {
				players[i][j] = rand.Intn(6) + 1
			}

			fmt.Printf("\tPemain #%d (%d): %s \n", i+1, playerScores[i], printDice(players[i]))
			for j:=0; j<len(players[i]); j++  {
				if players[i][j] == 6 {
					playerScores[i]++

					if j < 0 || j >= len(players[i]) {
						players[i] = []int{}
					}else{
						players[i] = removeItem(players[i], j)
					}
					j--
				}else if players[i][j] == 1 {
					nextPlayer := i+1
					if nextPlayer == playerCount{
						nextPlayer = 0
					}
					
					appendToPlayers[nextPlayer] = append(appendToPlayers[nextPlayer], 1)
					players[i] = removeItem(players[i], j)
				}
			}
		}

		fmt.Println("Setelah evaluasi:")
		for i:= range players{
			for j:= range appendToPlayers[i]{
				players[i] = append(players[i], appendToPlayers[i][j])
			}

			fmt.Printf("\tPemain #%d (%d): %s \n", i+1, playerScores[i], printDice(players[i]))
		}
		fmt.Println("==================")
	}
}