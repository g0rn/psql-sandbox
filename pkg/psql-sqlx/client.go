package psqlsqlx

import (
	"github.com/g0rn/psql-sandbox/pkg/entities"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewClient() (*sqlx.DB, error) {
	connStr := "postgres://g0rn:h0wbvky2@192.168.1.226/test"
	return sqlx.Open("postgres", connStr)
}

func InsertPerson(db *sqlx.DB, person *entities.Person) error {
	query := `INSERT INTO persons(first_name, last_name, age, email, role, sex) VALUES (:first_name, :last_name, :age, :email, :role, :sex)`
	_, err := db.NamedExec(query, person)
	return err
}

func GetPerson(db *sqlx.DB, id string) (*entities.Person, error) {
	query := `SELECT * FROM persons WHERE id = $1`
	person := entities.Person{}
	err := db.Get(&person, query, id)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func ListPersons(db *sqlx.DB) ([]entities.Person, error) {
	query := `SELECT * FROM persons`
	persons := []entities.Person{}
	err := db.Select(&persons, query)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func DeletePerson(db *sqlx.DB, id string) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

func DeletePersons(db *sqlx.DB) error {
	query := `DELETE FROM persons;`
	_, err := db.Exec(query)
	return err
}
