// display contents in URL
// add http prefix if missing
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, arg := range os.Args[1:] {
		url := getUrlFromArg(arg, prefix)
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func getUrlFromArg(arg string, prefix string) string {
	if strings.HasPrefix(arg, prefix) {
		return arg
	} else {
		return prefix + arg
	}
}
