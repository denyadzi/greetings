package greetings

import (
	"regexp"
	"testing"
)

type NameList []string

type SingleGreeting string

func (nl *NameList) ShiftNextName() string {
	strs := []string(*nl)
	if len(strs) == 0 {
		return ""
	}

	name := strs[0]
	*nl = strs[1:]
	return name
}

func (sg SingleGreeting) GetGreetingFormat() string {
	return string(sg)
}

func TestHelloName(t *testing.T) {
	name := "Less"
	var gs SingleGreeting = "Hello"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name, gs)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q`, name, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	var gs SingleGreeting = "Hello"
	_, err := Hello("", gs)
	if err == nil {
		t.Fatalf("Must fail with empty name")
	}
}

func TestHellosNames(t *testing.T) {
	passes := make(map[string]bool)
	names := &NameList{"Jojo", "Mojo"}
	var greet_source SingleGreeting = "Let's do some testing, %s!"
	messages, err := Hellos(names, greet_source)
	enter := false
	for name, msg := range messages {
		enter = true
		want := regexp.MustCompile(`\b` + name + `\b`)
		passes[name] = want.MatchString(msg)
	}
	if !enter || err != nil {
		t.Fatalf("Has error")
	}
	for name, pass := range passes {
		if !pass {
			t.Fatalf(`Hasn't passed with name %v`, name)
		}
	}
}
