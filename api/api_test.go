package api

import "testing"

func TestDoSomething(t *testing.T) {
	if err := DoSomething(); err != nil {
		t.Error(err)
	}
}
