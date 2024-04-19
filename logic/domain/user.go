package domain

const (
	UserTable = "User"
)

type User struct {
	id
}

func (User) TableName() string {
	return UserTable
}
