package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	path := flag.String("path", "./problem.csv", "Path to csv file")
	flag.Parse()

	file, _ := os.Open(*path)
	csvReader := csv.NewReader(bufio.NewReader(file))

	var problemSet []Problem

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		problemSet = append(problemSet, Problem{
			line[0],
			line[1],
		})
	}

	var inputSolution string
	reader := bufio.NewReader(os.Stdin)
	score := 0

	fmt.Println("Starting Quiz...")

	for i := 0; i < len(problemSet); i++ {
		fmt.Println(problemSet[i].Question)
		inputSolution, _ = reader.ReadString('\n')
		if strings.TrimSpace(inputSolution) != strings.TrimSpace(problemSet[i].Answer) {
			fmt.Println("Incorrect Answer")
			break
		}
		score++
	}
	fmt.Println("Total Score:", score)
}
