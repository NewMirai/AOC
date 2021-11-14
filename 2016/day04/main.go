package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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

type Counter map[string]int
type Pairs struct {
	key   string
	value int
}
type ByOccurences []Pairs

func (o ByOccurences) Len() int      { return len(o) }
func (o ByOccurences) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o ByOccurences) Less(i, j int) bool {
	if o[i].value == o[j].value {
		return o[i].key < o[j].key
	}
	return o[i].value > o[j].value
}

func InitCounter(name string) Counter {
	c := make(Counter)
	for _, r := range name {
		s := string(r)
		if _, ok := c[s]; ok {
			c[s]++
		} else {
			c[s] = 1
		}
	}
	return c
}

// Solve problem Part 1
func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	re := regexp.MustCompile(`([a-z -]+)(\d{3})\[([a-z]+)\]`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		name := strings.ReplaceAll(matches[1], "-", "")
		sector_id, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err)
		}
		checksum := matches[3]
		c := InitCounter(name)
		c_struct := make([]Pairs, len(c))
		i := 0
		for k, v := range c {
			c_struct[i] = Pairs{k, v}
			i++
		}
		sort.Sort(ByOccurences(c_struct))
		var checksum_c string
		for _, k := range c_struct {
			checksum_c += k.key
		}
		checksum_c = checksum_c[0:5]
		if checksum == checksum_c {
			total += sector_id
		}
	}
	return total
}

func LowerAlpha() (map[byte]int, string) {
	p := make([]byte, 26)
	alpha_map := make(map[byte]int, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
		alpha_map[p[i]] = i
	}
	return alpha_map, string(p)
}

// DecryptCaesarShift
func DecryptCaesarShift(id int,
	m map[byte]int,
	lalpha string,
	parts []string) (name string) {
	dparts := make([]string, len(parts))
	for _, word := range parts {
		var dword string
		for i := range word {
			c := word[i]
			pos := m[c]
			new_pos := (pos + id) % 26
			switch {
			case new_pos < 0:
				new_pos += 26
			case new_pos > 26:
				new_pos -= 26
			}
			new_c := lalpha[new_pos]
			dword += string(new_c)
		}
		dparts = append(dparts, dword)
	}
	name = strings.Join(dparts, " ")
	return name
}

// Solve problem Part 2
func Solve2(s *string) (res int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	re := regexp.MustCompile(`([a-z -]+)(\d{3})\[([a-z]+)\]`)
loop:
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		parts := strings.Split(
			strings.TrimSpace(
				strings.ReplaceAll(
					matches[1],
					"-",
					" ")),
			" ")
		sector_id, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err)
		}
		mapPosition, loweralpha := LowerAlpha()
		dname := DecryptCaesarShift(sector_id,
			mapPosition,
			loweralpha,
			parts)
		dname = strings.TrimSpace(dname)
		if strings.Contains(dname, "orth") {
			res = sector_id
			break loop
		}
	}
	return res
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
