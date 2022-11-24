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

type Coordinates struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

const (
	gridSizeX = 990
	gridSizeY = 990
)

func main() {
	list := loadFileToList("Input.txt")

	coordinatesList, maxX, maxY := mapCoordinates(list)
	fmt.Printf("Make sure gridSizeX and gridSizeY are set to: %v and %v\n", maxX+1, maxY+1)

	partOneGrid := partOne(coordinatesList)
	partOneAnswer := countGrid(partOneGrid)
	partTwoGrid := partTwo(partOneGrid, coordinatesList)
	partTwoAnswer := countGrid(partTwoGrid)

	fmt.Println("The answer to part one is:", partOneAnswer)
	fmt.Println("The answer to part two is:", partTwoAnswer)
}

func loadFileToList(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	defer file.Close()

	return lines
}

func mapCoordinates(lines []string) ([]Coordinates, int, int) {
	var maxX, maxY float64
	var coordinatesList []Coordinates
	for _, s := range lines {
		coordinates := generateCoordinates(s)
		maxX, maxY = updateMaxXMaxY(maxX, maxY, coordinates)
		coordinatesList = append(coordinatesList, coordinates)
	}

	return coordinatesList, int(maxX), int(maxY)
}

func partOne(list []Coordinates) [gridSizeY][gridSizeX]int {
	var grid [gridSizeY][gridSizeX]int
	for _, s := range list {
		if s.x1 == s.x2 {
			if s.y1 < s.y2 {
				for i := s.y1; i <= s.y2; i++ {
					grid[i][s.x1]++
				}
			} else {
				for i := s.y2; i <= s.y1; i++ {
					grid[i][s.x1]++
				}
			}
		} else if s.y1 == s.y2 {
			if s.x1 < s.x2 {
				for i := s.x1; i <= s.x2; i++ {
					grid[s.y1][i]++
				}
			} else {
				for i := s.x2; i <= s.x1; i++ {
					grid[s.y1][i]++
				}
			}
		}
	}
	return grid
}

func partTwo(grid [gridSizeY][gridSizeX]int, list []Coordinates) [gridSizeY][gridSizeX]int {
	for _, s := range list {
		if coordinatesAreDiagonal(s) {
			if s.x1 < s.x2 && s.y1 < s.y2 {
				j := 0
				for i := s.x1; i <= s.x2; i++ {
					grid[s.y1+j][i]++
					j++
				}
			} else if s.x1 < s.x2 && s.y2 < s.y1 {
				j := 0
				for i := s.x1; i <= s.x2; i++ {
					grid[s.y1-j][i]++
					j++
				}
			} else if s.x2 < s.x1 && s.y1 < s.y2 {
				j := 0
				for i := s.x1; i >= s.x2; i-- {
					grid[s.y1+j][i]++
					j++
				}
			} else if s.x2 < s.x1 && s.y2 < s.y1 {
				j := 0
				for i := s.x1; i >= s.x2; i-- {
					grid[s.y1-j][i]++
					j++
				}
			}
		}
	}
	return grid
}

func countGrid(grid [gridSizeY][gridSizeX]int) int {
	count := 0
	for _, s := range grid {
		for _, v := range s {
			if v > 1 {
				count++
			}
		}
	}
	return count
}

func coordinatesAreDiagonal(s Coordinates) bool {
	return s.x1 != s.x2 && math.Abs(float64((s.y1-s.y2)/(s.x1-s.x2))) == 1
}

func updateMaxXMaxY(maxX, maxY float64, s Coordinates) (float64, float64) {
	maxX = math.Max(maxX, float64(s.x1))
	maxX = math.Max(maxX, float64(s.x2))
	maxY = math.Max(maxY, float64(s.y1))
	maxY = math.Max(maxY, float64(s.y2))
	return maxX, maxY
}

func generateCoordinates(text string) Coordinates {
	splitString := strings.Split(text, " -> ")
	firstCoordinate := strings.Split(splitString[0], ",")
	secondCoodinate := strings.Split(splitString[1], ",")
	return Coordinates{x1: stringToInt(firstCoordinate[0]), y1: stringToInt(firstCoordinate[1]),
		x2: stringToInt(secondCoodinate[0]), y2: stringToInt(secondCoodinate[1])}
}

func stringToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	return number
}
