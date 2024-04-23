package domain

const (
	UserTable = "User"
)

type User struct {
	ID       *uint  `gorm:"column:Id;"`
	UserName string `gorm:"column:UserName;"`
	Password string `gorm:"column:Password;"`
}

func (User) TableName() string {
	return UserTable
}
