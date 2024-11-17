package dao

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"github/free-order-be/config"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
	"google.golang.org/api/option"
)

const (
	COUNTER_TABLE = "Counter"
	USER_TABLE    = "User"
	SHEET_TABLE   = "Sheet"
	ORDER_TABLE   = "Order"
	SHEET_SK      = "INFO#METADATA"
)

type DAO struct {
	client   *dynamo.DB
	UserDAO  IUserDAO
	SheetDAO ISheetDAO
	OrderDAO IOrderDAO
}

func NewDAO(ctx context.Context, db *dynamo.DB) (*DAO, error) {
	var (
		credential = []byte(config.Values.FirebaseCredential)
		err        error
	)
	if isEncodeBase64(config.Values.FirebaseCredential) {
		if credential, err = base64.StdEncoding.DecodeString(config.Values.FirebaseCredential); err != nil {
			return nil, err
		}
	}
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	firestore, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &DAO{
		UserDAO:  NewUserDAO(db),
		SheetDAO: NewSheetDAO(db),
		OrderDAO: NewOrderDAO(db, firestore),
		client:   db,
	}, nil
}

func NewDAORef(db *dynamo.DB) *DAO {
	return &DAO{
		client: db,
	}
}

func isEncodeBase64(str string) bool {
	_, err := base64.StdEncoding.DecodeString(str)
	return err == nil
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
