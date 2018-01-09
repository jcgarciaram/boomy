package dynahelpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/jcgarciaram/demoPark/utils"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DynamoGet queries Dynamo and retrieves the object being queried. object should be a pointer to a struct being populated
func DynamoGet(ID string, object interface{}) error {

	// Open new session for Dynamo
	db := getDynamoDB()

	tableName := utils.GetType(object)

	// Table name
	dynamoTable := db.Table(dbEC.AwsRegion + "_" + tableName)

	// Get object
	if err := dynamoTable.Get("ID", ID).One(object); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"ID":        ID,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
		}).Warn("Error getting from Dynamo")

		return err
	}

	return nil

}

// DynamoGetByField queries Dynamo and retrieves the object being queried. object should be a pointer to a struct being populated
func DynamoGetByField(fieldName string, value, objectType, returnSlice interface{}) error {

	// Open new session for Dynamo
	db := getDynamoDB()

	tableName := utils.GetType(objectType)

	// Table name
	dynamoTable := db.Table(dbEC.AwsRegion + "_" + tableName)

	// Get object
	if err := dynamoTable.Scan().Filter(fmt.Sprintf("'%s' = ?", fieldName), value).All(returnSlice); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"fieldName": fieldName,
			"value":     value,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
		}).Warn("Error getting from Dynamo")

		return err
	}

	return nil

}

// DynamoGetAll queries Dynamo and retrieves all objects for the table being queried. object should be a pointer to the individual object. sliceObject should be a pointer to a slice of the struct being populated
func DynamoGetAll(object interface{}, sliceObject interface{}) error {

	// Open new session for Dynamo
	db := getDynamoDB()

	tableName := utils.GetType(object)

	// Table name
	dynamoTable := db.Table(dbEC.AwsRegion + "_" + tableName)

	// Get object
	if err := dynamoTable.Scan().All(sliceObject); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
		}).Warn("Error getting from Dynamo")

		return err
	}

	return nil

}

// DynamoPut updates an object in Dynamo. object should be a pointer to an object
func DynamoPut(object interface{}) error {

	// Open new session for Dynamo
	db := getDynamoDB()

	tableName := utils.GetType(object)

	// Table name
	dynamoTable := db.Table(dbEC.AwsRegion + "_" + tableName)

	// Put object
	put := dynamoTable.Put(object)
	if err := put.Run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
			"object":    object,
		}).Warn("Error putting into Dynamo")

		return err
	}

	return nil
}

// CreateTable creates a new table in Dynamo if it does not exist
func CreateTable(object interface{}) error {

	// Open new session for Dynamo
	db := getDynamoDB()

	tableName := utils.GetType(object)

	// Table name
	dynamoTableName := dbEC.AwsRegion + "_" + tableName

	// Table name
	dynamoTable := db.Table(dynamoTableName)

	// Describe table. Done to see if table already exists
	describe := dynamoTable.Describe()
	if description, err := describe.Run(); err != nil && !strings.HasPrefix(err.Error(), "ResourceNotFoundException") {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
			"object":    object,
		}).Warn("Error describing table from Dynamo")

		return err
	} else if description.Active() {
		return nil
	}

	log.Printf("Creating %s table\n", dynamoTableName)

	// Create table
	create := db.CreateTable(dynamoTableName, object)
	if err := create.Run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"tableName": tableName,
			"awsRegion": dbEC.AwsRegion,
			"object":    object,
		}).Warn("Error creating table in Dynamo")

		return err
	}

	return nil
}

func getDynamoDB() *dynamo.DB {
	// Open new session for Dynamo
	return dynamo.New(
		session.New(),

		&aws.Config{
			Region:   aws.String(dbEC.AwsRegion),
			Endpoint: &dbEC.DynamoEndpoint,
		})

}
