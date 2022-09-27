package transactiondto

import (
	"backend_project/models"
	"time"
)

type TransactionResponse struct {
	ID       int                 `json:"id"`
	Stardate time.Time           `json:"startdate" form:"startdate"`
	Duedate  time.Time           `json:"duedate" form:"duedate"`
	UserID   int                 `json:"user_id" form:"user_id"`
	User     models.UserResponse `json:"user"`
	Attache  string              `json:"attache" form:"attache"`
	Status   string              `json:"status" form:"status"`
	Email    string              `json:"email" form:"email"`
}
