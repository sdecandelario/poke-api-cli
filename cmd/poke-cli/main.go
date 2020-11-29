package main

import (
	pokemonscli "github.com/sdecandelario/poke-api-cli/internal"
	"github.com/sdecandelario/poke-api-cli/internal/cli"
	"github.com/sdecandelario/poke-api-cli/internal/storage/pokeapi"
	"github.com/spf13/cobra"
)

func main() {

	var repo pokemonscli.PokemonRepo
	repo = pokeapi.NewPokeAPIRepository()

	rootCmd := &cobra.Command{Use: "pokemon-cli"}
	rootCmd.AddCommand(cli.InitPokemonCmd(repo))
	rootCmd.Execute()
}
