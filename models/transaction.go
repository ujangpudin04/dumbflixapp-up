package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"dueDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attache   string       `json:"attache"`
	Price     int          `json:"price"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID        int          `json:"id"`
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"dueDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attache   string       `json:"attache"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
