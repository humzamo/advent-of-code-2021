package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	bits = 12
)

func main() {
	var inputFileName = "Input.txt"
	var list = loadFileToList(inputFileName)

	partOneAnswer := partOne(list)
	fmt.Println("The answer to part one is:", partOneAnswer)
	partTwoAnswer := partTwo(list)
	fmt.Println("The answer to part two is:", partTwoAnswer)
}

func loadFileToList(inputFileName string) [][]int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var list [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var intSlice []int
		splitString := strings.Split(scanner.Text(), "")
		for _, s := range splitString {
			binary, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error!")
			}
			intSlice = append(intSlice, binary)
		}
		list = append(list, intSlice)
	}
	defer file.Close()
	return list
}

func partOne(list [][]int) int {
	var numberOfOnes [bits]int
	for _, s := range list {
		for i, v := range s {
			numberOfOnes[i] += v
		}
	}

	var gammaRate, epsilonRate string
	midpointOfRows := len(list) / 2
	for _, s := range numberOfOnes {
		if s > midpointOfRows {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	gammaRateInt := convertBinaryStringToInt(gammaRate)
	epsilonRateInt := convertBinaryStringToInt(epsilonRate)

	return gammaRateInt * epsilonRateInt
}

func partTwo(list [][]int) int {
	oxygen := calclulateOxygen(list)
	co2 := calclulateCo2(list)

	var oxygenString, co2String string
	for i := 0; i < bits; i++ {
		oxygenString += strconv.Itoa(oxygen[i])
		co2String += strconv.Itoa(co2[i])
	}

	oxygenRateInt := convertBinaryStringToInt(oxygenString)
	co2RateInt := convertBinaryStringToInt(co2String)

	return oxygenRateInt * co2RateInt

}

func calclulateOxygen(list [][]int) []int {
	for i := 0; i < bits; i++ {
		if len(list) == 1 {
			return list[0]
		}

		numberOfZeros, numberOfOnes := calculateZerosAndOnes(list, i)

		if numberOfOnes >= numberOfZeros {
			list = updateList(list, 1, i)
		} else {
			list = updateList(list, 0, i)
		}
	}
	return list[0]
}

func calclulateCo2(list [][]int) []int {
	for i := 0; i < bits; i++ {
		if len(list) == 1 {
			return list[0]
		}

		numberOfZeros, numberOfOnes := calculateZerosAndOnes(list, i)

		if numberOfOnes < numberOfZeros {
			list = updateList(list, 1, i)
		} else {
			list = updateList(list, 0, i)
		}
	}
	return list[0]
}

func convertBinaryStringToInt(binaryString string) int {
	number, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		fmt.Println("Error!")
	}
	return int(number)
}

// Calculated the number of zeros and ones at a specified index in the list.
func calculateZerosAndOnes(list [][]int, i int) (int, int) {
	var numberOfOnes int
	for _, s := range list {
		numberOfOnes += s[i]
	}
	numberOfZeros := len(list) - numberOfOnes
	return numberOfZeros, numberOfOnes
}

// Keeps items in the list and at a specified index that start with a specified number.
func updateList(list [][]int, number int, index int) [][]int {
	var updatedlist [][]int
	for _, s := range list {
		if s[index] == number {
			updatedlist = append(updatedlist, s)
		}
	}
	return updatedlist
}
