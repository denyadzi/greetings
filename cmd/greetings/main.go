package main

import (
	"fmt"
	"log"

	"github.com/denyadzi/greetings/internal/greetings"
	"github.com/denyadzi/greetings/pkg/person"
	"github.com/denyadzi/greetings/pkg/random_greet"
)

func main() {
	p1 := person.Person{"Phillip", 33}
	p2 := person.Person{"Angie", 31}
	p3 := person.Person{"Julia", 29}
	p4 := person.Person{"John", 34}
	pl := &person.PersonList{p1, p2, p3, p4}
	source := random_greet.RandomGreetingSource{
		"Hi, nice to meet you, %s",
		"Hello! You're welcome, lovely %s!",
		"Good day, %s!",
	}
	messages, err := greetings.Hellos(pl, source)
	if err != nil {
		log.Fatal("Error occured")
	}
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
