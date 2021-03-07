package main

import "github.com/spf13/cobra"

func main() {
	cmd := &cobra.Command{
		Use:   "talaria",
		Short: "Email simplified",
	}

	cmd.AddCommand(NewServeCmd())

	cmd.Execute()
}
