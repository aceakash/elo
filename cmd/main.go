package main

import "github.com/aceakash/elo"
import "fmt"

func main() {
	table := elo.NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	table.AddResult("bruce", "clark")
	fmt.Println(table)
}
