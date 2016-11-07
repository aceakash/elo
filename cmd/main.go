package main

import (
	"fmt"
	"github.com/aceakash/elo"
	"strings"
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
		fmt.Println("(5) Recalculate ratings from game log")
		fmt.Println("(6) Quit")
		fmt.Print("\nChoose an option: ")
		option := 0
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Err", err)
		}
		if option < 1 || option > 7 {
			fmt.Println("Oops. That's an invalid option - valid entries are 1 - 6.")
			continue prompt
		}
		switch option {
		case 7:
			fmt.Print()
			break prompt
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
				fmt.Println("\n!!!!!!!!!! Oops - that's not a registered player.")
				continue prompt
			}
			fmt.Print("\nWho lost? ")
			fmt.Scanln(&loser)
			if _, exists := table.Players[loser]; !exists {
				fmt.Println("\n!!!!!!!!!! Oops - that's not a registered player.")
				continue prompt
			}
			table.AddResult(winner, loser)

			store.Save(table)
			fmt.Println("\nResult saved")
		case 4:
			fmt.Println("\n\n")
			for _, gle := range table.GameLog.Entries {
				created := gle.Created.Format("_2 Jan 2006")
				fmt.Printf("[%s] %17s (%d -> %d)   defeated %17s (%d -> %d)\n", created, gle.Winner, gle.WinnerChange.Before, gle.WinnerChange.After, gle.Loser, gle.LoserChange.Before, gle.LoserChange.After)
			}
			fmt.Println("\n\n")
		case 5:
			fmt.Println("\n\nAre you sure you want to recreate the ratings table from the log? (y/n): ")
			var answer string
			fmt.Scanln(&answer)
			answer = strings.ToLower(answer)
			if answer != "y" && answer != "n" {
				fmt.Println("!!!!!!!!!! Oops - please only answer in y or n")
				continue prompt
			}
			if answer == "n" {
				fmt.Println("No worries, table left untouched.")
				continue prompt
			}
			err := table.RecalculateRatingsFromLog()
			if err != nil {
				fmt.Println("!!!!!!!!!! Oops - something went wrong! Here's the error:")
				fmt.Print(err)
				continue prompt
			}
			store.Save(table) // todo: handle errors
			fmt.Println("\nAll done, ratings table is now in sync with the game log.")

		case 6:
			var player string
			fmt.Scanln(&player)
			normalisedPlayer := player
			playerH2HRecords, err := table.HeadToHeadAll(normalisedPlayer)
			if err != nil {
				if err == elo.PlayerDoesNotExist {
					//fmt.Fprintf(w, "%s does not seem to be registered", against)
					return
				}
			}
			fmt.Println("H2H for ", player, playerH2HRecords)
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
		fmt.Println("!!!!!!!!!! Oops - that player has already been registered\n")
		return
	}
	fmt.Printf("\n%s has been registered\n", name)
}

func printEloTable(table elo.Table) {
	if len(table.Players) == 0 {
		fmt.Println("\nNo players registered!")
	}
	for _, player := range table.GetPlayersSortedByRating() {
		fmt.Printf("\n%25s (%d) - Played %2d, Won %2d, Lost %2d", player.Name, player.Rating, player.Played, player.Won, player.Lost)
	}
	fmt.Println("")
}
