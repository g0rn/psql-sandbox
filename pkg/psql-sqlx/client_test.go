package psqlsqlx_test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/g0rn/psql-sandbox/pkg/entities"
	psqlsqlx "github.com/g0rn/psql-sandbox/pkg/psql-sqlx"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = psqlsqlx.NewClient()
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkSQLXInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		psqlsqlx.InsertPerson(db, &entities.Person{
			FirstName: "FirstName_Test",
			LastName:  "LastName_Test",
			Email:     "test@test.org",
			Age:       23,
			Role:      "test-user",
			Sex:       "Female",
		})
	}

	psqlsqlx.DeletePersons(db)
}

func BenchmarkSQLXList(b *testing.B) {
	for i := 0; i < 1000; i++ {
		psqlsqlx.InsertPerson(db, &entities.Person{
			FirstName: "FirstName_Test",
			LastName:  "LastName_Test",
			Email:     "test@test.org",
			Age:       23,
			Role:      "test-user",
			Sex:       "Female",
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		persons, err := psqlsqlx.ListPersons(db)
		if err != nil {
			b.Error(err)
		}
		_ = persons
	}

	b.StopTimer()
	psqlsqlx.DeletePersons(db)
	b.StartTimer()
}

func BenchmarkSQLXGet(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		psqlsqlx.InsertPerson(db, &entities.Person{
			FirstName: "FirstName_Test",
			LastName:  "LastName_Test",
			Email:     "test@test.org",
			Age:       23,
			Role:      "test-user",
			Sex:       "Female",
		})
	}

	persons, err := psqlsqlx.ListPersons(db)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		person, err := psqlsqlx.GetPerson(db, persons[rand.Intn(len(persons))].ID)
		if err != nil {
			b.Error(err)
		}
		_ = person
	}

	b.StopTimer()
	psqlsqlx.DeletePersons(db)
	b.StartTimer()
}
