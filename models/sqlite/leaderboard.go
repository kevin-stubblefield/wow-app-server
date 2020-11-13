package sqlite

import (
	"database/sql"

	"stubblefield.io/wow-leaderboard-api/models"
)

type PvpLeaderboardStore struct {
	DB *sql.DB
}

func (store *PvpLeaderboardStore) FetchAllByBracket(pvpBracket, classList, specList string) ([]models.LeaderboardEntry, error) {
	// var stmt string
	// var rows *sql.Rows
	// var args []interface{}
	// classes := strings.Split(classList, ",")
	// specs := strings.Split(specList, ",")

	// if len(specs) > 0 {
	// 	stmt = "SELECT * FROM leaderboard WHERE bracket = ? AND character_class IN ("
	// 	stmt += strings.Repeat("?, ", len(classes)-1) + "?"
	// 	stmt += ") AND character_spec IN ("
	// 	stmt += strings.Repeat("?, ", len(specs)-1) + "?"
	// 	stmt += ")"
	// } else if len(classes) > 0 {
	// 	stmt = "SELECT * FROM leaderboard WHERE bracket = ? AND character_class IN ("
	// 	stmt += strings.Repeat("?, ", len(classes)-1) + "?"
	// 	stmt += ")"
	// }

	stmt := "SELECT * FROM leaderboard WHERE bracket = ?"
	stmt += " LIMIT 20 OFFSET 0"
	rows, err := store.DB.Query(stmt, pvpBracket)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	leaderboard, err := populatePvpLeaderboard(rows)

	return leaderboard, nil
}

func populatePvpLeaderboard(rows *sql.Rows) ([]models.LeaderboardEntry, error) {
	var leaderboard []models.LeaderboardEntry

	for rows.Next() {
		var e models.LeaderboardEntry

		err := rows.Scan(
			&e.Id,
			&e.Rank,
			&e.Rating,
			&e.CharacterName,
			&e.CharacterId,
			&e.CharacterRealmSlug,
			&e.CharacterRealmId,
			&e.CharacterFaction,
			&e.CharacterRace,
			&e.CharacterClass,
			&e.CharacterSpec,
			&e.GamesPlayed,
			&e.GamesWon,
			&e.GamesLost,
			&e.Bracket,
		)
		if err != nil {
			return nil, err
		}

		leaderboard = append(leaderboard, e)
	}

	return leaderboard, nil
}
