package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	Input  []string
	Output map[string][]string
}

var testSet = []TestData{
	{
		Input: []string{"пятак", "тяпка", "стол", "слот", "пятка", "гусь", "тпкяа", "усь", "тпка"},
		Output: map[string][]string{
			"пятак": {"тяпка", "пятка", "тпкяа"},
			"стол":  {"слот"},
		},
	},
	{
		Input: []string{"aaaa", "aaaa", "абра", "бара", "абраабра", "браа", "ауф", "фау", "уфа", "баар"},
		Output: map[string][]string{
			"абра": {"бара", "браа", "баар"},
			"ауф":  {"фау", "уфа"},
		},
	},
}

func TestNewDict(t *testing.T) {
	for _, data := range testSet {
		testOutput := NewDict(data.Input)
		if !reflect.DeepEqual(testOutput, data.Output) {
			t.Fatalf("error: test output inequal to prepared output: %v and %v", testOutput, data.Output)
		}
	}
}
