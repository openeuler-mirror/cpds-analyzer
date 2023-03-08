package main

import (
	analyzer "cpds/cpds-analyzer/pkg/cpds-analyzer/server"
)

func main() {
	cmd := analyzer.NewCommand()

	if err := cmd.Execute(); err != nil {

	}
}
