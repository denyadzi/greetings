package random_greet

import (
	"math/rand"
	"time"
)

type RandomGreetingSource []string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s RandomGreetingSource) GetGreetingFormat() string {
	formats := []string(s)
	return formats[rand.Intn(len(formats))]
}
