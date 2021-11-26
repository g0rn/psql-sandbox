package psqlpq_test

import (
	"database/sql"
	"log"
	"math/rand"
	"testing"

	"github.com/g0rn/psql-sandbox/pkg/entities"
	psqlpq "github.com/g0rn/psql-sandbox/pkg/psql-pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = psqlpq.NewClient()
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkNativeInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		psqlpq.InsertPerson(db, &entities.Person{
			FirstName: "FirstName_Test",
			LastName:  "LastName_Test",
			Email:     "test@test.org",
			Age:       23,
			Role:      "test-user",
			Sex:       "Female",
		})
	}

	psqlpq.DeletePersons(db)
}

func BenchmarkNativeList(b *testing.B) {
	for i := 0; i < 1000; i++ {
		psqlpq.InsertPerson(db, &entities.Person{
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
		persons, err := psqlpq.ListPersons(db)
		if err != nil {
			b.Error(err)
		}
		_ = persons
	}

	b.StopTimer()
	psqlpq.DeletePersons(db)
	b.StartTimer()
}

func BenchmarkNativeGet(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		psqlpq.InsertPerson(db, &entities.Person{
			FirstName: "FirstName_Test",
			LastName:  "LastName_Test",
			Email:     "test@test.org",
			Age:       23,
			Role:      "test-user",
			Sex:       "Female",
		})
	}

	persons, err := psqlpq.ListPersons(db)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		person, err := psqlpq.GetPerson(db, persons[rand.Intn(len(persons))].ID)
		if err != nil {
			b.Error(err)
		}
		_ = person
	}

	b.StopTimer()
	psqlpq.DeletePersons(db)
	b.StartTimer()
}
