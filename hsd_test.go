package hsd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestStringDistanceTable(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{name: "lhs is longer than rhs", lhs: "foo", rhs: "fo", want: -1},
		{name: "lhs is shorter than rhs", lhs: "fo", rhs: "foo", want: -1},
		{name: "No diff", lhs: "foo", rhs: "foo", want: 0},
		{name: "1 diff", lhs: "foo", rhs: "foh", want: 1},
		{name: "2 diff", lhs: "foo", rhs: "fhh", want: 2},
		{name: "3 diff", lhs: "foo", rhs: "hhh", want: 3},
		{name: "multibyte", lhs: "あいう", rhs: "あいえ", want: 1},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v \n", tc.name, tc.want, got)
		} else {
			fmt.Printf("%s: expected: %v, got: %v \n", tc.name, tc.want, got)
		}
	}
}
