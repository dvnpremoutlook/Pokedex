package PokeCmds

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/dvnpremoutlook/Pokedex/PokeCache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type pokemon struct {
	name string
	url  string
}

type user struct {
	name     string
	Pokemons map[string]pokemon
	Location string
}

type config struct {
	player   user
	cache    pokecache.Cache
	next     string
	previous string
}

func (c *config) cliCommands(command, arguments string) cliCommand {

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display first 20 locations",
			callback:    func() error { return c.commandMap() },
		},
		"mapb": {
			name:        "mapb",
			description: "Display Previous 20 locations",
			callback:    func() error { return c.commandMapb() },
		},
		"explore": {
			name:        "explore <location>",
			description: "Given Location Display all Pokemons",
			callback:    func() error { return c.explore(arguments) },
		},
		"catch": {
			name:        "explore <location>",
			description: "Given Location Display all Pokemons",
			callback:    func() error { return c.catch(arguments) },
		},
		"inspect": {
			name:        "explore <location>",
			description: "Given Location Display all Pokemons",
			callback:    func() error { return c.inspect(arguments) },
		},
		"Pokedex": {
			name:        "explore <location>",
			description: "Given Location Display all Pokemons",
			callback:    func() error { return c.pokedex() },
		},
	}

	return commands[command]
}

func PokeCmds(c config) {
	scanner := bufio.NewScanner(os.Stdin)

	interval := 5
	c.cache = pokecache.NewCache(interval)
	c.player.Pokemons = make(map[string]pokemon)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		arguments := strings.Split(text, " ")

		if len(arguments) == 1 {
			arguments = append(arguments, " ")
		}

		command := c.cliCommands(arguments[0], arguments[1])
		if command.name == "" {
			fmt.Println("Command Not Found")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
			break
		}

	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

}

func Config(next, previous string) config {
	return config{}
}
