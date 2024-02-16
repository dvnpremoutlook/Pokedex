package PokeCmds

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/dvnpremoutlook/Pokedex/PokeAPI"
)

func (c *config) checkCache(API string) (value []byte) {

	value, errx := c.cache.Get(API)

	if errx != nil {
		value, errx = PokeAPI.PokeAPI(API)
		if errx != nil {
			fmt.Println(errx)
		}
		c.cache.Add(API, value)
	}

	return value
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	return errors.New("")
}

func (c *config) commandMap() error {

	// API
	API := "https://pokeapi.co/api/v2/location-area"

	if c.next != "" {
		API = c.next
	}

	value := c.checkCache(API)

	locations, errx := PokeAPI.Locations(value)

	if errx != nil {
		fmt.Println(errx)
	}

	c.next = locations.Next
	c.previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf("%s \n", loc.Name)
	}

	return nil
}

func (c *config) commandMapb() error {

	API := c.previous
	if API == "" {
		fmt.Println("Please Try Again : you are on first page")
		return nil
	}

	value := c.checkCache(API)

	locations, errx := PokeAPI.Locations(value)

	if errx != nil {
		fmt.Println(errx)
	}

	c.next = locations.Next
	c.previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf("%s \n", loc.Name)
	}

	return nil
}

func (c *config) explore(location string) error {
	fmt.Println(location)

	API := "https://pokeapi.co/api/v2/location-area/" + location

	value := c.checkCache(API)

	pokemons, errx := PokeAPI.PokemonsEncounters(value)

	if errx != nil {
		fmt.Println(errx)
	}

	c.player.Location = location

	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func (c *config) catch(poke string) error {
	P_API := "https://pokeapi.co/api/v2/pokemon/" + poke
	E_API := P_API + "/encounters"

	p_value := c.checkCache(P_API)
	e_value := c.checkCache(E_API)

	encounters, e_errx := PokeAPI.PokeEncounter(e_value)
	Pokemon, p_errx := PokeAPI.Pokemons(p_value)

	if e_errx != nil || p_errx != nil {
		fmt.Println("Something is wrong with Catch Function Json Retrival")
	}

	for _, loc := range encounters {
		if loc.LocationArea.Name == c.player.Location {
			chance := rand.Intn(1000)
			if Pokemon.BaseExperience < chance {
				c.player.Pokemons[Pokemon.Name] = pokemon{
					name: Pokemon.Name,
					url:  Pokemon.HeldItems[0].Item.URL,
				}

				fmt.Println("You Caught", poke, chance, Pokemon.BaseExperience)

				return nil
			}
			fmt.Println(poke, "Dodged your capture", chance, Pokemon.BaseExperience)

			return nil
		}
	}

	return nil

}

func (c *config) inspect(poke string) error {

	for _, p := range c.player.Pokemons {
		if p.name == poke {
			value := c.checkCache("https://pokeapi.co/api/v2/pokemon/" + p.name)

			pokemonDetails, p_errx := PokeAPI.Pokemons(value)

			if p_errx != nil {
				fmt.Println("Something is wrong with Inspect Function Json Retrival")
			}

			fmt.Println("Name:", pokemonDetails.Name)
			fmt.Println("Weight:", pokemonDetails.Weight)
			fmt.Println("Height:", pokemonDetails.Height)
			fmt.Println("Stats:")

			for _, s := range pokemonDetails.Stats {
				fmt.Println("-", s.Stat.Name, " ", s.BaseStat)
			}
			fmt.Println("Types:")
			for _, t := range pokemonDetails.Types {
				fmt.Println("-", t.Type.Name)
			}

			return nil

		}

	}

	fmt.Println("You dont have", poke)

	return nil

}

func (c *config) pokedex() error {

	for poke := range c.player.Pokemons {
		fmt.Println(poke)
	}

	return nil

}
