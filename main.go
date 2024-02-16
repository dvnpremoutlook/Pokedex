package main

import (
	"github.com/dvnpremoutlook/Pokedex/PokeCmds"
)

func main() {

	c := PokeCmds.Config(
		"",
		"",
	)

	PokeCmds.PokeCmds(c)

}
