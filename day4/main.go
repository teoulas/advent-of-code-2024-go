package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "f", "", "input file path")
	flag.Parse()
	if path == "" {
		panic("input file path is required")
	}
	g := readInput(path)
	num := findString(g, "XMAS")
	log.Printf("Number of XMAS found: %d", num)
	num2 := findXmas(g)
	log.Printf("Number of MAS in X form found: %d", num2)
}

//  1      2      3      4
// ---    ---    ---    ---
// M M    M S    S S    S M
//  A      A      A      A
// S S    M S    M M    S M

func findXmas(g Grid) int {
	n := 0
	for r := range g.Rows {
		if r == 0 || r == len(g.Rows)-1 { // we're looking for A so the edges are not interesting
			continue
		}
		for c := range g.Rows[r].Cells {
			if c == 0 || c == len(g.Rows[r].Cells)-1 { // we're looking for A so the edges are not interesting
				continue
			}
			if g.Get(r, c) == 'A' {
				n += checkForXmas(g, r, c)
			}
		}
	}
	return n
}

var xDirs = [][]int{
	{-1, -1}, // N W
	{-1, 1},  // N E
	{1, 1},   // S E
	{1, -1},  // S W
}

func checkForXmas(g Grid, r int, c int) int {
	// clock-wise NW, NE, SE, SW
	// valid values: MMSS, MSSM, SSMM, SMMS
	n := 0
	str := ""
	for i := range xDirs {
		str += string(g.Get(r+xDirs[i][0], c+xDirs[i][1]))
	}
	if str == "MMSS" || str == "MSSM" || str == "SSMM" || str == "SMMS" {
		n++
	}
	return n
}

func findString(g Grid, str string) int {
	n := 0
	runes := []rune(str)
	for r := range g.Rows {
		for c := range g.Rows[r].Cells {
			n += checkForString(g, runes, r, c)
		}
	}
	return n
}

func checkForString(g Grid, runes []rune, r int, c int) int {
	n := 0
	for _, d := range directions {
		found := true
		for i := range runes {
			if g.Get(r+d[0]*i, c+d[1]*i) != runes[i] {
				found = false
				break
			}
		}
		if found {
			n++
		}
	}
	return n
}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Grid struct {
	Rows []Row
}

func (g Grid) Get(r, c int) rune {
	if r < 0 || r >= len(g.Rows) {
		return '?'
	}
	row := g.Rows[r]
	if c < 0 || c >= len(row.Cells) {
		return '?'
	}
	return row.Cells[c]
}

func (g Grid) String() string {
	s := ""
	for i := range g.Rows {
		s += string(g.Rows[i].Cells) + "\n"
	}
	return s
}

type Row struct {
	Cells []rune
}

func readInput(path string) Grid {
	g := Grid{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return g
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		r := Row{}
		for _, c := range line {
			r.Cells = append(r.Cells, c)
		}
		g.Rows = append(g.Rows, r)
	}
	return g
}
