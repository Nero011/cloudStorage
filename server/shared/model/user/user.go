package user

type User struct {
	Name     string `json:"name", gorm:"name"`
	Password string `json:"password", gorm:"password"`
}
