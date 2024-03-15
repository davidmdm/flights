package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/davidmdm/x/xerr"
	"gopkg.in/yaml.v3"
)

//go:embed flights.yaml
var flightData []byte

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var flights []struct {
		Repo    string `yaml:"repo"`
		Version string `yaml:"version"`
	}
	if err := yaml.Unmarshal(flightData, &flights); err != nil {
		return err
	}

	// spinner, stop := ansi.Stdout.Spinner(ansi.SpinnerOptions{})
	// defer stop()

	var errs []error
	for _, flight := range flights {
		// spinner <- fmt.Sprintf("generating flight for %s %s", flight.Repo, flight.Version)

		cmd := exec.Command(
			"go", "run", "github.com/davidmdm/yoke/cmd/helm2go@main",
			"--repo", flight.Repo,
			"--outdir", path.Base(flight.Repo),
		)

		if flight.Version != "" {
			cmd.Args = append(cmd.Args, "--version", flight.Version)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			errs = append(errs, fmt.Errorf("%s:%s: %w", flight.Repo, flight.Version, err))
		}
	}

	return xerr.MultiErrFrom("generating flight(s)", errs...)
}
