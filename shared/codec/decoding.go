package codec

import (
	"fmt"
	"strconv"
	"strings"
)

type decodablePacket struct {
	packetNo int
	payload  [3]string
}

func DecodeDataFromFqdns(fqdns []string, baseDomainName string) (string, error) {
	decodablePackets := []decodablePacket{}

	for _, fqdn := range fqdns {
		dp, err := convertFqdnToDecodablePacket(fqdn, baseDomainName)
		if err != nil {
			return "", err
		}
		decodablePackets = append(decodablePackets, dp)
	}

	return convertDecodablePacketsToFullPayload(decodablePackets)
}

func convertFqdnToDecodablePacket(fqdn string, mainDomain string) (decodablePacket, error) {
	if !strings.HasSuffix(fqdn, mainDomain) {
		return decodablePacket{}, fmt.Errorf("FQDN does not end with main domain")
	}

	subdomains := strings.TrimSuffix(fqdn, "."+mainDomain)

	dp := decodablePacket{}

	parts := strings.Split(subdomains, ".")

	copy(dp.payload[:], parts[1:])

	metadata := strings.Split(parts[0], "-")
	packetNo, err := strconv.Atoi(metadata[1])

	if err != nil {
		return decodablePacket{}, fmt.Errorf("packetNo is not a number")
	}

	dp.packetNo = packetNo

	return dp, nil
}

func convertDecodablePacketsToFullPayload(packets []decodablePacket) (string, error) {
	result := ""
	for _, packet := range packets {
		for _, part := range packet.payload {
			decoded, err := base36ToString(part)
			if err != nil {
				return "", err
			}
			result += decoded
		}
	}
	return result, nil
}
