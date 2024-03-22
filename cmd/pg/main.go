package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davidmdm/flights/postgresql"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	resources, err := postgresql.RenderChart(os.Args[0], "", nil)
	if err != nil {
		return err
	}
	return json.NewEncoder(os.Stdout).Encode(resources)
}
