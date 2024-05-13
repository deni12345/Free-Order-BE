package domain

const (
	UserTable = "User"
)

type Users []*User
type User struct {
	ID       *uint  `gorm:"column:Id;"`
	UserName string `gorm:"column:UserName;"`
	Password string `gorm:"column:Password;"`
}

func (User) TableName() string {
	return UserTable
}

func (mu *User) CheckNil() *User {
	if mu.ID != nil {
		return mu
	}
	return nil
}
