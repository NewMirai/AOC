package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"
)

// GetInput Make a request to AOC website
func GetInput(year int, day int) []byte {
	url := fmt.Sprint("https://adventofcode.com/", year, "/day/", day, "/input")
	session_number := os.Getenv("AOC_SESSION")
	client := &http.Client{}
	cookie := &http.Cookie{
		Name:  "session",
		Value: session_number,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	return body
}

// main function
func main() {
	var flagday int
	var flagyear int
	var filename string

	flag.IntVar(&flagday, "day", 1, "Day of AOC")
	flag.IntVar(&flagyear, "year", 2020, "Year of AOC")

	flag.Parse()

	data := GetInput(flagyear, flagday)
	if flagday < 10 {
		filename = fmt.Sprint("./", flagyear, "/day0", flagday, "/input.txt")
	} else {
		filename = fmt.Sprint("./", flagyear, "/day", flagday, "/input.txt")
	}
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
