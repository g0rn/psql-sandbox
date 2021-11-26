package psqlpq

import (
	"database/sql"

	"github.com/g0rn/psql-sandbox/pkg/entities"
	_ "github.com/lib/pq"
)

func NewClient() (*sql.DB, error) {
	connStr := "postgres://g0rn:h0wbvky2@192.168.1.226/test"
	return sql.Open("postgres", connStr)
}

func InsertPerson(db *sql.DB, person *entities.Person) error {
	insertQuery := `INSERT INTO persons(first_name, last_name, age, email, role, sex) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(insertQuery, person.FirstName, person.LastName, person.Age, person.Email, person.Role, person.Sex)
	return err
}

func GetPerson(db *sql.DB, id string) (*entities.Person, error) {
	query := `SELECT * FROM persons WHERE id = $1`
	row := db.QueryRow(query, id)
	person := entities.Person{}
	err := row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age, &person.Email, &person.Role, &person.Sex)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func ListPersons(db *sql.DB) ([]entities.Person, error) {
	query := `SELECT * FROM persons`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	persons := make([]entities.Person, 0)
	for rows.Next() {
		person := entities.Person{}

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age, &person.Email, &person.Role, &person.Sex)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}
	return persons, nil
}

func DeletePerson(db *sql.DB, id string) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

func DeletePersons(db *sql.DB) error {
	query := `DELETE FROM persons;`
	_, err := db.Exec(query)
	return err
}
