package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected: %v, got:%v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	type test map[string]struct {
		input string
		sep   string
		want  []string
	}

	tests := test{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, v := range tests {

		t.Run(name, func(t *testing.T) {
			got := Split(v.input, v.sep)
			if !reflect.DeepEqual(v.want, got) {
				t.Fatalf("Expected: %v, Got: %v", v.want, got)
			}
		})
	}
}
