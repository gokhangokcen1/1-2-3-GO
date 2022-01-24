package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	f, _ := os.Open("/home/runner/scraper/buster/wordlist.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		directory(line)

	}
	fmt.Println("*****Done*****")
}

func directory(dir string) {
	url := "https://tryhackme.com/" + dir
	res, _ := http.Get(url)

	if res.StatusCode == 200 {
		fmt.Printf("/%v\n", dir)

	}
}
