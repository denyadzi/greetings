package person

import "testing"

func TestPersonType(t *testing.T) {
	p := Person{"Jojo", 33}

	if p.Name != "Jojo" || p.Age != 33 {
		t.Fatal("Failed person init")
	}
}

func TestPersonListType(t *testing.T) {
	p1 := Person{"Jojo", 33}
	p2 := Person{"Mojo", 30}
	pl := PersonList{p1, p2}

	name1 := pl.ShiftNextName()
	name2 := pl.ShiftNextName()
	name3 := pl.ShiftNextName()

	if name1 != "Jojo" {
		t.Fatalf("Person 1 name is %s", name1)
	}
	if name2 != "Mojo" {
		t.Fatalf("Person 2 name is %s", name2)
	}
	if name3 != "" {
		t.Fatalf("Empty string is %s", name3)
	}
}
