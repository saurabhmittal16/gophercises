package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// parse limit and random flag
	csvFileName := flag.String("file", "problems.csv", "give filename containing problems")
	limit := flag.Int("limit", 30, "set duration for level")
	random := flag.Bool("shuffle", false, "show problems in random order")
	flag.Parse()

	// opens csv file
	file, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatal(err)
	}

	// NewReader accepts types with Reader interface
	// os.Open returns type file which implements Reader
	r := csv.NewReader(file)

	// read all lines at once
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to parse the provided CSV file.")
	}

	// if random flag, randomize records
	if *random == true {
		shuffleRecords(records)
	}

	// start game
	var input string
	correct := 0

	fmt.Println("Press s to start game")
	fmt.Scanf("%s", &input)

	if input == "s" {
		// start timer
		timer := time.NewTimer(time.Second * time.Duration(*limit))

		for i, record := range records {
			fmt.Printf("Problem #%d: %s = ", i+1, record[0])
			answerCh := make(chan string)

			go func() {
				var answer string
				// read input string from stdin
				fmt.Scanf("%s\n", &answer)
				answerCh <- answer
			}()

			select {
			case <-timer.C:
				fmt.Printf("\nYou scored %d out of %d.\n", correct, len(records))
				return
			case ans := <-answerCh:
				if ans == record[1] {
					correct++
				}
			}
		}

		fmt.Printf("You scored %d out of %d.\n", correct, len(records))
	}
}

func shuffleRecords(records [][]string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range records {
		newpos := r.Intn(len(records) - 1)
		records[i], records[newpos] = records[newpos], records[i]
	}
}
