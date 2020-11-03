package sqlite

import (
	"database/sql"

	"stubblefield.io/wow-leaderboard-api/models"
)

type PvpLeaderboardStore struct {
	DB *sql.DB
}

func (store *PvpLeaderboardStore) FetchAllByBracket(pvpBracket string) ([]models.LeaderboardEntry, error) {
	stmt := "SELECT * FROM leaderboard WHERE bracket = ?"
	rows, err := store.DB.Query(stmt, pvpBracket)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaderboard []models.LeaderboardEntry

	for rows.Next() {
		var e models.LeaderboardEntry

		err = rows.Scan(
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
