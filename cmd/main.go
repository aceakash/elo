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

	prompt:
	for {
		fmt.Println("\nWhat would you like to do?\n")
		fmt.Println("(1) View the ratings table")
		fmt.Println("(2) Register a new player")
		fmt.Println("(3) Record outcome of a game")
		fmt.Println("(4) View all recorded games")
		fmt.Println("(5) Quit")
		fmt.Print("\nChoose an option: ")
		option := 0
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Err", err)
		}
		if option < 1 || option > 5 {
			fmt.Println("Oops. That's an invalid option - valid entries are 1 - 5.")
			continue prompt
		}
		switch option {
		case 5:
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
			if _, exists := table.Players[winner]; !exists {
				fmt.Println("\nOops - that's not a registered player.")
				continue prompt
			}
			fmt.Print("\nWho lost? ")
			fmt.Scanln(&loser)
			if _, exists := table.Players[winner]; !exists {
				fmt.Println("\nOops - that's not a registered player.")
				continue prompt
			}
			table.AddResult(winner, loser)

			store.Save(table)
			fmt.Println("\nResult saved")
		case 4:
			fmt.Println("\n\n")
			for _, gle := range table.GameLog.Entries {
				fmt.Printf("[%s] %s defeated %s\n", gle.Created, gle.Winner, gle.Loser)
			}
			fmt.Println("\n\n")
		default:
			fmt.Println("You chose ", option)
		}
	}
}
func registerNewPlayer(table *elo.Table) {
	fmt.Print("\nEnter player name: ")
	name := ""
	fmt.Scanln(&name)
	err := table.Register(name)
	if err == elo.PlayerAlreadyExists {
		fmt.Println("Oops. That player has already been registered\n")
		return
	}
	fmt.Printf("\n%s has been registered\n", name)
}

func printEloTable(table elo.Table) {
	if len(table.Players) == 0 {
		fmt.Println("\nNo players registered!")
	}
	for _, player := range table.GetPlayersSortedByRating() {
		fmt.Printf("\n%10s (%d) - Played %2d, Won %2d, Lost %2d", player.Name, player.Rating, player.Played, player.Won, player.Lost)
	}
	fmt.Println("")
}

