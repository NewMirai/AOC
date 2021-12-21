package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type Packet interface {
	Decode()
	SumVersion() int
}

type Header struct {
	packetVersion int
	typeID        int
}

// A single packet
type LiteralPacket struct {
	repr   string
	header Header
	value  int
}

func (lp *LiteralPacket) GetNumber() (number int) {
	sRepr := lp.repr[6:]
	if IsLastGroup(sRepr[0]) {
		number = ReadBitNumber(sRepr[1:])
		return
	}
	N := len(sRepr)
	nParts := N / 5
	parts := make([]string, nParts)
	start, end := 0, 5
	for i := 0; i < nParts; i++ {
		switch i {
		case nParts - 1:
			parts[i] = sRepr[start+1 : start+1+4]
		default:
			parts[i] = sRepr[start+1 : end]
		}
		start += 5
		end += 5
	}
	binStr := strings.Join(parts, "")
	number = ReadBitNumber(binStr)
	return
}

func (lp *LiteralPacket) Decode() {
	value := lp.GetNumber()
	lp.value = value
}

func (lp *LiteralPacket) SumVersion() int { return lp.header.packetVersion }

// Contains 1 or more packets
type OperatorPacket struct {
	repr      string
	header    Header
	lenTypeID int
	bitLen    int // total bit length of subpacket
	nPackets  int
	packets   []Packet
}

func (op *OperatorPacket) Decode() {
	runtime.Breakpoint()
	switch op.lenTypeID {
	case 0:
		op.bitLen = ReadBitNumber(op.repr[7 : 7+15])
		subPackets := op.repr[7+15 : 7+15+op.bitLen]
		op.nPackets = len(subPackets) / 11
		start, end := 0, 11
		packets := make([]Packet, op.nPackets)
		for i := 0; i < op.nPackets; i++ {
			switch i {
			case op.nPackets - 1:
				rawPacket := subPackets[start:]
				p := bin2Packet(rawPacket)
				packets[i] = p

			default:
				rawPacket := subPackets[start:end]
				p := bin2Packet(rawPacket)
				packets[i] = p
			}
			start += 11
			end += 11
		}
		for _, packet := range packets {
			packet.Decode()
		}
		op.packets = packets
	case 1:
		op.nPackets = ReadBitNumber(op.repr[7 : 7+11])
		subPackets := op.repr[7+11:]
		start, end := 0, 11
		packets := make([]Packet, op.nPackets)
		for i := 0; i < op.nPackets; i++ {
			rawPacket := subPackets[start:end]
			p := bin2Packet(rawPacket)
			packets[i] = p
			start += 11
			end += 11
		}
		for _, packet := range packets {
			packet.Decode()
		}
		op.packets = packets
	}
}

func (op *OperatorPacket) SumVersion() (total int) {
	for _, packet := range op.packets {
		total = packet.SumVersion()
	}
	return
}

// Convert hex to binary
// hex -> byte -> binary string
func Hex2Binary(s string) (decoded string) {
	byteArray, _ := hex.DecodeString(s)
	binArray := make([]string, len(byteArray))
	for i := range byteArray {
		binArray[i] = fmt.Sprintf("%08b", byteArray[i])
	}
	decoded = strings.Join(binArray, "")
	return
}

// Read bit number
func ReadBitNumber(s string) (number int) {
	fmt.Sscanf(s, "%b", &number)
	return
}

// Check if last group
func IsLastGroup(r byte) bool {
	if r == '0' {
		return true
	}
	return false
}

// Convert string to Packet partially initialize
func bin2Packet(s string) (p Packet) {
	packetVersion, typeID := s[0:3], s[3:6]
	header := Header{
		packetVersion: ReadBitNumber(packetVersion),
		typeID:        ReadBitNumber(typeID),
	}
	switch ReadBitNumber(typeID) {
	case 4:
		p = &LiteralPacket{
			repr:   s,
			header: header,
		}
	default:
		lenTypeID, _ := strconv.Atoi(string(s[6]))
		p = &OperatorPacket{
			repr:      s,
			header:    header,
			lenTypeID: lenTypeID,
		}
	}
	return
}

func Solve(s *string) (total int) {
	raw := strings.TrimSpace(*s)
	rawDecoded := Hex2Binary(raw)
	p := bin2Packet(rawDecoded)
	p.Decode()
	total = p.SumVersion()
	return total
}

func main() {
	problem_input := ReadInput("input.test4")
	fmt.Println(Solve(&problem_input))
}
