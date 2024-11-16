package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/guregu/dynamo/v2"
)

type IUserDAO interface {
	Create(context.Context, *d.User) error
	FindByID(context.Context, uint) (*d.User, error)
	FindByEmail(context.Context, string) (*d.User, error)
}

func NewUserDAO(client *dynamo.DB) IUserDAO {
	return &UserImpl{
		dao:   NewDAORef(client),
		table: client.Table(USER_TABLE),
	}
}

type UserImpl struct {
	dao   *DAO
	table dynamo.Table
}

func (u *UserImpl) TableName() string {
	return u.table.Name()
}

func (u *UserImpl) Create(ctx context.Context, user *d.User) error {
	newID, err := u.dao.NextID(ctx, USER_TABLE)
	if err != nil {
		return err
	}

	user.ID = createUserPK(newID)
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

func (u *UserImpl) FindByEmail(ctx context.Context, email string) (*d.User, error) {
	var results = d.Users{}
	err := u.table.Scan().Filter("Email=?", email).Limit(1).All(ctx, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return &d.User{}, nil
	}
	return results[0], nil
}

func createUserPK(id *uint) *string {
	return aws.String(fmt.Sprintf("USER#%v", *id))
}
