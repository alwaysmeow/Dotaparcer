package dotabase

import (
	"database/sql"
	"dotaparser/packages/types"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func (db *dotabase) InsertMatch(match *types.Match, pro bool) error {
	if db == nil {
		return fmt.Errorf("database instance is nil")
	}
	if match == nil {
		return fmt.Errorf("match is nil")
	}

	query := `
	INSERT INTO matches (id, rWon, radiant, dire, metaDif, pro)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id) DO UPDATE SET
		rWon = EXCLUDED.rWon,
		radiant = EXCLUDED.radiant,
		dire = EXCLUDED.dire,
		metaDif = EXCLUDED.metaDif,
		pro = EXCLUDED.pro
	`

	var radiant []int
	var dire []int

	for i := 0; i < 5; i++ {
		radiant = append(radiant, match.Radiant.Heroes[i].Id)
		dire = append(dire, match.Dire.Heroes[i].Id)
	}

	rWon := match.Winner == types.Radiant

	_, err := db.db.Exec(query, match.Id, rWon, pq.Array(radiant), pq.Array(dire), match.MetaDif, pro)
	if err != nil {
		return fmt.Errorf("data insert error: %v", err)
	}
	return nil
}

func (db *dotabase) MatchExist(id int) bool {
	query := fmt.Sprintf("SELECT id, rWon, radiant, dire, metaDif FROM matches WHERE id = %d;", id)

	match := types.Match{}

	rWin := false
	var radiant []sql.NullInt64
	var dire []sql.NullInt64

	row := db.db.QueryRow(query)
	err := row.Scan(&match.Id, &rWin, pq.Array(&radiant), pq.Array(&dire), &match.MetaDif)

	if err == sql.ErrNoRows {
		// No match found
		return false
	} else if err != nil {
		// Some other error occurred
		log.Printf("QueryRow error: %v", err)
		return false
	}

	return true
}
