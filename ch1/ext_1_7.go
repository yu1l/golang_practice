// use io.Copy(dst, src) instead of ioutil.ReadAll
// also, make sure detect io.Copy error
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err1 := io.Copy(os.Stdout, response.Body); err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err1)
			os.Exit(1)
		}
		response.Body.Close()
	}
}
