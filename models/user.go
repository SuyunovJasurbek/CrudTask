package models

type User struct {
	ID int 
	FullName string
	NickName string
	Photo string
	Birthday string
	Location string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type UserList struct {
	ID int 
	FullName string
	NickName string
	Photo string
	Birthday string
}