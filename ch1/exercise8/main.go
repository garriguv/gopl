package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { // FIXME: What about all the other schemes :)
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, cpErr := io.Copy(os.Stdout, resp.Body)
		if cpErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, cpErr)
			os.Exit(1)
		}
	}
}
