package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: renamer listA listB")
	}

	fA, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening %s: %v\n", os.Args[1], err)
	}
	defer fA.Close()

	fB, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalf("Error opening %s: %v\n", os.Args[2], err)
	}
	defer fB.Close()

	sA := bufio.NewScanner(fA)
	sB := bufio.NewScanner(fB)

	for {
		if !sA.Scan() || !sB.Scan() {
			break
		}

		lineA := strings.TrimSpace(sA.Text())
		lineB := strings.TrimSpace(sB.Text())

		if lineA == lineB {
			continue
		}

		if err := os.Rename(lineA, lineB); err != nil {
			log.Printf("Error renaming %s -> %s: %v\n", lineA, lineB, err)
		}
	}
	if err := sA.Err(); err != nil {
		log.Fatalf("Error reading %s: %v\n", os.Args[1], err)
	}
	if err := sB.Err(); err != nil {
		log.Fatalf("Error reading %s: %v\n", os.Args[2], err)
	}
}
