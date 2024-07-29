package dotabase

import (
	"database/sql"
	"dotaparser/packages/types"
	"fmt"

	"github.com/lib/pq"
)

func GetHero(db *sql.DB, id int) (types.Hero, error) {
	var hero types.Hero
	query := `
    SELECT id, name, winrates, matches
    FROM heroes
    WHERE id = $1
    `
	row := db.QueryRow(query, id)
	var winrates []float64
	var matches []sql.NullInt64
	err := row.Scan(&hero.Id, &hero.Name, pq.Array(&winrates), pq.Array(&matches))
	if err != nil {
		return hero, fmt.Errorf("ошибка при извлечении данных: %v", err)
	}

	for i := 0; i < 5; i++ {
		hero.Winrate[i] = winrates[i]
		if matches[i].Valid {
			hero.Matches[i] = int(matches[i].Int64)
		} else {
			hero.Matches[i] = 0
		}
	}

	fmt.Println(hero)
	return hero, nil
}

func InsertHero(db *sql.DB, hero types.Hero) error {
	query := `
    INSERT INTO heroes (id, name, winrates, matches)
    VALUES ($1, $2, $3, $4)
    ON CONFLICT (id) DO UPDATE SET
        name = EXCLUDED.name,
        winrates = EXCLUDED.winrates,
        matches = EXCLUDED.matches
    `
	_, err := db.Exec(query, hero.Id, hero.Name, pq.Array(hero.Winrate), pq.Array(hero.Matches))
	if err != nil {
		return fmt.Errorf("data insert error: %v", err)
	}
	return nil
}
