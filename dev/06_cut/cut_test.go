package main

import (
	"log"
	"testing"
)

func TestParseColumns(t *testing.T) {
	var validTestData = []string{
		"1-8",
		"1,9",
		"2",
	}

	var invalidTestData = []string{
		"afafaffaafaf",
		"-2",
		"",
	}

	for _, datum := range validTestData {
		_, err := ParseColumns(datum)
		if err != nil {
			t.Error(err)
		}
	}

	for _, datum := range invalidTestData {
		cols, err := ParseColumns(datum)
		if err == nil {
			t.Errorf("something goes wrong: %v", cols)
		}
	}
}

func TestFieldCut(t *testing.T) {
	var TabTestData = []string{
		"hello\tholla\thallo",
		"foobar\tfoooo\tfeeee",
	}

	var CommaTestData = []string{
		"hello,holla,hallo",
		"foobar,foooo,feeee",
	}

	cut, err := FieldCut(TabTestData, "\t", false, 0, 2)
	if err != nil {
		t.Error(err)
	}

	out := "hello hallo \nfoobar feeee\n"

	if len(out) != len(cut) {
		t.Errorf("invalid return %d and %d\n", len(out), len(cut))
		log.Print(cut)
		log.Print(out)
	}

	cut, err = FieldCut(TabTestData, "\t", false, 4)

	if err == nil {
		t.Error(err, cut)
	}

	cut, err = FieldCut(CommaTestData, ",", false, 0, 2)
	if err != nil {
		t.Error(err)
	}

	if len(out) != len(cut) {
		t.Errorf("invalid return %d and %d\n", len(out), len(out))
		log.Print(cut)
		log.Print(out)
	}

}
