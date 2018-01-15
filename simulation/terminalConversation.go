package simulation

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jcgarciaram/boomy/boomyDB"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/boomyAPI"
)

// StartTerminalConversation allows for testing in the terminal.
func StartTerminalConversation() {
	baseURL := "http://localhost:9999"

	body, err := post(baseURL+"/v1/api/boomy/resident/conversation", "", boomyDB.Resident{FirstName: "Pipo"})
	if err != nil {
		log.Fatal(err)
	}

	var m boomyAPI.Message

	// Unmarshal body into object pointer passed in request
	if err := json.Unmarshal(body, &m); err != nil {

		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error marshaling JSON to Message struct")

		log.Fatal(err)
	}

	token := strings.Trim(string(m.Token), "\n")
	token = strings.Trim(token, "\"")

	fmt.Println(m.Message)
	for _, qr := range m.QuickReplies {
		fmt.Printf("\t%s\n", qr)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		inM := boomyAPI.Message{
			Message: text[:len(text)-1],
		}

		body, err := post(baseURL+"/v1/api/boomy/resident/message", token, inM)
		if err != nil {
			log.Fatal(err)
		}

		var m boomyAPI.Message

		// Unmarshal body into object pointer passed in request
		if err := json.Unmarshal(body, &m); err != nil {

			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Warn("Error marshaling JSON to Message struct")

			log.Fatal(err)
		}

		fmt.Println(m.Message)
		for _, qr := range m.QuickReplies {
			fmt.Printf("\t%s\n", qr)
		}
	}
}
