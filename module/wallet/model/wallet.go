package model

import (
	"errors"
	"time"
)

var (
	ErrUserIDInvalid         = errors.New("UserID invalid")
	ErrBankIDInvalid         = errors.New("BankID invalid")
	ErrAccountIDInvalid      = errors.New("AccountID invalid")
	ErrItemNotFound          = errors.New("item not found")
	ErrCannotUpdateCreatedAt = errors.New("can not update CreatedAt")
)

/*
	`id` int not null AUTO_INCREMENT, primary key (`id`, `user_id`), # user_id để partition
    `user_id` int unsigned NOT NULL,
    `bank_id` int unsigned NOT NULL,
    `account_id` int unsigned NOT NULL,
    `amount` int NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
*/

type Wallet struct {
	Id        int        `json:"id" gorm:"column:id;"`
	UserID    int        `json:"user_id" gorm:"column:user_id;"`
	BankID    int        `json:"bank_id" gorm:"column:bank_id;"`
	AccountID int        `json:"account_id" gorm:"column:account_id;"`
	Amount    int        `json:"amount" gorm:"column:amount;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}

func (Wallet) TableName() string { return "wallets" }

func (item Wallet) Validate() error {
	switch {
	case item.UserID <= 0:
		return ErrUserIDInvalid
	case item.BankID <= 0:
		return ErrBankIDInvalid
	case item.BankID <= 0:
		return ErrBankIDInvalid
	case item.AccountID <= 0:
		return ErrAccountIDInvalid
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
