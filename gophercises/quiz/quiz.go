package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"time"
)

var csvFile *string = flag.String("csv", "example.csv", "CSV file with all the questions")
var timeout *int = flag.Int("timeout", 10, "Timeout in seconds")
var questions [][]string

func check(msg string, err error) {
	if err != nil {
		glog.Errorf(msg+" Error: %v", err)
		panic(err)
	}
}

func readQuestions() {
	glog.V(3).Infof("Reading questions from file %s", *csvFile)
	f, err := os.Open(*csvFile)
	check(fmt.Sprintf("Failed to open csvFile %s.", *csvFile), err)
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	questions, err = r.ReadAll()
	check(fmt.Sprintf("Failed to read from csv file %s.", csvFile), err)
}

func init() {
	flag.Parse()
	glog.V(3).Infof("init() method called")
	readQuestions()
}

func getInput(inputChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
}

func startTimer(c chan bool) {
	glog.V(3).Infof("Starting timer for %v seconds", *timeout)
	time.Sleep(time.Duration(*timeout) * time.Second)
	c <- true
}

func startQuiz() {
	glog.V(3).Infof("Starting quiz")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press enter to start quiz")
	if confirm := scanner.Scan(); !confirm {
		check("Failed to read user input", scanner.Err())
	}
	timerChan := make(chan bool, 1)
	go startTimer(timerChan)
	inputChan := make(chan string, 1)
	go getInput(inputChan)
	score := 0
Loop:
	for _, line := range questions {
		if len(line) != 2 {
			glog.Errorf("Skipping invalid question: %v", line)
			continue
		}
		question := line[0]
		answer := line[1]
		fmt.Printf("%s :", question)
		select {
		case <-timerChan:
			fmt.Println("Time is up!")
			break Loop
		case resp := <-inputChan:
			if resp == answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong!")
			}

		}
	}
	fmt.Printf("You scored %d/%d\n", score, len(questions))
}

func main() {
	startQuiz()
}
