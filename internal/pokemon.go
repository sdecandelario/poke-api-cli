package pokemonscli

// PokemonRepo definiton of methods to access a data pokemon
type PokemonRepo interface {
	GetPokemonByName(name string) (Pokemon, error)
	GetPokemonByID(id int) (Pokemon, error)
}

// Pokemon representation of a Pokemon
type Pokemon struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Order int    `json:"order"`
}
