package main

import (
	"fmt"
	"math/rand"
	"time"
)

func stillPlay(players [][]int) bool {
	return true;
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

	for z:=0 ; stillPlay(players); z++ {
		fmt.Println("Giliran", (z+1), "lempar dadu:")

		for i := range players {
			for j := range players[i] {
				players[i][j] = rand.Intn(6) + 1
			}

			fmt.Printf("\tPemain #%d (%d)\n", i+1, playerScores[i])

			for j := range players[i] {
				if players[i][j] == 6 {
					
				}
			}
		}
	}

	// while stillPlay(players) {
	// 	fmt.Println("==================")
	// }

}