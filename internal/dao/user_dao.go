package dao

import (
	"context"
	d "github/free-order-be/internal/domain"

	"github.com/guregu/dynamo/v2"
)

type IUserDAO interface {
	Create(context.Context, *d.User) error
	FindByID(context.Context, uint) (*d.User, error)
	FindByEmail(context.Context, string) (d.Users, error)
}

func NewUserDAO(client *dynamo.DB) IUserDAO {
	return &UserImpl{
		client: client,
		table:  client.Table(USER_TABLE),
	}
}

type UserImpl struct {
	dao    *DAO
	client *dynamo.DB
	table  dynamo.Table
}

func (u *UserImpl) TableName() string {
	return u.table.Name()
}

func (u *UserImpl) Create(ctx context.Context, user *d.User) error {
	newID, err := u.dao.NextID(ctx, USER_TABLE)
	if err != nil {
		return err
	}

	user.ID = newID
	return u.table.Put(user).Run(ctx)
}

func (u *UserImpl) FindByID(ctx context.Context, ID uint) (*d.User, error) {
	var result = &d.User{}
	err := u.table.Get("Id", ID).One(ctx, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserImpl) FindByEmail(ctx context.Context, email string) (d.Users, error) {
	var results = d.Users{}
	err := u.table.Scan().Filter("$=?", "Email", email).All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
