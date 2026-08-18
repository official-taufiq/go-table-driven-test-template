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
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, v := range tests {
		got := Split(v.input, v.sep)
		if !reflect.DeepEqual(v.want, got) {
			t.Fatalf("Expected: %v, Got: %v", v.want, got)
		}
	}
}
