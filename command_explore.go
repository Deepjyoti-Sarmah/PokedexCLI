package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 0 {
		return errors.New("No location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon in %s:\n", locationArea.Name)

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func callbackMapb(cfg *config) error {
	// pokeapiClient := pokeapi.NewClient()
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil

}
