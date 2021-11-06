package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	flag.IntVar(&flagday, "day", 1, "Day of AOC")
	flag.IntVar(&flagyear, "year", 2020, "Year of AOC")

	flag.Parse()

	data := GetInput(flagyear, flagday)
	err := os.WriteFile("input.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
