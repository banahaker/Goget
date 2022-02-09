package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getData(url string) (string, time.Duration, int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data), duration, resp.StatusCode
}

func writeFile(file string, content string) {
	os.Mkdir("data", os.ModePerm)
	f, err := os.Create(fmt.Sprintf("data\\%s", file))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func main() {
	COLORS := map[string]string{
		"DEFAULT": "\033[0m",
		"RED":     "\033[31m",
		"GREEN":   "\033[32m",
		"YELLOW":  "\033[33m",
		"BLUE":    "\033[34m",
		"PURPLE":  "\033[35m",
		"WHITE":   "\033[97m",
	}

	url := flag.String("u", "", "Url to request")
	location := flag.String("loc", "", "location to place response content")
	flag.Parse()
	if *url == "" {
		fmt.Println("Please Type the URL to request")
	}
	data, time, status := getData(*url)
	if *location != "" {
		writeFile(*location, data)
	}
	if status == 200 {
		fmt.Println("Status Code: ", COLORS["GREEN"], status, COLORS["DEFAULT"])
	} else {
		fmt.Println("Status Code: ", COLORS["RED"], status, COLORS["DEFAULT"])
	}
	fmt.Println("Time       : ", COLORS["YELLOW"], time, COLORS["DEFAULT"])
	fmt.Println("Response   : ", data)
}
