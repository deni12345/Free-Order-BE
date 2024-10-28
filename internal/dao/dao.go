package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
)

const (
	COUNTER_TABLE = "Counter"
	USER_TABLE    = "User"
)

type DAO struct {
	client  *dynamo.DB
	UserDAO IUserDAO
}

func NewDAO(db *dynamo.DB) *DAO {
	daoRef := &DAO{
		client: db,
	}

	return &DAO{
		UserDAO: &UserImpl{
			dao:   daoRef,
			table: db.Table(USER_TABLE),
		},
		client: db,
	}
}

func (dao *DAO) NextID(ctx context.Context, table string) (*uint, error) {
	output, err := dao.client.Client().UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(COUNTER_TABLE),
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

	return aws.Uint(uint(newID)), nil
}
