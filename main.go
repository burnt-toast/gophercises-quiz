package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", " a CSV filein the format of 'question,answer'")
	flag.Parse()

	//fmt.Println("The csvFileName is: ", *csvFileName)
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Unable to open file: ", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV")
	}

	problems := parseLines(lines)
	fmt.Println("Number of problems loaded: ", len(problems))

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

func parseLines(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}
	}
	return problems
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
