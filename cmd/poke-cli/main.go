package main

import (
	"github.com/sdecandelario/poke-api-cli/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "pokemon-cli"}
	rootCmd.AddCommand(cli.InitPokemonCmd())
	rootCmd.Execute()
}
