package dotabase

import (
	"database/sql"
	"dotaparser/packages/types"
	"fmt"

	"github.com/lib/pq"
)

func (db *dotabase) GetHeroes() (types.Heroes, error) {
	heroes := types.Heroes{}

	query := `SELECT * FROM heroes;`

	rows, err := db.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var hero types.Hero
		var winrates []float64
		var matches []sql.NullInt64

		err := rows.Scan(&hero.Id, &hero.Name, pq.Array(&winrates), pq.Array(&matches))
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}

		for i := 0; i < 5; i++ {
			hero.Winrate[i] = winrates[i]
			hero.Matches[i] = int(matches[i].Int64)
		}

		heroes[hero.Id] = hero
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке строк: %v", err)
	}

	heroes.CalcMeta()

	return heroes, nil
}

func (db *dotabase) InsertHeroes(heroes types.Heroes) error {
	for _, hero := range heroes {
		err := db.InsertHero(hero)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
