package main

import (
	"log"
	"reflect"
	"sort"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	testVal := []string{
		"fooo",
		"bar",
		"fooo",
		"fefefe",
	}

	out := UniqueStrings(testVal)

	if len(out) != 3 {
		t.Errorf("not uniqe %v\n", out)
	}

	out = UniqueStrings([]string{})

	if len(out) != 0 {
		t.Errorf("not empty %v\n", out)
	}
}

func TestUnpackStrings(t *testing.T) {
	testVal := []string{
		"fooo foa fob\n",
		"bar ber bir\n",
		"fooo feee feea\n",
	}

	out := UnpackStrings(testVal)
	validateVal := [][]string{
		{"fooo", "foa", "fob"},
		{"bar", "ber", "bir"},
		{"fooo", "feee", "feea"},
	}
	if !reflect.DeepEqual(out, validateVal) {
		t.Error("wrong!!")
		log.Print(out)
	}

}

func TestSortingAction(t *testing.T) {
	testVal := [][]string{
		{"1"},
		{"4"},
		{"5"},
		{"67"},
		{"-1"},
		{"7"},
	}

	validateTestVal := [][]string{
		{"-1"},
		{"1"},
		{"4"},
		{"5"},
		{"7"},
		{"67"},
	}

	_, err := SortingAction(testVal, 0, true, false, func(x interface{}, less func(i int, j int) bool) bool {
		sort.Slice(x, less)
		return true
	})

	if !reflect.DeepEqual(testVal, validateTestVal) {
		t.Error("wrong!!")
		log.Print(testVal)
	}
	if err != nil {
		t.Error(err)
	}
}
