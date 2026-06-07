package postgres

import (
	"database/sql"
	"students-api/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Database, error) {
	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &Database{
		Db: db,
	}, nil
}

func (d *Database) CreateStudent(name string, email string, age int) (int64, error) {
	var id int64

	err := d.Db.QueryRow(`
		INSERT INTO students (name, email, age)
		VALUES ($1, $2, $3)
		RETURNING id
	`, name, email, age).Scan(&id)

	if err != nil {
		return 0, nil
	}

	return id, nil
}
