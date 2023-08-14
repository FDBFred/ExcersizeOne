package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var score float64 = 0

func main() {

	file, err := os.Open("quiz-master/problems.csv")
	if err != nil {
		log.Fatalln("Error opening file", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Read the file
	reader := csv.NewReader(file)

	reader.Comma = ','

	records, err := reader.ReadAll()

	records = shuffle(records)

	fmt.Println("Welcome to the Quiz Game")
	fmt.Println("Please answer the following questions")

	time.Sleep(5 * time.Second)

	startTime := time.Now()

	for i := 0; i < len(records); i++ {
		fmt.Println(records[i][0])

		var answer int

		_, err := fmt.Scanf("%d\n", &answer)
		if err != nil {
			fmt.Println("Enter A Number")
		}

		ansCorr, err := strconv.Atoi(records[i][1])

		if answer == ansCorr {
			score += 1
		}
	}

	endTime := time.Now()

	scoreTime := endTime.Sub(startTime)

	fmt.Println("Time taken: ", scoreTime)

	fmt.Println("Your score was: ", score)

	scoreTimeSecs := scoreTime.Seconds()

	LengthAsFloat := float64(len(records))

	fmt.Println("Your overall score is: ", ((LengthAsFloat-score)+1)*scoreTimeSecs)

	time.Sleep(10 * time.Second)

}

func shuffle(array [][]string) [][]string {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}
