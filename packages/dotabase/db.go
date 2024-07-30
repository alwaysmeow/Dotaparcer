package dotabase

import (
	"database/sql"
	"fmt"
)

type dotabase struct {
	db *sql.DB
}

func GetDB() *dotabase {
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

	return &dotabase{db: db}
}

func (db *dotabase) Init() error {
	// SQL-запрос для создания таблицы
	query := `
    CREATE TABLE IF NOT EXISTS heroes (
    	id INT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        winrates FLOAT8[5],
        matches INT4[5]
    );
	CREATE TABLE IF NOT EXISTS matches (
        id INT PRIMARY KEY,
		radiant INT4[5],
		dire INT4[5],
		metaDif FLOAT8,
		pro BOOLEAN
    );
    `

	// Выполняем запрос
	_, err := db.db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы: %v", err)
	}

	return nil
}

func (db *dotabase) Close() {
	db.db.Close()
}
