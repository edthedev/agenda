package main

import (
	"reflect"
	"testing"
)

func TestGetAgenda(t *testing.T) {
	expected := []string{
		"- 7 am Breakfast",
		"- 9 am Second Breakfast",
		"- 11:00 am Elevensies",
		"- 12:15 pm Lunch",
		"- 2 pm Afternoon Tea",
	}

	result := getAgenda("agenda_test_file.md")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}
