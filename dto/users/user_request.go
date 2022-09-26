package usersdto

type CreateUserRequest struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe string `json:"subscribe" gorm:"type: varchar(255)"`
	Image     string `json:"image"`
}

type UpdateUserRequest struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Subscribe string `json:"subscribe"`
	Image     string `json:"image"`
}

type CreateUserFotoProfilRequest struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Image string `json:"image"`
}
