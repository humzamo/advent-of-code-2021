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

	//For part one, compare one number with the number directly in front.
	partOneAnswer := GetNumberOfIncreases(numberList, 1)
	fmt.Println("The answer to part one is:", partOneAnswer)

	// For part two, instead of comparing the sliding window sum of three numbers, simply compare
	// one number with the number three places ahead (the middle numbers in the sum windows cancel out!).
	partTwoAnswer := GetNumberOfIncreases(numberList, 3)
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
	defer file.Close()
	return numberList
}

// Gets number of times a number as increased compared to a number a certain offset away.
func GetNumberOfIncreases(numberList []int, offset int) int {
	numberOfIncreases := 0
	for i := 0; i < len(numberList)-offset; i++ {
		if numberList[i] < numberList[i+offset] {
			numberOfIncreases++
		}
	}
	return numberOfIncreases
}
