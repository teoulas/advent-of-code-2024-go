package main

import (
	"bufio"
	"flag"
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
	rules, pageUpdates := readInput(path)
	sumCorrect := 0
	updatesToFix := make(PageUpdates, 0)
	for _, pages := range pageUpdates {
		qualify := true
		for i := 1; i < len(pages); i++ {
			x := pages[i-1]
			y := pages[i]
			if !rules.AllowXAfterY(y, x) {
				qualify = false
				updatesToFix = append(updatesToFix, pages)
				break
			}
		}
		if qualify {
			middlePage := findMiddle(pages)
			sumCorrect += middlePage
		}
	}
	log.Printf("sum of correct updates: %d", sumCorrect)
	// part2
	sum2 := 0
	for _, pages := range updatesToFix {
		// log.Printf("Checking line %v", pages)
		for i := 1; i < len(pages); i++ {
			x := pages[i-1]
			y := pages[i]
			if !rules.AllowXAfterY(y, x) {
				pages[i-1], pages[i] = pages[i], pages[i-1]
				// log.Printf("Swapped %s and %s => %v", x, y, pages)
				if i > 1 { // go back to check the previous pair
					i -= 2
				}
				continue
			}
		}
		// assume it's fixed and calculate the middle page
		middlePage := findMiddle(pages)
		sum2 += middlePage
	}
	log.Printf("sum of updates after fixing: %d", sum2)
}

type OrderRule struct {
	Left  string
	Right string
}
type OrderRules []OrderRule

func (o *OrderRules) AddRule(left string, right string) {
	*o = append(*o, OrderRule{Left: left, Right: right})
}

// if not found, assume allowed, so we only need to find rules that disallow
func (o OrderRules) AllowXAfterY(x, y string) bool {
	for _, rule := range o {
		if rule.Left == x && rule.Right == y {
			return false
		}
	}
	return true
}

type PageUpdates [][]string

func findMiddle(pages []string) int {
	midIdx := len(pages) / 2
	midNum, err := strconv.Atoi(pages[midIdx])
	if err != nil {
		return 0
	}
	return midNum
}

func readInput(path string) (OrderRules, PageUpdates) {
	rules := make(OrderRules, 0)
	updates := make(PageUpdates, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return rules, updates
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}
		if len(line) == 5 && line[2] == '|' {
			rules.AddRule(line[0:2], line[3:5])
		}
		if line[2] == ',' {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	return rules, updates
}
