package main

import (
	"log"
	"net/http"

	"github.com/jcgarciaram/demoPark/chatbot"
	"github.com/jcgarciaram/demoPark/demoParkDB"
	"github.com/jcgarciaram/demoPark/dynahelpers"
)

var (
	secret string
)

func init() {
	awsRegion := "us-west-2"
	dynamoEndpoint := "http://localhost:8000"
	secret = "script penates onion potence spinning exocrine"

	// Send environment variables to Dynamo package which will be interacting with Dynamo
	// Dynamo package returns a channel signaling that it has been initialized
	dynamoInitChan := dynahelpers.SetEnvironmentVariables(awsRegion, dynamoEndpoint)

	// Send Dynamo initialization channel to all packages that will be interacting with Dynamo
	// Receive from these packages a channel that lets us know they have also initialized
	dbInitChan := demoParkDB.InitializeDB(dynamoInitChan)
	convInitChan := chatbot.InitializeConv(dynamoInitChan)

	// Receive from all packages that let us know that they have been initialized
	<-dbInitChan
	log.Println("demoParkDB package initialized")
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

	// if err := demoParkConversation.BuildNewUserConversation(); err != nil {
	// 	log.Fatal(err)
	// }

	chatbot.StartTerminalConversation("user1")

}
