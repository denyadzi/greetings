package greetings

import (
	"errors"
	"fmt"
)

type Names interface {
	ShiftNextName() string
}

type GreetingSource interface {
	GetGreetingFormat() string
}

func Hello(name string, greet_source GreetingSource) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}
	return fmt.Sprintf(greet_source.GetGreetingFormat(), name), nil
}

func Hellos(names Names, greet_source GreetingSource) (map[string]string, error) {
	messages := make(map[string]string)
	for name := names.ShiftNextName(); name != ""; name = names.ShiftNextName() {
		msg := fmt.Sprintf(greet_source.GetGreetingFormat(), name)
		messages[name] = msg
	}
	return messages, nil
}
