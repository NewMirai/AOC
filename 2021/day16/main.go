package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

func readBinary(s string, start int, end int) int {
	var number int
	fmt.Sscanf(s[start:end], "%b", &number)
	return number
}

func DecodeHex(s string, total *int) {
	packetVersion := readBinary(s, 0, 3)
	typeID := readBinary(s, 3, 6)
	packetType := CheckID(typeID)
	if packetType == "literal" {
		DecodeLiteral(s, packetVersion, typeID, total)
	} else {
		DecodeOperator(s, packetVersion, typeID, total)
	}
}

func DecodeLiteral(s string, version int, typeID int, total *int) {
	var value int
	A, B, C := s[6:11], s[11:16], s[16:21]
	n1, n2, n3 := A[1:], B[1:], C[1:]
	tmp := n1 + n2 + n3
	fmt.Sscanf(tmp, "%b", &value)
	*total += value
}

func DecodeOperator(s string, version int, typeID int, total *int) {
	var subLength int
	lBitNum := GetLengthBit(s[6:7])
	subLengthS := s[7 : 7+lBitNum]
	fmt.Sscanf(subLengthS, "%b", &subLength)
	switch typeID {
	case 6:
		c := 0
		for c < 2 {
			switch c {
			case 0:
				packet := s[7+lBitNum : 7+lBitNum+11]
				DecodeHex(packet, total)
			}
			c++
		}
	}
}

func GetLengthBit(ltID string) int {
	cTable := map[string]int{
		"0": 15,
	}
	return cTable[ltID]
}

func CheckID(id int) string {
	switch id {
	case 4:
		return "literal"
	default:
		return "operator"
	}
}

func IsSmall(s string) bool {
	if s[0] == '0' {
		return true
	}
	return false
}

func Solve(s *string) (total int) {
	raw := strings.TrimSpace(*s)
	hex2Binary := map[string]string{
		"0":  "0000",
		"1":  "0001",
		"2":  "0010",
		"3":  "0011",
		"4 ": "0100",
		"5":  "0101",
		"6":  "0110",
		"7":  "0111",
		"8":  "1000",
		"9":  "1001",
		"A":  "1010",
		"B":  "1011",
		"C":  "1100",
		"D":  "1101",
		"E":  "1110",
		"F":  "1111",
	}
	var rawDecoded string
	for _, c := range raw {
		rawDecoded += hex2Binary[string(c)]
	}
	runtime.Breakpoint()
	DecodeHex(rawDecoded, &total)
	return total
}

func main() {
	problem_input := ReadInput("input.test2")
	fmt.Println(Solve(&problem_input))
}
