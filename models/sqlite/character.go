package sqlite

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"stubblefield.io/wow-leaderboard-api/models"
)

// CharacterStore holds the database object to interact with character data
type CharacterStore struct {
	DB *sqlx.DB
}

// Fetch retrieves one character given the realm and character name
func (store *CharacterStore) Fetch(realmSlug, characterName string) (*models.Character, error) {
	row := store.DB.QueryRow("SELECT character_name, character_realm, character_realm_slug FROM leaderboard WHERE character_realm_slug = ? AND character_name = ?", realmSlug, characterName)
	character, err := populateCharacter(row)
	if err != nil {
		return nil, err
	}

	rows, err := store.DB.Query("SELECT * FROM equipment WHERE character_realm_slug = ? AND character_name = ?", realmSlug, characterName)
	if err != nil {
		return nil, err
	}

	equipment, err := populateEquipment(rows)
	if err != nil {
		return nil, err
	}

	character.Equipped = equipment

	return character, nil
}

func populateCharacter(row *sql.Row) (*models.Character, error) {
	var c models.Character

	err := row.Scan(
		&c.Name,
		&c.Realm,
		&c.RealmSlug,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func populateEquipment(rows *sql.Rows) ([]models.Equipment, error) {
	var equipped []models.Equipment

	for rows.Next() {
		var e models.Equipment

		err := rows.Scan(
			&e.ID,
			&e.CharacterRealmSlug,
			&e.CharacterName,
			&e.ItemID,
			&e.ItemSlot,
			&e.ItemName,
		)
		if err != nil {
			return nil, err
		}

		equipped = append(equipped, e)
	}

	return equipped, nil
}
