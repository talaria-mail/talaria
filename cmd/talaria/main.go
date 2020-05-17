package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "talaria",
	Short: "Simple and efficient Email server",
}

func main() {
	cmd.Execute()
}
