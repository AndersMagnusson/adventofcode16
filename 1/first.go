package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := `R3, R1, R4, L4, R3, R1, R1, L3, L5, L5, L3, R1, R4, L2, L1, R3, L3, R2, R1, R1, L5, L2, L1, R2, L4, R1, L2, L4, R2, R2, L2, L4, L3, R1, R4, R3, L1, R1, L5, R4, L2, R185, L2, R4, R49, L3, L4, R5, R1, R1, L1, L1, R2, L1, L4, R4, R5, R4, L3, L5, R1, R71, L1, R1, R186, L5, L2, R5, R4, R1, L5, L2, R3, R2, R5, R5, R4, R1, R4, R2, L1, R4, L1, L4, L5, L4, R4, R5, R1, L2, L4, L1, L5, L3, L5, R2, L5, R4, L4, R3, R3, R1, R4, L1, L2, R2, L1, R4, R2, R2, R5, R2, R5, L1, R1, L4, R5, R4, R2, R4, L5, R3, R2, R5, R3, L3, L5, L4, L3, L2, L2, R3, R2, L1, L1, L5, R1, L3, R3, R4, R5, L3, L5, R1, L3, L5, L5, L2, R1, L3, L1, L3, R4, L1, R3, L2, L2, R3, R3, R4, R4, R1, L4, R1, L5`
	// input := "R2, L3"
	way := strings.Split(input, ", ")
	x := 0
	y := 0
	nose := "up"
	for _, instruction := range way {

		direction := string(instruction[0])
		end := len(instruction)
		length, err := strconv.Atoi(string(instruction[1:end]))
		if err != nil {
			log.Fatal("Failed to pase length")
		}
		log.Print(length)
		switch nose {
		case "up":
			if direction == "R" {
				x = x + length
				nose = "right"
			} else {
				x = x - length
				nose = "left"
			}
		case "right":
			if direction == "R" {
				y = y - length
				nose = "down"
			} else {
				y = y + length
				nose = "up"
			}
		case "down":
			if direction == "R" {
				x = x - length
				nose = "left"
			} else {
				x = x + length
				nose = "right"
			}
		case "left":
			if direction == "R" {
				y = y + length
				nose = "up"
			} else {
				y = y - length
				nose = "down"
			}

		}
	}
	if x < 0 {
		x = x * -1
	}
	if y < 0 {
		y = y * -1
	}
	fmt.Printf("length: %d", x+y)
}
