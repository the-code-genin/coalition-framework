package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/the-code-genin/coalition-p2p"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Remote node address must be specified as the argument")
	}

	host, err := coalition.NewHost(3001)
	if err != nil {
		panic(err)
	}
	defer host.Close()

	peerKey := host.PeerKey()
	fmt.Printf("Sending ping from %s\n", hex.EncodeToString(peerKey[:]))
	response, err := host.SendMessage(os.Args[1], 1, coalition.PingMethod, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.(string))
	fmt.Printf("Peers: %d\n", len(host.Peers()))
}
