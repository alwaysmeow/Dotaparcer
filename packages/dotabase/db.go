package dotabase

import (
	"database/sql"
	"dotaparser/packages/types"
	"fmt"

	"github.com/lib/pq"
)

func GetDB() *sql.DB {
	connStr := "user=meowmeow dbname=dotabase sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка открытия соединения:", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ошибка при пинге базы данных:", err)
		return nil
	}

	return db
}

func DBinit(db *sql.DB) error {
	// SQL-запрос для создания таблицы
	query := `
    CREATE TABLE IF NOT EXISTS heroes (
    	id INT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        winrates FLOAT8[5],
        matches INT4[5]
    );
    `

	// Выполняем запрос
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы: %v", err)
	}

	return nil
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

func InsertHeroes(db *sql.DB, heroes types.Heroes) error {
	for _, hero := range heroes {
		err := InsertHero(db, hero)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
