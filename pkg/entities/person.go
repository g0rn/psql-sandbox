package entities

import "fmt"

type Person struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Age       int16  `db:"age"`
	Role      string `db:"role"`
	Sex       string `db:"sex"`
}

func (p Person) String() string {
	return fmt.Sprintf(
		"ID:\t%v\nName:\t%v %v\nEmail:\t%v\nAge:\t%v\nRole:\t%v\nSex:\t%v\n",
		p.ID,
		p.FirstName,
		p.LastName,
		p.Email,
		p.Age,
		p.Role,
		p.Sex,
	)
}
