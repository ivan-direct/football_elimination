package teams

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	team := New()
	xType := fmt.Sprintf("%T", team)
	if xType != "*teams.Team" {
		t.Fatalf("expecting *teams.Team, was%v\n", xType)
	}
}
