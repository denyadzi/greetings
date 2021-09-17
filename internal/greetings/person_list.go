package greetings

import "github.com/denyadzi/greetings/pkg/person"

type PersonList person.PersonList

func (pl *PersonList) ShiftNextName() string {
	l := []person.Person(*pl)
	if len(l) == 0 {
		return ""
	}
	first_person := l[0]
	*pl = PersonList(l[1:])
	return first_person.Name
}
