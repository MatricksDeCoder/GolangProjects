package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/csv"
	"io"
	"github.com/gookit/color"
	"flag"
	"time"
)

const nameCSVFlag ="csv"
const defaultFile = "problems.csv"
const messageCSVFlag = "a csv file in the format of 'question,answer'"
const messageCSVFail = "Failed to open the CSV file: "
const messageStartQuiz = "Starting quiz.....\n"
const messageCorrect = "\nHere are the correct answers..."
const nameTimeFlag = "limit"
const defaultTime = 30
const messageTimeFlag = "the time limit for the quiz in seconds"
const messageTimeup = "\nSorry! Time limit has elaspsed"

type Problem struct {
	asked   string
	answer  string
}

func main() {
	// Quiz program max 100 questions
	// Enter or pipe in command line -csv flag  for file with questions
	csvFilename := flag.String(nameCSVFlag, defaultFile, messageCSVFlag)
	timeLimit := flag.Int(nameTimeFlag, defaultTime, messageTimeFlag)
	flag.Parse()

	// Open file to read content
	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println(messageCSVFail, *csvFilename)
		os.Exit(1)
	}
	defer csvFile.Close()

	// Save lines in map[string]string key=question, value=answer
	csvReader   := csv.NewReader(csvFile)

	//	Create array of problems
	questions := make([]Problem,0)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // or handle it another way
		}
		question := Problem{asked:row[0], answer:row[1]}
		questions = append(questions,question)
	}
	// Implement Quiz
	numQuestions := len(questions)
	correct := 0
	check := make([]bool,numQuestions)
	answers := make([]string, numQuestions)
	var score float64

	color.Green.Println(messageStartQuiz)

	// Implement Quiz timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for i,question := range questions {
		color.Yellow.Printf("Problem #%d: %s = ", i+1, question.asked)
		answerCh := make(chan string)
		go func() {
			userInput := ""
			inputScanner := bufio.NewScanner(os.Stdin)
			inputScanner.Scan()
			userInput = inputScanner.Text()
			userInput = strings.Trim(userInput, " ")
			answerCh <- userInput
		}()

		select {
		case <- timer.C : {
			fmt.Println(messageTimeup)
			fmt.Printf("Answered %d of %d",i,numQuestions)
			break problemLoop
		}
		case answers[i] = <- answerCh : {
			if answers[i] == question.answer {
				correct++
				check[i] = true
			}
		}
		}
	}
	score  = float64(correct)/float64(numQuestions) * 100
	fmt.Printf("\nYou scored: %d out of %d - Mark: %.2f\n",correct, numQuestions, score)
	//Re print question and answers with red color for wrong answers
	color.Yellow.Println(messageCorrect)
	for i,question := range questions {
		if check[i] == true {
			color.Green.Printf("Problem #%d: %s = %s\n", i+1, question.asked,question.answer)
		} else {
			color.Red.Printf("Problem #%d: %s = %s - You answered %s\n", i+1, question.asked, question.answer, answers[i])
		}
	}

}