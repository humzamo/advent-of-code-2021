package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var inputFileName = "Input.txt"

	var numberList = loadFileToListofNumbers(inputFileName)
	partOneAnswer := partOne(numberList)
	fmt.Println("The answer to part one is:", partOneAnswer)

	partTwoAnswer := partTwo(numberList)
	fmt.Println("The answer to part two is:", partTwoAnswer)
}

func loadFileToListofNumbers(inputFileName string) []int {

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var numberList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("The input list contains a value which is not an integer!")
			log.Fatal(err)
		}
		numberList = append(numberList, number)
	}
	file.Close()
	return numberList
}

// For part one, compare one number with the number directly in front.
func partOne(numberList []int) int {
	numberOfIncreases := 0
	for i := 0; i < len(numberList)-1; i++ {
		if numberList[i] < numberList[i+1] {
			numberOfIncreases++
		}
	}
	return numberOfIncreases
}

// For part two, instead of comparing the sliding window sum of three numbers, simply compare
// one number with the number three places ahead (the middle number in the sum windows cancel out!).
func partTwo(numberList []int) int {
	numberOfIncreases := 0
	for i := 0; i < len(numberList)-3; i++ {
		if numberList[i] < numberList[i+3] {
			numberOfIncreases++
		}
	}
	return numberOfIncreases
}
