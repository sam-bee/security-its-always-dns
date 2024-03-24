package codec

import (
	"regexp"
	"strings"
	"testing"
)

func TestEncoding(t *testing.T) {

	input := "A"
	expectedEndOfDomain := "-0.1t.example.com"
	expectedPatternForStartOfDomain := regexp.MustCompile(`^[0-9a-f]{1,6}`)

	result := GetDomainsToLookUp(input, "example.com")

	if len(result) != 1 {
		t.Errorf("Expected 1; got %d", len(result))
	}

	if !strings.HasSuffix(result[0], expectedEndOfDomain) {
		t.Errorf("Expected string ending in %s; got %s", expectedEndOfDomain, result[0])
	}

	if !expectedPatternForStartOfDomain.MatchString(result[0]) {
		t.Errorf("Expected string starting with up to 6 hex characters; got %s", result[0])
	}
}

func TestDecoding(t *testing.T) {

	input := []string{"abc123-0.1t.example.com"}
	expected := "A"

	result, err := DecodeDataFromFqdns(input, "example.com")

	if err != nil {
		t.Errorf("Expected no error; got %s", err)
	}

	if result != expected {
		t.Errorf("Expected %s; got %s", expected, result)
	}
}

func TestEncodingAndDecoding(t *testing.T) {

	type test struct {
		data   string
		reason string
	}

	tests := []test{
		{
			data:   "A",
			reason: "String is shorter than a single 63 byte subdomain level",
		},
		{
			data:   "Emoji 😊",
			reason: "Test a string that has weird characters",
		},
		{
			data:   "0123456789012345678901234567890123456789",
			reason: "String just fits in one 63 byte subdomain level",
		},
		{
			data:   "01234567890123456789012345678901234567890123456789012345678901234567890123456789",
			reason: "String requires two subdomain levels",
		},
		{
			data:   "01234567890123456789012345678901234567890",
			reason: "String requires one full and one partial subdomain level",
		},
		{
			data:   "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
			reason: "String requires a full packet of three subdomain levels",
		},
		{
			data:   "0123456789012345678901234567890123456789012345678901234567891234012345678901234567890123456789012345678901234567890123456789123401234567890123456789012345678901234567890123456789012345678912345",
			reason: "String is big and requires more than one exfil packet",
		},
	}

	for _, tc := range tests {

		exfilPacketDomains := GetDomainsToLookUp(tc.data, "example.com")

		decodedData, err := DecodeDataFromFqdns(exfilPacketDomains, "example.com")

		if err != nil {
			t.Errorf("Failing feature: %s; Expected no error; got %s", tc.reason, err)
		}

		if decodedData != tc.data {
			t.Errorf("Failing feature: %s; Expected %s; got %s", tc.reason, tc.data, decodedData)
		}
	}
}
