package simulation

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/utils"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/boomyAPI"
)

type residentToken struct {
	ID    uint
	Token string
}

// StartTerminalConversation allows for testing in the terminal.
func StartTerminalConversation(conn utils.Conn) {
	baseURL := "http://localhost:9999"

	// Get all existing residents
	var rsdnts boomyDB.Residents
	if err := rsdnts.Find(conn); err != nil {
		log.Fatal(err)
	}

	// Read residents.json
	jsonFile, err := os.OpenFile("simulation/residents.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Error opening residents.json: ", err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading residents.json: ", err)
	}
	var rtSlice []residentToken
	if err := json.Unmarshal(byteValue, &rtSlice); err != nil {
		log.Fatal("Error unmarshaling residents.json to rtSlice: ", err)
	}

	fmt.Println("len(rtSlice)", len(rtSlice), "rtSlice:", rtSlice)

	var token string
	for {
		fmt.Println("Please select a resident or type NEW to create a new resident:")
		for _, r := range rsdnts {
			fmt.Printf("\t%d - %s\n", r.ID, r.FirstName+" "+r.LastName)
		}

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		choice := text[:len(text)-1]

		if len(rsdnts) == 0 || choice == "NEW" {
			body, err := post(baseURL+"/v1/api/boomy/resident/conversation", "", boomyDB.Resident{})
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

			token = strings.Trim(string(m.Token), "\n")
			token = strings.Trim(token, "\"")

			fmt.Println(m.Message)
			for _, qr := range m.QuickReplies {
				fmt.Printf("\t%s\n", qr)
			}

			rt := residentToken{
				ID:    m.ID,
				Token: m.Token,
			}

			rtSlice = append(rtSlice, rt)
			byteSlice, err := json.Marshal(rtSlice)
			if err != nil {
				log.Fatal("Error marshaling rtSlice: ", err)
			}
			if err := ioutil.WriteFile("simulation/residents.json", byteSlice, 0755); err != nil {
				log.Fatal("Error writing residents.json: ", err)
			}
			break

		} else {

			tokenFound := false
			for _, rt := range rtSlice {
				choiceInt, _ := strconv.Atoi(choice)
				if uint(choiceInt) == rt.ID {
					token = rt.Token
					tokenFound = true
					break
				}
			}

			body, err := get(baseURL+"/v1/api/boomy/conversation/resident", token)
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
			if tokenFound {
				break
			}
		}
		fmt.Printf("No valid resident selected\n\n")
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
