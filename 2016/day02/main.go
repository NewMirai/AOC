package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadInput of the problem
func ReadInput(f string) string {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type PositionTable map[rune]string

// Solve problem
func Solve(s *string) string {
	instructions := strings.Split(strings.TrimSpace(*s), "\n")
	position := "5"
	var string_result string
	for _, instruction := range instructions {
		for _, movement := range instruction {
			switch position {
			case "1":
				table := PositionTable{
					'U': "1",
					'D': "4",
					'R': "2",
					'L': "1",
				}
				position = table[movement]

			case "2":
				table := PositionTable{
					'U': "2",
					'D': "5",
					'R': "3",
					'L': "1",
				}
				position = table[movement]

			case "3":
				table := PositionTable{
					'U': "3",
					'D': "6",
					'R': "3",
					'L': "2",
				}
				position = table[movement]
			case "4":
				table := PositionTable{
					'U': "1",
					'D': "7",
					'R': "5",
					'L': "4",
				}
				position = table[movement]
			case "5":
				table := PositionTable{
					'U': "2",
					'D': "8",
					'R': "6",
					'L': "4",
				}
				position = table[movement]
			case "6":
				table := PositionTable{
					'U': "3",
					'D': "9",
					'R': "6",
					'L': "5",
				}
				position = table[movement]
			case "7":
				table := PositionTable{
					'U': "4",
					'D': "7",
					'R': "8",
					'L': "7",
				}
				position = table[movement]
			case "8":
				table := PositionTable{
					'U': "5",
					'D': "8",
					'R': "9",
					'L': "7",
				}
				position = table[movement]
			case "9":
				table := PositionTable{
					'U': "6",
					'D': "9",
					'R': "9",
					'L': "8",
				}
				position = table[movement]

			}
		}
		string_result += position
	}
	return string_result
}

func Solve2(s *string) string {
	instructions := strings.Split(strings.TrimSpace(*s), "\n")
	position := "5"
	var string_result string
	for _, instruction := range instructions {
		for _, movement := range instruction {
			switch position {
			case "1":
				table := PositionTable{
					'U': "1",
					'D': "3",
					'R': "1",
					'L': "1",
				}
				position = table[movement]

			case "2":
				table := PositionTable{
					'U': "2",
					'D': "6",
					'R': "3",
					'L': "2",
				}
				position = table[movement]

			case "3":
				table := PositionTable{
					'U': "1",
					'D': "7",
					'R': "4",
					'L': "2",
				}
				position = table[movement]
			case "4":
				table := PositionTable{
					'U': "4",
					'D': "8",
					'R': "4",
					'L': "3",
				}
				position = table[movement]
			case "5":
				table := PositionTable{
					'U': "5",
					'D': "5",
					'R': "6",
					'L': "5",
				}
				position = table[movement]
			case "6":
				table := PositionTable{
					'U': "2",
					'D': "A",
					'R': "7",
					'L': "5",
				}
				position = table[movement]
			case "7":
				table := PositionTable{
					'U': "3",
					'D': "B",
					'R': "8",
					'L': "6",
				}
				position = table[movement]
			case "8":
				table := PositionTable{
					'U': "4",
					'D': "C",
					'R': "9",
					'L': "7",
				}
				position = table[movement]
			case "9":
				table := PositionTable{
					'U': "9",
					'D': "9",
					'R': "9",
					'L': "8",
				}
				position = table[movement]

			case "A":
				table := PositionTable{
					'U': "6",
					'D': "A",
					'R': "B",
					'L': "A",
				}
				position = table[movement]
			case "B":
				table := PositionTable{
					'U': "7",
					'D': "D",
					'R': "C",
					'L': "A",
				}
				position = table[movement]
			case "C":
				table := PositionTable{
					'U': "8",
					'D': "C",
					'R': "C",
					'L': "B",
				}
				position = table[movement]
			case "D":
				table := PositionTable{
					'U': "B",
					'D': "D",
					'R': "D",
					'L': "D",
				}
				position = table[movement]
			}
		}
		string_result += position
	}
	return string_result
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
