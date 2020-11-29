package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pokemonscli "github.com/sdecandelario/poke-api-cli/internal"
)

const (
	pokemonEndpoint = "/pokemon/"
	baseURL         = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
	url string
}

// NewPokeAPIRepository fetch pokemon data from poke api
func NewPokeAPIRepository() pokemonscli.PokemonRepo {
	return &pokemonRepo{url: baseURL}
}

func (p *pokemonRepo) GetPokemonByName(name string) (pokemon pokemonscli.Pokemon, err error) {
	response, _ := http.Get(fmt.Sprintf("%v%v%s", p.url, pokemonEndpoint, name))

	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &pokemon)

	return
}

func (p *pokemonRepo) GetPokemonByID(id int) (pokemon pokemonscli.Pokemon, err error) {
	response, _ := http.Get(fmt.Sprintf("%v%v%d", p.url, pokemonEndpoint, id))

	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &pokemon)

	return
}
