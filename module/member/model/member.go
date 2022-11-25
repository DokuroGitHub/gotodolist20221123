package model

import (
	"errors"
	"time"
)

var (
	ErrFirstNameIsBlank      = errors.New("FirstName can not be blank")
	ErrLastNameIsBlank       = errors.New("LastName can not be blank")
	ErrUserNameIsBlank       = errors.New("UserName can not be blank")
	ErrItemNotFound          = errors.New("item not found")
	ErrCannotUpdateCreatedAt = errors.New("can not update CreatedAt")
)

/*
   `firstname` varchar(25) NOT NULL,
   `lastname` varchar(25) NOT NULL,
   `username` varchar(16) NOT NULL,
   `email` varchar(35),
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` timestamp ON UPDATE CURRENT_TIMESTAMP,
   `deleted_at` timestamp NULL DEFAULT NULL
*/

type Member struct {
	FirstName string     `json:"firstname" gorm:"column:firstname;"`
	LastName  string     `json:"lastname" gorm:"column:lastname;"`
	UserName  string     `json:"username" gorm:"column:username;"`
	Email     string     `json:"email" gorm:"column:email;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}

func (Member) TableName() string { return "members" }

func (item Member) Validate() error {
	switch {
	case item.FirstName == "":
		return ErrFirstNameIsBlank
	case item.LastName == "":
		return ErrLastNameIsBlank
	case item.UserName == "":
		return ErrLastNameIsBlank
	default:
		return nil
	}
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
