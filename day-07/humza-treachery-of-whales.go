package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	list, min, max := loadFileToList("Input.txt")

	fmt.Println("The answer to part one is:", solveForPart("one", list, min, max))
	fmt.Println("The answer to part two is:", solveForPart("two", list, min, max))
}

func loadFileToList(inputFileName string) ([]int, int, int) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	var numbers []int
	for _, s := range strings.Split(text, ",") {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	defer file.Close()

	min, max := MinMax(numbers)
	return numbers, min, max
}

func solveForPart(part string, list []int, min, max int) int {
	var possibleFuels []int
	for i := min; i <= max; i++ {
		fuel := 0
		for _, s := range list {
			distance := int(math.Abs(float64(s - i)))
			fuel += fuelCalculation(part, distance)
		}
		possibleFuels = append(possibleFuels, fuel)
	}
	minFuel, _ := MinMax(possibleFuels)
	return minFuel
}

func fuelCalculation(part string, distance int) int {
	switch part {
	case "one":
		return distance
	case "two":
		// Using the formula for the sum of N numbers.
		return int((distance) * (distance + 1) / 2)
	}
	return 0
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
