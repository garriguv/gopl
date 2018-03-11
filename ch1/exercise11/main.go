package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	urls, err := loadUrls("top500.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: error loading urls: %v\n", err)
		os.Exit(1)
	}
	tmp, err := ioutil.TempFile(".", "fetchall")
	defer tmp.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Fprintln(tmp, <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func loadUrls(file string) (urls []string, err error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		urls = append(urls, strings.Split(line, ",")[1])
	}
	return
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
