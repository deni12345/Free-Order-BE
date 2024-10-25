package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/ptr"
	"github.com/guregu/dynamo/v2"
)

const (
	UserTable = "user"
)

type IUserDAO interface {
	Create(context.Context, *d.User) error
	FindByEmail(context.Context, string) (d.Users, error)
}

type UserImpl struct {
	client *dynamo.DB
	table  dynamo.Table
}

func (u *UserImpl) TableName() string {
	return u.table.Name()
}

func (u *UserImpl) Create(ctx context.Context, user *d.User) error {
	newID, err := u.nextID(ctx, UserTable)
	if err != nil {
		return err
	}
	user.ID = newID
	put := u.table.Put(user)
	return put.Run(ctx)
}

func (u *UserImpl) FindByEmail(ctx context.Context, userEmail string) (d.Users, error) {
	var result = d.Users{}
	err := u.table.Scan().Filter("$ = ?", "Email", userEmail).Limit(1).All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserImpl) nextID(ctx context.Context, table string) (*uint, error) {
	output, err := u.client.Client().UpdateItem(ctx, &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"CounterName": &types.AttributeValueMemberS{Value: table},
		},
		UpdateExpression: aws.String("SET CurrentValue = if_not_exists(CurrentValue, :start) + :increment"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":start":     &types.AttributeValueMemberN{Value: "0"},
			":increment": &types.AttributeValueMemberN{Value: "1"},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to increment counter: %w", err)
	}

	newID, err := strconv.ParseUint(output.Attributes["CurrentValue"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse counter value: %w", err)
	}

	return ptr.Uint(uint(newID)), nil
}
