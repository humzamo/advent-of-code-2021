package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instructionPair struct {
	Direction string
	Places    int
}

func main() {
	var inputFileName = "Input.txt"
	var instructionList = loadFileToList(inputFileName)

	solution := solveTravel(instructionList)
	fmt.Println("The answer to part one is:", solution[0])
	fmt.Println("The answer to part two is:", solution[1])
}

func loadFileToList(inputFileName string) []instructionPair {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var instructionList []instructionPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructionList = append(instructionList, splitInstruction(scanner.Text()))
	}
	defer file.Close()
	return instructionList
}

func solveTravel(instructionList []instructionPair) [2]int {
	var horizontalTravel, verticalTravel, aim int
	for _, s := range instructionList {
		switch s.Direction {
		case "forward":
			horizontalTravel += s.Places
			verticalTravel += aim * s.Places
		case "down":
			aim += s.Places
		case "up":
			aim -= s.Places
		}
	}

	// For both parts, the horizontal travel is the same.
	// For part one, the aim is equivalent to the vertical travel.
	partOneAnswer := horizontalTravel * aim
	partTwoAnswer := horizontalTravel * verticalTravel

	return [2]int{partOneAnswer, partTwoAnswer}
}

// Helper method to split a string into an instruction pair.
func splitInstruction(command string) instructionPair {
	splitCommand := strings.Split(command, " ")
	places, err := strconv.Atoi(splitCommand[1])
	if err != nil {
		log.Fatal(err)
	}

	return instructionPair{
		Direction: splitCommand[0],
		Places:    places,
	}
}
