package cli

import (
	"encoding/csv"
	"os"
	"strconv"

	pokemonscli "github.com/sdecandelario/poke-api-cli/internal"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"
const nameFlag = "name"
const fileNameFlag = "filename"

// InitPokemonCmd initialize pokemon command
func InitPokemonCmd(repository pokemonscli.PokemonRepo) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemons",
		Short: "Print data about pokemons",
		Run:   runPokemonsFn(repository),
	}

	pokemonCmd.Flags().IntP(idFlag, "i", 0, "ID of the pokemon")
	pokemonCmd.Flags().StringP(nameFlag, "n", "", "Name of the pokemon")
	pokemonCmd.Flags().StringP(fileNameFlag, "f", "pokemon.csv", "Name of csv file")

	return pokemonCmd
}

func runPokemonsFn(repository pokemonscli.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt(idFlag)
		name, _ := cmd.Flags().GetString(nameFlag)
		filename, _ := cmd.Flags().GetString(fileNameFlag)

		var pokemon pokemonscli.Pokemon

		if name != "" {
			pokemon, _ = repository.GetPokemonByName(name)
		} else {
			pokemon, _ = repository.GetPokemonByID(id)
		}

		os.MkdirAll("data", 0700)
		csvFile, _ := os.Create("data/" + filename)
		csvwriter := csv.NewWriter(csvFile)
		csvInfo := []string{
			strconv.Itoa(pokemon.ID), pokemon.Name, strconv.Itoa(pokemon.Order),
		}
		_ = csvwriter.Write(csvInfo)
		csvwriter.Flush()
		csvFile.Close()
	}
}
