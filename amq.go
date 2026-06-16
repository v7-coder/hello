package hello

import "fmt"

func SendMessage(message string) {
	fmt.Println(fmt.Sprintf("Hello %s", message))
}
