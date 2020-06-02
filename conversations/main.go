package main

/*
import (
	"log"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/boomyAPI"
	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbServer = os.Getenv("MYSQL_SERVER")
	dbPort   = os.Getenv("MYSQL_PORT")
	dbUser   = os.Getenv("MYSQL_USER")
	dbPass   = os.Getenv("MYSQL_PASS")

	db *gorm.DB
)

func init() {

	// open mysql connection
	var err error
	db, err = gorm.Open("mysql", "%s:%s@%s:%s/?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbServer, dbPort)
	if err != nil {
		logrus.Fatal("error opening connection to DB")
	}

	// Share db variable with all packages that will be interacting with MySQL
	// Receive from these packages a channel that lets us know they have initialized
	convInitChan := chatbot.InitializeConv(db)
	boomyAPI.InitializeDB(db)

	// Receive from all packages that let us know that they have been initialized
	<-convInitChan
	log.Println("chatbot has been initialized")
}

func main() {
	defer db.Close()

	newResidentConvTree := buildNewResidentConversation()
	newResidentConvTree.SaveEntireTree()

}
*/
