package dotabase

import (
	"dotaparser/packages/types"
	"fmt"

	"github.com/lib/pq"
)

func (db *dotabase) InsertMatch(match *types.Match, pro bool) error {
	query := `
	INSERT INTO matches (id, radiant, dire, metaDif, pro)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (id) DO UPDATE SET
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

	_, err := db.db.Exec(query, match.Id, pq.Array(radiant), pq.Array(dire), match.MetaDif, pro)
	if err != nil {
		return fmt.Errorf("data insert error: %v", err)
	}
	return nil
}
