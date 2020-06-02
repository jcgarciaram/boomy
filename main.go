package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/boomyAPI"
	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jcgarciaram/boomy/simulation"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbServer = os.Getenv("MYSQL_SERVER")
	dbPort   = os.Getenv("MYSQL_PORT")
	dbUser   = os.Getenv("MYSQL_USER")
	dbPass   = os.Getenv("MYSQL_PASS")
	dbName   = os.Getenv("MYSQL_DB")

	db *gorm.DB
)

func init() {

	// open mysql connection
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&clientFoundRows=true", dbUser, dbPass, dbServer, dbPort, dbName))
	if err != nil {
		logrus.Fatalf("error opening connection to DB: %s", err)
	}
	db.SingularTable(true)

	// Share db variable with all packages that will be interacting with MySQL
	// Receive from these packages a channel that lets us know they have initialized
	dbInitChan := boomyDB.InitializeDB(db)
	convInitChan := chatbot.InitializeConv(db)

	// Receive from all packages that let us know that they have been initialized
	<-dbInitChan
	log.Println("boomyDB package initialized")
	<-convInitChan
	log.Println("chatbot has been initialized")

	boomyAPI.Initialize(db)
}

func main() {
	defer db.Close()

	go func() {
		// Router
		router := NewRouter()
		http.Handle("/", &MyServer{router})
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()

	simulation.StartTerminalConversation(db)

}
