package model

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

func (u *User) TableName() string {
	return "user"
}
