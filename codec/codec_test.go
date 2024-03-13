package codec

import (
	"testing"
)

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

		exfilPacketDomains := EncodeDataToExfilPackets(tc.data, "example.com")

		fqdns := []string{}

		for _, ep := range exfilPacketDomains {
			fqdns = append(fqdns, ep.ToFqdn())
		}

		decodedData, err := DecodeDataFromFqdns(fqdns,  "example.com")

		if err != nil {
			t.Errorf("Failing feature: %s; Expected no error; got %s", tc.reason, err)
		}

		if decodedData != tc.data {
			t.Errorf("Failing feature: %s; Expected %s; got %s", tc.reason, tc.data, decodedData)
		}
	}
}
