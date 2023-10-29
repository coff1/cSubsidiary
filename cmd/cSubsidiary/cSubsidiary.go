package main

import (
	"github.com/coff1/cSubsidiary/internal/runner"
)

func main() {
	options := runner.ParseOptions()

	runner.RunEnumeration(options)
}

