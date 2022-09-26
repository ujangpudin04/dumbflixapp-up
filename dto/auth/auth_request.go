package authdto

type RegisterRequest struct {
	Fullname  string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email     string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password  string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender    string `json:"gender"  form:"gender" validate:"required"`
	Phone     string `json:"phone"  form:"phone" validate:"required"`
	Address   string `json:"address"  form:"address" validate:"required"`
	Subscribe string `json:"subscribe"  form:"form"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
