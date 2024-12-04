package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

var mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var mulRegex2 = regexp.MustCompile(`(do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\))`)

func main() {
	var path string
	flag.StringVar(&path, "f", "", "input file path")
	flag.Parse()
	if path == "" {
		panic("input file path is required")
	}
	sum1 := parseInput(path, false)
	log.Printf("Sum of all multiplications: %d", sum1)
	sum2 := parseInput(path, true)
	log.Printf("Sum of all multiplications with enabled flag: %d", sum2)

}

func tokenSplit(data []byte, atEOF bool) (int, []byte, error) {
	// log.Printf("data: %s, atEOF: %v", data, atEOF)
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	loc := mulRegex2.FindIndex(data)
	if loc != nil {
		tkn := data[loc[0]:loc[1]]
		adv := loc[1]
		// log.Printf("  found token: %s - advancing: %d", tkn, adv)
		return adv, tkn, nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func parseInput(path string, withToggle bool) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sum := 0
	enabled := true
	sc.Split(tokenSplit)
	for sc.Scan() {
		t := sc.Text()
		// log.Printf("token: %s", t)
		if withToggle {
			if t == "do()" {
				enabled = true
				continue
			} else if t == "don't()" {
				enabled = false
				continue
			}
		}
		if enabled {
			matches := mulRegex.FindAllStringSubmatch(t, -1)
			for _, m := range matches {
				a, err := strconv.Atoi(m[1])
				if err != nil {
					log.Fatalf("error converting %s to int: %v", m[1], err)
				}
				b, err := strconv.Atoi(m[2])
				if err != nil {
					log.Fatalf("error converting %s to int: %v", m[2], err)
				}
				sum += a * b
			}
		}
	}
	return sum
}
