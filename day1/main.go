package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"slices"
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
	lists := readInput(path)
	for i := range lists {
		slices.Sort(lists[i])
	}
	sum1 := 0 // part 1
	sum2 := 0 // part 2
	for i := range lists[0] {
		d := lists[1][i] - lists[0][i]
		sum1 += abs(d)
		sim := occursIn(lists[0][i], lists[1])
		sum2 += sim * lists[0][i]
	}
	log.Printf("sum: %d", sum1)
	log.Printf("similarity: %d", sum2)
}

type NumList []int

func readInput(path string) []NumList {
	l1 := make([]int, 0)
	l2 := make([]int, 0)

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
		str := strings.Split(line[0:len(line)-1], "   ")
		n1, err := strconv.Atoi(str[0])
		if err != nil {
			log.Fatalf("error converting string to int: %v", err)
			break
		}
		n2, err := strconv.Atoi(str[1])
		if err != nil {
			log.Fatalf("error converting string to int: %v", err)
			break
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	return []NumList{l1, l2}
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func occursIn(n int, l NumList) int {
	occ := 0
	for i := range l {
		if n == l[i] {
			occ++
		}
	}
	return occ
}
