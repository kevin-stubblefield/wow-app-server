package sqlite

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"stubblefield.io/wow-leaderboard-api/models"
)

type SpecializationStore struct {
	DB *sqlx.DB
}

func (store *SpecializationStore) FetchSpecs() ([]models.Specialization, error) {
	query := "SELECT * FROM specializations"

	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	return populateSpecializations(rows)
}

func populateSpecializations(rows *sql.Rows) ([]models.Specialization, error) {
	var specs []models.Specialization

	for rows.Next() {
		var s models.Specialization

		err := rows.Scan(
			&s.ID,
			&s.Class,
			&s.ClassSlug,
			&s.Spec,
			&s.SpecSlug,
			&s.SpecRole,
		)
		if err != nil {
			return nil, err
		}

		specs = append(specs, s)
	}

	return specs, nil
}
