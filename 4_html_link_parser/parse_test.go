package link

import (
	"fmt"
	"os"
	"testing"
)

func TestCode(t *testing.T) {
	s, err := os.Open("ex2.html")
	if err != nil {
		panic(err)
	}

	links, err := Parse(s)
	if err != nil {
		panic(err)
	}

	exp := "[{Href:https://www.twitter.com/joncalhoun Text:Check me out on twitter} {Href:https://github.com/gophercises Text:Gophercises is on Github !}]\n"
	actual := fmt.Sprintf("%+v\n", links)

	if actual != exp {
		t.Fatalf("expected: %+v\n actual: %+v", exp, actual)
	}
}
