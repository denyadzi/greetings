package random_greet

import "testing"

func TestRandomGreetingSourceType(t *testing.T) {
	_ = RandomGreetingSource{"One", "Two", "Three"}
}
