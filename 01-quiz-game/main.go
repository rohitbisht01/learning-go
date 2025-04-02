package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	csvPath = flag.String("csv", "problems.csv", "A CSV file containing questions and answers")
	limit   = flag.Int("limit", 30, "The time limit to answer each question")
)

func main() {
	flag.Parse()

	// Open CSV file
	file, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Read CSV data
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	questions := make(map[string]string, len(csvData))
	for _, row := range csvData {
		questions[row[0]] = row[1]
	}

	scanner := bufio.NewScanner(os.Stdin)
	correct := 0
	total := len(questions)

	fmt.Println("Quiz Game Starts!")
	var ind int64

	// Loop through all questions
	for question, answer := range questions {
		fmt.Printf("Question %d: %s ", ind+1, question)
		ind++

		// Create a channel for user input
		userAnswerCh := make(chan string)

		// Goroutine to read user input
		go func() {
			scanner.Scan()
			userAnswerCh <- strings.TrimSpace(strings.ToLower(scanner.Text()))
		}()

		select {
		case userAnswer := <-userAnswerCh:
			// If user answers within time
			if strings.ToLower(answer) == userAnswer {
				fmt.Println("✅ Correct!")
				correct++
			} else {
				fmt.Println("❌ Incorrect! The correct answer was:", answer)
			}
		case <-time.After(time.Duration(*limit) * time.Second):
			// If time runs out
			fmt.Println("\n⏰ Time's up! Moving to the next question.")
		}
	}

	// Final Score
	fmt.Println("\nTotal Questions:", total)
	fmt.Println("Correct Answers:", correct)
	fmt.Println("Incorrect Answers:", total-correct)
}
