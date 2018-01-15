package main

import (
	"log"
	"net/http"

	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/simulation"
)

func init() {
	awsRegion := "us-west-2"
	dynamoEndpoint := "http://localhost:8000"

	// Send environment variables to Dynamo package which will be interacting with Dynamo
	// Dynamo package returns a channel signaling that it has been initialized
	dynamoInitChan := dynahelpers.SetEnvironmentVariables(awsRegion, dynamoEndpoint)

	// Send Dynamo initialization channel to all packages that will be interacting with Dynamo
	// Receive from these packages a channel that lets us know they have also initialized
	dbInitChan := boomyDB.InitializeDB(dynamoInitChan)
	convInitChan := chatbot.InitializeConv(dynamoInitChan)

	// Receive from all packages that let us know that they have been initialized
	<-dbInitChan
	log.Println("boomyDB package initialized")
	<-convInitChan
	log.Println("chatbot has been initialized")
}

func main() {

	go func() {
		// Router
		router := NewRouter()
		http.Handle("/", &MyServer{router})
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()

	simulation.StartTerminalConversation()

}
