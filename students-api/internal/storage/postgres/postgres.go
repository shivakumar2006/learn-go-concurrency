package postgres

import (
	"database/sql"
	"fmt"
	"students-api/internal/config"
	"students-api/internal/types"

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
		return 0, err
	}

	return id, nil
}

func (d *Database) GetStudentById(id int64) (types.Student, error) {
	var student types.Student
	err := d.Db.QueryRow(`
		SELECT id, name, email, age
		FROM students 
		WHERE id = $1
		LIMIT 1
	`, id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with id : %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("Query error : %w", err)
	}

	return student, nil
}

func (d *Database) GetStudents() ([]types.Student, error) {
	rows, err := d.Db.Query(`
		SELECT id, name, email, age FROM students
	`)
	if err != nil {
		return nil, fmt.Errorf("Query error : %w", err)
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		if err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age); err != nil {
			return nil, fmt.Errorf("Error while scanning %w", err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating rows %w", err)
	}

	return students, nil
}

func (d *Database) UpdateStudent(id int64, name, email string, age int) (int64, error) {
	var updatedId int64
	err := d.Db.QueryRow(`
		UPDATE students
		SET name = $1, email = $2, age = $3
		WHERE id = $4
		RETURNING id
	`, name, email, age, id).Scan(&updatedId)

	if err != nil {
		return 0, fmt.Errorf("error while updating student data : %w", err)
	}

	return updatedId, nil
}

func (d *Database) DeleteStudent(id int64) (int64, error) {
	var deleteId int64
	err := d.Db.QueryRow(`
		DELETE FROM students
		WHERE id = $1
		RETURNING id
	`, id).Scan(&deleteId)

	if err != nil {
		return 0, fmt.Errorf("error while deleting the student data : %w", err)
	}

	return deleteId, nil
}
