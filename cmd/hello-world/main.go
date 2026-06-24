package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebmaspd/hello-spec-it/internal/greeter"
)

func main() {
	if len(os.Args) < 2 || strings.TrimSpace(os.Args[1]) == "" {
		fmt.Fprintln(os.Stderr, "Usage: hello-world <name>")
		os.Exit(1)
	}
	fmt.Println(greeter.Greet(os.Args[1]))
}
