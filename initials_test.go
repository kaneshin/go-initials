package initials

import (
	"reflect"
	"testing"
)

func TestInitial(t *testing.T) {
	expected1 := "SK"
	res1 := GetInitialsIfASCII("Shintaro", "Kaneko")
	if !reflect.DeepEqual(res1, expected1) {
		t.Fatalf("Expected %v, but %v:", expected1, res1)
	}

	expected2 := "SK"
	res2 := GetInitialsIfASCII("shintaro", "kaneko")
	if !reflect.DeepEqual(res2, expected2) {
		t.Fatalf("Expected %v, but %v:", expected2, res2)
	}

	expected3 := "SKFBBQ"
	res3 := GetInitialsIfASCII("shintaro", "kaneko", "foo", "bar", "baz", "qux")
	if !reflect.DeepEqual(res3, expected3) {
		t.Fatalf("Expected %v, but %v:", expected3, res3)
	}
}
