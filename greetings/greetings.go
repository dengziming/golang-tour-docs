package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

import "errors"

func Hello(name string) (message string, err error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Hellos(names[]string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	for _,name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v, Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
