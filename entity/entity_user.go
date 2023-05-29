package entity

type Age struct {
	Date, Month, Year int
}

type User struct {
	Id                          int
	Username, Email, Membership string
	Born                        Age
}

type DatabaseUser struct {
	Data User
	Next *DatabaseUser
}

var UserMembership = []string{"Bronze", "Silver", "Gold", "Platinum", "Diamond"}
