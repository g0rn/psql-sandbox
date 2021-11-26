package main

import (
	"log"

	"github.com/g0rn/psql-sandbox/pkg/entities"
	psqlsqlx "github.com/g0rn/psql-sandbox/pkg/psql-sqlx"
)

func main() {
	db, err := psqlsqlx.NewClient()
	if err != nil {
		log.Fatalf("Can't connect to PSQL, err: %v\n", err)
	}
	defer db.Close()

	err = psqlsqlx.InsertPerson(db, &entities.Person{
		FirstName: "Alexei",
		LastName:  "Bychkovsky",
		Email:     "Alexei.Bychkovsky@gmail.com",
		Age:       31,
		Role:      "Owner",
		Sex:       "Male",
	})
	if err != nil {
		log.Fatalf("Can't insert person, err: %v\n", err)
	}

	person, err := psqlsqlx.GetPerson(db, "73d894ed-87b1-4a45-b92a-47c861fa01ee")
	if err != nil {
		log.Fatalf("Can't get person, err: %v\n", err)
	}
	log.Println(person)

	persons, err := psqlsqlx.ListPersons(db)
	if err != nil {
		log.Fatalf("Can't list persons, err: %v\n", err)
	}
	for i := range persons {
		log.Println(persons[i])
	}

	psqlsqlx.DeletePerson(db, "73d894ed-87b1-4a45-b92a-47c861fa01ee")
}
