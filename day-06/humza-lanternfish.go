package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := loadFileToList("Input.txt")

	partOneList := recursion(list, 0, 80)
	partTwoList := recursion(partOneList, 80, 256)
	fmt.Println("The answer to part one is:", sum(partOneList))
	fmt.Println("The answer to part two is:", sum(partTwoList))
}

func loadFileToList(inputFileName string) map[int]int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	dict := map[int]int{}
	for _, s := range strings.Split(text, ",") {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		dict[number]++
	}

	defer file.Close()

	return dict
}

func recursion(list map[int]int, start, end int) map[int]int {
	for i := start; i < end; i++ {
		list = updateDictionary(list)
	}
	return list
}

func updateDictionary(list map[int]int) map[int]int {
	updatedDict := map[int]int{
		8: list[0],
		7: list[8],
		6: list[0] + list[7],
		5: list[6],
		4: list[5],
		3: list[4],
		2: list[3],
		1: list[2],
		0: list[1],
	}
	return updatedDict
}

func sum(list map[int]int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
