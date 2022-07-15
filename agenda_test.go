package main

import (
	"reflect"
	"testing"
	"time"
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

func TestGetAgendaFile(t *testing.T) {

	testDate := time.Date(2022, 7, 15, 0, 0, 0, 0, time.UTC)
	expected := "/home/agendas/2022/07-15.md"
	format := "/home/agendas/{YYYY}/{MM}-{DD}.md"

	result := getAgendaFile(testDate, format)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}
