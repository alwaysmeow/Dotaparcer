package dotabase

import (
	"database/sql"
	"fmt"
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
