package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type numberPair struct {
	Number int
	Called bool
}

type bingoCard struct {
	Rows []row
}

type row struct {
	Values []numberPair
}

const (
	gridLength = 5
)

func main() {
	var inputFileName = "Input.txt"
	list := loadFileToList(inputFileName)

	bingoNumbers := getBingoNumbers(list[0])
	bingoCards := getBingoBoards(list)

	fmt.Println("The answer to part one is:", partOne(&bingoNumbers, &bingoCards))
	fmt.Println("The answer to part two is:", partTwo(&bingoNumbers, &bingoCards))
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
	lines = append(lines, "")
	defer file.Close()

	return lines
}

func partOne(bingoNumbers *[]int, bingoCards *[]bingoCard) int {
	for _, s := range *bingoNumbers {
		for _, v := range *bingoCards {
			crossNumberOnBoard(s, &v)
			validRow := isRowValid(&v)
			validColumn := isColumnValid(&v)

			if validColumn != nil || validRow != nil {
				return s * sumUncheckedNumbersOnBoard(&v)
			}
		}
	}
	return 0
}

func partTwo(bingoNumbers *[]int, bingoCards *[]bingoCard) int {
	for i, s := range *bingoNumbers {
		for _, v := range *bingoCards {
			crossNumberOnBoard(s, &v)
			*bingoCards = removeCompletedBoard(bingoCards, &v)

			if len(*bingoCards) == 1 {
				remainingNumbers := (*bingoNumbers)[i:len(*bingoNumbers)]
				return partOne(&remainingNumbers, bingoCards)
			}
		}
	}
	return 0
}

func removeCompletedBoard(bingoCards *[]bingoCard, board *bingoCard) []bingoCard {
	var updatedCards []bingoCard
	for _, s := range *bingoCards {
		validRow := isRowValid(&s)
		validColumn := isColumnValid(&s)
		if validColumn == nil && validRow == nil {
			updatedCards = append(updatedCards, s)
		}
	}
	return updatedCards
}

func crossNumberOnBoard(number int, board *bingoCard) bingoCard {
	for i, s := range board.Rows {
		for j, v := range s.Values {
			if v.Number == number {
				board.Rows[i].Values[j].Called = true
			}
		}
	}
	return *board
}

func sumUncheckedNumbersOnBoard(board *bingoCard) int {
	sum := 0
	for _, s := range board.Rows {
		for _, v := range s.Values {
			if v.Called == false {
				sum += v.Number
			}
		}
	}
	return sum
}

func isRowValid(board *bingoCard) []int {
	count := 0
	for _, s := range board.Rows {
		for _, v := range s.Values {
			if v.Called == false {
				count = 0
				break
			}
			count++
		}
		if count == gridLength {
			var array []int
			for j := 0; j < gridLength; j++ {
				array = append(array, s.Values[j].Number)
			}
			return array
		}
	}
	return nil
}

func isColumnValid(board *bingoCard) []int {
	count := 0
	for i := 0; i < gridLength; i++ {
		for _, s := range board.Rows {
			if s.Values[i].Called == false {
				count = 0
				break
			}
			count++
		}
		if count == gridLength {
			var array []int
			for j := 0; j < gridLength; j++ {
				array = append(array, board.Rows[j].Values[i].Number)
			}
			return array
		}
	}
	return nil
}

func convertStringToRow(line *string) row {
	var row row
	var pairs []numberPair
	splitLine := strings.Split(*line, " ")
	for _, s := range splitLine {
		if s == "" {
			continue
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		pair := numberPair{Number: number, Called: false}
		pairs = append(pairs, pair)
	}
	row.Values = pairs
	return row
}

func getBingoNumbers(line string) []int {
	var bingoNumbers []int
	for _, s := range strings.Split(line, ",") {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		bingoNumbers = append(bingoNumbers, number)
	}
	return bingoNumbers
}

func getBingoBoards(lines []string) []bingoCard {
	var bingoCards []bingoCard
	var rows []row

	for i := 2; i < len(lines); i++ {
		if lines[i] != "" {
			row := convertStringToRow(&lines[i])
			rows = append(rows, row)
		} else {
			bingoCards = append(bingoCards, bingoCard{Rows: rows})
			rows = []row{}
		}
	}
	return bingoCards
}
