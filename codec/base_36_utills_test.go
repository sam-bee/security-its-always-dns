package codec

import (
	"testing"
)

func TestBase36Encoding(t *testing.T) {

	input := "A"
	expected := "1t"
	result := stringToBase36([]byte(input))

	if result != expected {
		t.Errorf("Expected %s; got %s", expected, result)
	}
}

func TestDecodingFromBase36(t *testing.T) {

	input := "1t"
	expected := "A"
	result, _ := base36ToString(input)

	if result != expected {
		t.Errorf("Expected %s; got %s", expected, result)
	}
}
