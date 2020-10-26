package mdb

type FriendsShip struct {
	ID           int    `db:"id"`
	UserOne      string `db:"user_one_email"`
	UserTwo      string `db:"user_two_email"`
	UpdateStatus bool   `db:"update_status"`
}
