package elo

type H2HRecord struct {
	Opponent string
	Won      int
	Lost     int
}

func AddGameToH2HRecords(player string, playerH2HRecords []H2HRecord, game GameLogEntry) []H2HRecord {
	processed := playerH2HRecords
	if player == game.Winner {
		found := false
		for _, record := range processed {
			if record.Opponent == game.Loser {
				found = true
				record.Won++
			}
		}
		if !found {
			record := H2HRecord{
				Opponent: game.Loser,
				Won: 1,
			}
			processed = append(processed, record)
		}
	}
	if player == game.Loser {
		found := false
		for _, record := range processed {
			if record.Opponent == game.Winner {
				found = true
				record.Lost++
			}
		}
		if !found {
			record := H2HRecord{
				Opponent: game.Winner,
				Lost: 1,
			}
			processed = append(processed, record)
		}
	}
	return processed
}

//func findOrCreateH2HRecord(playerH2HRecords []H2HRecord, opponent string) (record H2HRecord, found bool) {
//	found = false
//	newH2HR := H2HRecord{
//		opponent: opponent,
//	}
//	for _, h2hr := range []H2HRecord(playerH2HRecords) {
//		if h2hr.opponent == opponent {
//			return h2hr, true
//		}
//	}
//	return newH2HR, false
//}

// HeadToHeadAll returns all the h2h stats for the specified player.
func (table Table) HeadToHeadAll(player string) ([]H2HRecord, error) {
	res := make([]H2HRecord, 0)
	if _, found := table.Players[player]; !found {
		return res, PlayerDoesNotExist
	}

	for _, game := range table.GameLog.Entries {
		res = AddGameToH2HRecords(player, res, game)
	}
	return res, nil
}
