package dynahelpers

import (
	"log"
	"time"
)

func init() {
	go func() {
		<-time.After(1 * time.Second)
		if dbEC.AwsRegion == "" {
			log.Fatal("AWS Region not set in dynahelpers")
		}

		if dbEC.DynamoEndpoint == "" {
			log.Fatal("Dynamo Endpoint not set in dynahelpers")
		}

		close(dynamoInitChan)
	}()

}
