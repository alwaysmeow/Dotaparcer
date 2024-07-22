package main

import (
	"database/sql"
	"dotaparser/types"
	"fmt"

	"github.com/lib/pq"
)

func main() {
	//heroes, _ := types.ParseHeroes()
	db := getdb()
	dbinit(db)
	//insertHero(db, (*heroes)[0])
	hero, _ := getHero(db, 1)
	hero.Log()
	db.Close()
}

func getdb() *sql.DB {
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

func dbinit(db *sql.DB) error {
	// SQL-запрос для создания таблицы
	query := `
    CREATE TABLE IF NOT EXISTS heroes (
        id SERIAL PRIMARY KEY,
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

func insertHero(db *sql.DB, hero types.Hero) error {
	query := `
    INSERT INTO heroes (name, winrates, matches)
    VALUES ($1, $2, $3)
    `
	fmt.Println("Executing query:", query)
	fmt.Printf("Parameters: name=%s, winrates=%v, matches=%v\n", hero.Name, hero.Winrate, hero.Matches)
	_, err := db.Exec(query, hero.Name, pq.Array(hero.Winrate), pq.Array(hero.Matches))
	if err != nil {
		return fmt.Errorf("ошибка при вставке данных: %v", err)
	}
	return nil
}

func getHero(db *sql.DB, id int) (types.Hero, error) {
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
