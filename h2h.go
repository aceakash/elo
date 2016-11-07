package elo

// H2HRecord is a head-to-head record for a player against a particular opponent.
type H2HRecord struct {
	Opponent string
	Won      int
	Lost     int
}

func registerGameForPlayerH2H(player string, playerH2HRecords []H2HRecord, game GameLogEntry) []H2HRecord {
	latestH2HRecords := playerH2HRecords
	if player == game.Winner {
		found := false
		for i, record := range latestH2HRecords {
			if record.Opponent == game.Loser {
				found = true
				record.Won++
				latestH2HRecords[i] = record
			}
		}
		if !found {
			record := H2HRecord{
				Opponent: game.Loser,
				Won: 1,
			}
			latestH2HRecords = append(latestH2HRecords, record)
		}
	}
	if player == game.Loser {
		found := false
		for i, record := range latestH2HRecords {
			if record.Opponent == game.Winner {
				found = true
				record.Lost++
				latestH2HRecords[i] = record
			}
		}
		if !found {
			record := H2HRecord{
				Opponent: game.Winner,
				Lost: 1,
			}
			latestH2HRecords = append(latestH2HRecords, record)
		}
	}
	return latestH2HRecords
}

// HeadToHeadAll returns all the h2h stats for the specified player.
func (table Table) HeadToHeadAll(player string) ([]H2HRecord, error) {
	res := make([]H2HRecord, 0)
	if _, found := table.Players[player]; !found {
		return res, PlayerDoesNotExist
	}

	for _, game := range table.GameLog.Entries {
		res = registerGameForPlayerH2H(player, res, game)
	}
	return res, nil
}
