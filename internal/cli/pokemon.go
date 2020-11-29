package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"
const nameFlag = "name"

// InitPokemonCmd initialize pokemon command
func InitPokemonCmd() *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemons",
		Short: "Print data about pokemons",
		Run:   runPokemonsFn(),
	}

	pokemonCmd.Flags().StringP(idFlag, "i", "", "ID of the pokemon")
	pokemonCmd.Flags().StringP(nameFlag, "n", "", "Name of the pokemon")

	return pokemonCmd
}

// Pokemon representation of a Pokemon
type Pokemon struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Order int    `json:"order"`
}

func runPokemonsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		name, _ := cmd.Flags().GetString(nameFlag)

		if id != "" || name != "" {
			var pokemon Pokemon
			var url string
			if id != "" {
				url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
			} else {
				url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
			}
			response, _ := http.Get(url)
			body, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &pokemon)
			os.MkdirAll("data", 0700)
			csvFile, _ := os.Create("data/pokemon.csv")
			csvwriter := csv.NewWriter(csvFile)
			csvInfo := []string{
				strconv.Itoa(pokemon.ID), pokemon.Name, strconv.Itoa(pokemon.Order),
			}
			_ = csvwriter.Write(csvInfo)
			csvwriter.Flush()
			csvFile.Close()
		} else {
			fmt.Println("no number")
		}
	}
}
