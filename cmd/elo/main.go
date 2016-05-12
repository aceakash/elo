package main

import (
	"fmt"
	"os"
)

func main() {
	if !isCommandValid(os.Args) {
		printUsage()
		os.Exit(1)
	}
}

func isCommandValid(args []string) bool {
	if isValidRegisterCommand(args) {
		return true
	}
	return false
}

func isValidRegisterCommand(args []string) bool {

}

func printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("elo register <nickname> -- register a new player")
	fmt.Println("elo result <winner> <loser> -- add results of a match")
	fmt.Println("elo ratings -- see ratings for all players")
	fmt.Println("elo ratings <count> -- see ratings for top <count> players")
	fmt.Println("elo ratings <player> -- see ratings for <player>")
	fmt.Println("elo h2h <player1> <player2> -- see win/loss numbers for matches played between <player1> and <player2>")
	fmt.Println("elo history <player> -- see matches played by <player>")
	fmt.Println("elo history <player1> <player2> -- see matches played by <player1> against <player2>")
}

/*
register akash
register yash
result yash akash
result akash yash
register reshma
result reshma yash
result reshma akash
ratings
ratings 10
ratings reshma
ratings akash
h2h reshma akash
history reshma akash
history reshma

register <player nickname>
	should add player name to db, along with when registered
	player
		Nickname (unique)
		Name
		JoinedOn
	should complain if already registered
	should only accept lowercase chars

result <winner nickname> <loser nickname>
	should add match to db
	match
		Winner
		Loser
		AddedOn
	should error if either players not found in players db
	should error if winner is same as loser

ratings
	calculate ratings for all players on the fly and show in inverse order of ratings

ratings <count>
	show only top <count> ratings

h2h <player1> <player2>
	show head to head stats for player1 and player2

history <player1> <player2>
	show matches where player1 played player2

history <player nickname>
	show matches played by player


*/
