package usersdto

type UserResponse struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender" `
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Subscribe string `json:"subscribe"`
	Image     string `json:"image"`
}

type UserFotoProfilResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Image string `json:"image"`
}
