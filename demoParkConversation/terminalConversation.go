package demoParkConversation

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func StartTerminalConversation(senderId string) chan string {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		response, qrs, err := getResponse(senderId, text[:len(text)-1])
		if err != nil {
			log.Println("Error from getResponse")
		}
		fmt.Println(response)
		for _, qr := range qrs {
			fmt.Printf("\t%s\n", qr.ReplyText)
		}
	}
}
