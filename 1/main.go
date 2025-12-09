package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type state struct {
	code  int
	arrow int
}

func (s *state) turn(line string) {
	dir := 1
	if line[:1] == "L" {
		dir = -1
	}

	move, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Println(err)
	}

	s.arrow = ((s.arrow+dir*move)%100 + 100) % 100

	if s.arrow == 0 {
		s.code += 1
	}
}

func main() {
	s := state{
		code:  0,
		arrow: 50,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		s.turn(line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Print(s.code)

}
