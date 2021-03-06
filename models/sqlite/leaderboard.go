package sqlite

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"stubblefield.io/wow-leaderboard-api/models"
)

// PvpLeaderboardStore holds the database object to interact with pvp leaderboard data
type PvpLeaderboardStore struct {
	DB *sqlx.DB
}

// FetchAllByBracket retrieves data for the given bracket
func (store *PvpLeaderboardStore) FetchAllByBracket(pvpBracket string, classes, specs []string, limit, offset int) ([]models.LeaderboardEntry, error) {
	var query string
	var args []interface{}
	var err error

	if len(classes) > 0 && len(specs) > 0 {
		query, args, err = sqlx.In("SELECT * FROM leaderboard WHERE bracket = ? AND character_class IN (?) AND character_spec IN (?)", pvpBracket, classes, specs)
	} else if len(classes) > 0 {
		query, args, err = sqlx.In("SELECT * FROM leaderboard WHERE bracket = ? AND character_class IN (?)", pvpBracket, classes)
	} else {
		query = "SELECT * FROM leaderboard WHERE bracket = ?"
		args = append(args, pvpBracket)
	}
	if err != nil {
		return nil, err
	}

	query += " ORDER BY rank"
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit)
	args = append(args, offset)

	query = store.DB.Rebind(query)
	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populatePvpLeaderboard(rows)
}

// GetClassAndSpecBreakdown gets the class/spec breakdown for a bracket
func (store *PvpLeaderboardStore) GetClassAndSpecBreakdown(bracket string) ([]models.ClassAndSpecBreakdown, error) {
	query := `select l.character_class, l.character_spec, count(*) as spec_count, 1.0 * count(*) / (select count(*) from leaderboard where bracket = ?) * 100 as percent 
	from leaderboard l 
	where bracket = ? and l.character_class  <> '' and l.character_spec <> '' 
	group by l.character_class, l.character_spec`

	rows, err := store.DB.Query(query, bracket, bracket)
	if err != nil {
		return nil, err
	}

	return populateSpecBreakdown(rows)
}

func (store *PvpLeaderboardStore) getSpecsForClasses(classes []string) ([]string, error) {
	query, args, err := sqlx.In("SELECT spec FROM specializations WHERE class IN (?)", classes)
	if err != nil {
		return nil, err
	}

	query = store.DB.Rebind(query)
	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, nil
}

func populatePvpLeaderboard(rows *sql.Rows) ([]models.LeaderboardEntry, error) {
	var leaderboard []models.LeaderboardEntry

	for rows.Next() {
		var e models.LeaderboardEntry

		err := rows.Scan(
			&e.ID,
			&e.Rank,
			&e.Rating,
			&e.CharacterName,
			&e.CharacterID,
			&e.CharacterRealm,
			&e.CharacterRealmSlug,
			&e.CharacterRealmID,
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

func populateSpecBreakdown(rows *sql.Rows) ([]models.ClassAndSpecBreakdown, error) {
	var breakdown []models.ClassAndSpecBreakdown

	for rows.Next() {
		var b models.ClassAndSpecBreakdown

		err := rows.Scan(
			&b.Class,
			&b.Spec,
			&b.SpecCount,
			&b.Percent,
		)
		if err != nil {
			return nil, err
		}

		breakdown = append(breakdown, b)
	}

	return breakdown, nil
}
