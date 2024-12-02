package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path string
	flag.StringVar(&path, "f", "", "input file path")
	flag.Parse()
	if path == "" {
		panic("input file path is required")
	}
	reports := readInput(path)
	numSafe := 0
	numSafeDampened := 0
	for i := range reports {
		s := reports[i].IsSafe()
		if s == 0 {
			numSafe++
		}
		if s == 1 {
			numSafeDampened++
		}
	}
	log.Printf("Number of safe rows: %d", numSafe)
	log.Printf("Number of safe rows with dampening: %d", numSafe+numSafeDampened)
}

type Report struct {
	Levels []int
}

func (r *Report) IsSafe() int {
	numFails := 0
	diff := 0
	for i := 0; i < len(r.Levels)-1; i++ {
		curr := r.Levels[i]
		next := r.Levels[i+1]
		if curr == next { // equal=fail
			numFails++
		}
		d := next - curr
		if d > 3 || d < -3 { // too big of a difference=fail
			numFails++
		}
		if diff == 0 {
			diff = d
		}
		if diff*d < 0 {
			numFails++ // changed direction
		}
	}
	return numFails
}

func readInput(path string) []Report {
	reports := make([]Report, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error reading file: %v", err)
			break
		}
		strs := strings.Split(line[0:len(line)-1], " ")
		nums := make([]int, len(strs))
		for i := range strs {
			n, err := strconv.Atoi(strs[i])
			if err != nil {
				log.Fatalf("error converting string to int: %v", err)
				break
			}
			nums[i] = n
		}
		reports = append(reports, Report{Levels: nums})
	}
	return reports
}
