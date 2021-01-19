package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var (
	// 0 indicates Guy de Chauliac
	DiceResults = map[int]int{
		2:  0,
		3:  3,
		4:  5,
		5:  7,
		6:  10,
		7:  12,
		8:  10,
		9:  4,
		10: 8,
		11: 8,
		12: 0,
	}
)

func Conduct(trials int) int {
	var totalDeaths int
	file := CreateFileWithRandomID()
	defer file.Close()
	for i := 1; i <= trials; i++ {
		dice, total := RollDice()
		deaths := DiceResults[total]
		if deaths == 0 {
			arr, _ := RollDice()
			if arr[0] == 6 {
				deaths = -15
			} else {
				deaths = -arr[0]
			}
		}
		str := strconv.Itoa(i) + "\t\t" + strconv.Itoa(dice[0]) + "--" + strconv.Itoa(dice[1]) + "\t\t" + strconv.Itoa(deaths) + "\n"
		fmt.Print(str)
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println("Could not write to file", err)
		}
		totalDeaths += deaths
	}
	_, err := file.WriteString("\nTotal Deaths: " + strconv.Itoa(totalDeaths))
	if err != nil {
		fmt.Println("Could not write total deaths to file", err)
	}
	return totalDeaths
}

func CreateFileWithRandomID() *os.File {
	f, err := os.Create("blackdeath-" + generateRandomID(6) + ".txt")
	if err != nil {
		log.Fatal("Could not create file", err)
	}
	return f
}

func generateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RollDice() ([2]int, int) {
	var arr [2]int
	min := 1
	max := 6
	arr[0] = rand.Intn(max-min+1) + min
	arr[1] = rand.Intn(max-min+1) + min
	return arr, arr[0] + arr[1]
}
