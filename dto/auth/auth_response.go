package authdto

type AuthResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
}

type LoginResponse struct {
	// Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	// Email  string `gorm:"type: varchar(255)" json:"email"`
	// Token  string `gorm:"type: varchar(255)" json:"token"`
	// Status string `gorm:"type: varchar(255)" json:"status"`

	ID        int    `json:"id"`
	Fullname  string `gorm:"type: varchar(255)" json:"fullname"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Status    string `gorm:"type: varchar(255)" json:"status"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe string `json:"subscribe" gorm:"type: varchar(255)"`
}

type RegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	Fullname  string `gorm:"type: varchar(255)" json:"fullname"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Status    string `gorm:"type: varchar(255)" json:"status"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe string `json:"subscribe" gorm:"type: varchar(255)"`
}
