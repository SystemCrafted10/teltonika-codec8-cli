package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"chrisfarrugia.dev/teltonika-codec8-cli/parser"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: teltonika-codec8-cli <hexstring>")
	}

	hexStr := os.Args[1]

	// Decode hex string to []byte
	data, err := hex.DecodeString(hexStr)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid hex string:", err)
		os.Exit(2)
	}

	// Codec ID is at position 8 (1 byte)
	codecID := int(data[8])

	var packet *parser.TeltonikaAvlPacket

	switch codecID {
	case 8:
		packet, err = parser.ParseCodec8(data)
	case 142:
		packet, err = parser.ParseCodec8Extended(data)
	default:
		packet, err = nil, fmt.Errorf("unsupported codec - %d", codecID)
	}

	// Parse packet

	if err != nil {
		fmt.Fprintln(os.Stderr, "Parse error", err)
		os.Exit(3)
	}

	// Output as JSON
	json.NewEncoder(os.Stdout).Encode(packet)

}

// go build -o teltonika-codec8-cli ./cmd/teltonika-codec8/main.go
// scp -i ~/.ssh_iot/id_ecdsa teltonika-codec8-cli root@51.89.4.109:/usr/bin/
// scp -i ~/.ssh_iot/id_ecdsa teltonika-codec8-cli ubuntu@57.129.34.12:/home/ubuntu
