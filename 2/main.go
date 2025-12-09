package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "18623-26004,226779-293422,65855-88510,868-1423,248115026-248337139,903911-926580,97-121,67636417-67796062,24-47,6968-10197,193-242,3769-5052,5140337-5233474,2894097247-2894150301,979582-1016336,502-646,9132195-9191022,266-378,58-91,736828-868857,622792-694076,6767592127-6767717303,2920-3656,8811329-8931031,107384-147042,941220-969217,3-17,360063-562672,7979763615-7979843972,1890-2660,23170346-23308802"

	invalids := 0

	for v := range strings.SplitSeq(input, ",") {
		r := strings.Split(v, "-")

		lBound, err := strconv.Atoi(r[0])
		if err != nil {
			log.Fatalf("error reading file: %s", err)
		}

		hBound, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatalf("error reading file: %s", err)
		}

		for i := lBound; i <= hBound; i++ {
			s := strconv.Itoa(i)
			l, r := s[:len(s)/2], s[len(s)/2:]

			if l == r {
				invalids += i
			}
		}
	}

	fmt.Print(invalids)
}
