package user

import "time"


type Users struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	Status		bool
}

func (this *Users) HigUser(id int, firstname string, lastname string, email string, createdat time.Time, status bool) {
	this.ID = id
	this.FirstName = firstname
	this.LastName = lastname
	this.Email = email
	this.CreatedAt = createdat
	this.Status = status
}
