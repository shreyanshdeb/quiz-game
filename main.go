package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Problem contains questions and answers
//obtained from the csv file
type Problem struct {
	Question string
	Answer   string
}

func main() {
	path := flag.String("csv", "./problem.csv", "Path to csv file")
	flag.Parse()

	problemSet := parseFile(*path)

	var inputSolution string
	reader := bufio.NewReader(os.Stdin)
	correctScore := 0
	incorrectScore := 0

	fmt.Println("Starting Quiz...")

	for i := 0; i < len(problemSet); i++ {
		fmt.Printf("Problem#%d: %s ", (i + 1), problemSet[i].Question)
		inputSolution, _ = reader.ReadString('\n')
		if strings.TrimSpace(inputSolution) != strings.TrimSpace(problemSet[i].Answer) {
			incorrectScore++
		} else {
			correctScore++
		}
	}
	fmt.Printf("You scored %d out of %d.", correctScore, len(problemSet))
}

func parseFile(path string) []Problem {
	file, _ := os.Open(path)
	csvReader := csv.NewReader(bufio.NewReader(file))

	var problemSet []Problem

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		problemSet = append(problemSet, Problem{
			strings.TrimSpace(line[0]),
			strings.TrimSpace(line[1]),
		})
	}

	return problemSet
}
