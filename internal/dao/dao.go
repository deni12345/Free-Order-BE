package dao

import (
	"github.com/guregu/dynamo/v2"
)

type DAO struct {
	UserDAO IUserDAO
}

func NewDAO(db *dynamo.DB) *DAO {

	return &DAO{
		UserDAO: &UserImpl{
			client: db,
			table:  db.Table(UserTable),
		},
	}
}
