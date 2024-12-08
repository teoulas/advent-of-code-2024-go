package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "f", "example.txt", "input file path")
	flag.Parse()
	gameMap := loadMap(path)
	s1 := solve1(gameMap)
	log.Printf("Solution #1: %d", s1)
}

func solve1(m Map) int {
	visited := make(map[Point]bool)
	current := m.Start
	dir := m.Direction
	visited[current] = true
	for {
		// dumpState(m, visited)
		// fmt.Scanln()

		dr := directions[dir][0]
		dc := directions[dir][1]
		next := Point{Row: current.Row + dr, Col: current.Col + dc}
		if next.Row < 0 || next.Row >= m.Rows || next.Col < 0 || next.Col >= m.Cols {
			break // out of bounds
		}
		// check if there is an obstacle in front
		if hasObstacle(next, m.Obstacles) {
			// turn right/90 deg clockwise
			dir = (dir + 1) % 4
			continue
		}
		visited[next] = true
		current = next
	}
	return len(visited)
}

func dumpState(m Map, visits map[Point]bool) {
	fmt.Printf("\033[2J\033[H")
	fmt.Printf("Visits: %d\n", len(visits))
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			p := Point{Row: r, Col: c}
			if p == m.Start {
				fmt.Print("^")
			} else if hasObstacle(p, m.Obstacles) {
				fmt.Print("#")
			} else if visits[p] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Point struct {
	Row, Col int
}

var directions = [][]int{
	{-1, 0}, // up/north
	{0, 1},  // right/east
	{1, 0},  // down/south
	{0, -1}, // left/west
}

type Map struct {
	Rows, Cols int
	Start      Point
	Direction  int // 0: north, 1: east, 2: south, 3: west
	Obstacles  []Point
}

func hasObstacle(p Point, obstacles []Point) bool {
	for _, o := range obstacles {
		if o == p {
			return true
		}
	}
	return false
}

func (m *Map) AddRow(row string) {
	rowIdx := m.Rows
	m.Rows++
	colLen := len(row)
	if colLen > m.Cols {
		m.Cols = colLen
	}
	for colIdx, cell := range row {
		if cell == '#' { // obstacle
			m.Obstacles = append(m.Obstacles, Point{Row: rowIdx, Col: colIdx})
		} else if cell == '^' { // start facing north
			m.Start = Point{Row: rowIdx, Col: colIdx}
		}
	}
}

func loadMap(path string) Map {
	m := Map{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return m
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}
		m.AddRow(line)
	}
	return m
}
