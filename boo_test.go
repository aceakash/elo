package elo

import "testing"

func BenchmarkBoo(b *testing.B) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	table.Register("diana")
	table.Register("jon")
	players := [...]string{"bruce", "clark", "diana", "jon"}

	for i := 0; i < b.N; i++ {
		table.AddResult(players[i%4], players[(i+1)%4])
	}
}
