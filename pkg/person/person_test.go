package person

import "testing"

func TestPersonType(t *testing.T) {
	p := Person{"Jojo", 33}

	if p.Name != "Jojo" || p.Age != 33 {
		t.Fatal("Failed person init")
	}
}
