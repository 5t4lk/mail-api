package user

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Output(message, title, to interface{}) (string, error) {
	currentTime := Nowadays()
	_, err := fmt.Printf("Message: [%s]\nTitle: [%s]\nSent to: [%s]\n", message, title, to)
	if err != nil {
		return "", err
	}

	return currentTime, nil
}

func AskUser() (interface{}, interface{}, interface{}, error) {
	var to, title, message interface{}
	to, err := input("Please enter email to delivery: ", to)
	if err != nil {
		return nil, nil, nil, err
	}
	title, err = input("Please enter the title of message: ", title)
	if err != nil {
		return nil, nil, nil, err

	}
	message, err = input("Please enter the message: ", message)
	if err != nil {
		return nil, nil, nil, err
	}

	return to, title, message, nil
}

func input(text, value interface{}) (interface{}, error) {
	fmt.Println(text)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value = scanner.Text()

	return value, nil
}

func Nowadays() string {
	now := time.Now()
	tTime := fmt.Sprintf("%s", now.Local().Format(time.RFC1123))

	return tTime
}
