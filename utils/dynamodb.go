package utils

import (
	"context"
	"fmt"

	"LayeredStructured/domain/entities"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func InsertData(client *dynamodb.Client) {
	product := entities.Product{
		ProductUUID: "asdasd",
		Name:        "asdad",
		Description: "sadad",
		Brand:       "adsad",
		Price:       "asdsadas",
		Category:    "asdad",
		DateAdded:   "asdadada",
	}
	out, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("products"),
		Item: map[string]types.AttributeValue{
			"Id":          &types.AttributeValueMemberS{Value: product.ProductUUID},
			"Name":        &types.AttributeValueMemberS{Value: product.Name},
			"Description": &types.AttributeValueMemberS{Value: product.Description},
			"Brand":       &types.AttributeValueMemberS{Value: product.Brand},
			"Price":       &types.AttributeValueMemberS{Value: product.Price},
			"Category":    &types.AttributeValueMemberS{Value: product.Category},
			"DateAdded":   &types.AttributeValueMemberS{Value: product.Description},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

}
