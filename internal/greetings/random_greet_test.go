package greetings

import "testing"

func TestRandomGreetingSourceType(t *testing.T) {
	_ = RandomGreetingSource{"One", "Two", "Three"}
}
