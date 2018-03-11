package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, cpErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if cpErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, cpErr)
			os.Exit(1)
		}
	}
}
