package models

type User struct {
	ID        int    `json:"id" db:"id" binding:"required"`
	FullName  string `json:"full_name" db:"full_name" binding:"required"`
	NickName  string `json:"nick_name" db:"nick_name" binding:"required"`
	Photo     string `json:"photo" db:"photo" binding:"required"`
	Birthday  string `json:"birth_day" db:"birth_day" binding:"required"`
	Location  string `json:"location" db:"location" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at" `
	UpdatedAt string `json:"updated_at" db:"updated_at" `
	DeletedAt string `json:"deleted_at" db:"deleted_at" `
}

type UserHandler struct {
	FullName  string `json:"full_name" db:"full_name" binding:"required"`
	NickName  string `json:"nick_name" db:"nick_name" binding:"required"`
	Photo     string `json:"photo" db:"photo" binding:"required"`
	Birthday  string `json:"birth_day" db:"birth_day" binding:"required"`
	Location  string `json:"location" db:"location" binding:"required"`	
}

type UserService struct {
	FullName  string `json:"full_name" db:"full_name" binding:"required"`
	NickName  string `json:"nick_name" db:"nick_name" binding:"required"`
	Photo     string `json:"photo" db:"photo" binding:"required"`
	Birthday  string `json:"birth_day" db:"birth_day" binding:"required"`
	Location  string `json:"location" db:"location" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at" `
	UpdatedAt string `json:"updated_at" db:"updated_at" `
	DeletedAt string `json:"deleted_at" db:"deleted_at" `
}

type UserUpdate struct {
	FullName  string `json:"full_name" db:"full_name" binding:"required"`
	NickName  string `json:"nick_name" db:"nick_name" binding:"required"`
	Photo     string `json:"photo" db:"photo" binding:"required"`
	Birthday  string `json:"birth_day" db:"birth_day" binding:"required"`
	Location  string `json:"location" db:"location" binding:"required"`
	UpdatedAt string `json:"updated_at" db:"updated_at" `
}
type Users struct {
	ID       int    `json:"id" db:"id" binding:"required"`
	FullName string `json:"full_name" db:"full_name" binding:"required"`
	NickName string `json:"nick_name" db:"nick_name" binding:"required"`
	Photo    string `json:"photo" db:"photo" binding:"required"`
	Birthday string `json:"birth_day" db:"birth_day" binding:"required"`
}

type Error struct {
	Message string
}

type Response struct {
	Message string
}