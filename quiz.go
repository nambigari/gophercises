package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type QA struct {
	number1       int
	number2       int
	correctAnswer int
	answerByUser  int
}

func (qa *QA) Init(record []string) {
	qa.number1, _ = strconv.Atoi(record[0])
	qa.number2, _ = strconv.Atoi(record[1])
	qa.correctAnswer, _ = strconv.Atoi(record[2])
}

func ReadFile(filePath string) []QA {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("No file")
	}
	r := csv.NewReader(strings.NewReader(string(data)))
	var quiz []QA
	for {
		record, err := r.Read()
		var qa QA
		if len(record) == 3 {
			qa.Init(record)
			quiz = append(quiz, qa)
		}
		if err == io.EOF {
			break
		}
	}
	return quiz
}

func Score(quiz []QA) (int, int) {
	NumberOfQuestions := len(quiz)
	var CorrectAnswers int
	for _, qa := range quiz {
		if qa.answerByUser == qa.correctAnswer {
			CorrectAnswers++
		}
	}
	return NumberOfQuestions, CorrectAnswers
}

func main() {
	argsWithProg := os.Args
	pathToProblems := "problems.csv"
	if len(argsWithProg) > 1 {
		pathToProblems = argsWithProg[1]
	}
	quiz := ReadFile(pathToProblems)
	for idx, qa := range quiz {
		fmt.Println(qa.number1, " + ", qa.number2, " ?")
		var foo int
		fmt.Scan(&foo)
		fmt.Println("Input read is :", foo)
		quiz[idx].answerByUser = foo
	}
	NumberOfQuestions, CorrectAnswers := Score(quiz)
	fmt.Println("Number of questions : ", NumberOfQuestions)
	fmt.Println("Correct ansers : ", CorrectAnswers)
	fmt.Println("pecent correct :", (float32(CorrectAnswers)/float32(NumberOfQuestions))*100)
}
