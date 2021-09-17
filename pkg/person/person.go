package person

type Person struct {
	Name string
	Age  int
}

type PersonList []Person

func (pl *PersonList) ShiftNextName() string {
	l := []Person(*pl)
	if len(l) == 0 {
		return ""
	}
	first_person := l[0]
	*pl = PersonList(l[1:])
	return first_person.Name
}
