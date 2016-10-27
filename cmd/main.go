package main

import (
	"fmt"
	"github.com/aceakash/elo"
)

func main() {
	store := elo.JsonFileTableStore{
		Filepath: "eloTable.json",
	}
	table, err := store.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Table loaded:", table)

	prompt:
	for {
		fmt.Println("\nWhat would you like to do?\n")
		fmt.Println("(1) View the ratings table")
		fmt.Println("(2) Register a new player")
		fmt.Println("(3) Record outcome of a game")
		fmt.Println("(4) Quit")
		fmt.Print("\nChoose an option: ")
		option := 0
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Err", err)
		}
		if option < 1 || option > 4 {
			fmt.Println("Oops. That's an invalid option - valid entries are 1 - 4.")
			continue
		}
		switch option {
		case 4:
			fmt.Print()
			break prompt;
		case 1:
			printEloTable(table)
		case 2:
			registerNewPlayer(&table)
			store.Save(table)
		case 3:
			var winner, loser string
			fmt.Print("\n\nWho won? ")
			fmt.Scanln(&winner)
			fmt.Print("\nWho lost? ")
			fmt.Scanln(&loser)
			table.AddResult(winner, loser)
			store.Save(table)
			fmt.Println("\nResult saved")
		default:
			fmt.Println("You chose ", option)
		}
	}
}
func registerNewPlayer(table *elo.Table) {
	fmt.Print("\nEnter player name: ")
	name := ""
	fmt.Scanln(&name)
	table.Register(name)
	fmt.Printf("\n%s has been registered\n", name)
}

func printEloTable(table elo.Table) {
	if len(table.Players) == 0 {
		fmt.Println("\nNo players registered!")
	}
	for name, player := range table.Players {
		fmt.Printf("\n%s (%d) - Played %d, Won %d, Lost %d", name, player.Rating, player.Played, player.Won, player.Lost)
	}
	fmt.Println("")
}
