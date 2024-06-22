package udpclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"pong/src/pong"
)

type Message struct {
	Key  string
	User int
}

func Get() pong.Game {
	udpAddr, err := net.ResolveUDPAddr("udp", "8000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send a message to the server
	buf := make([]byte, 0)

	// Read from the connection untill a new line is send
	bufio.NewReader(conn).Read(buf)
	g := pong.NewGame()
	json.Unmarshal(buf, g)
	if err != nil {
		panic(err)
	}
	return *g
}

func Request(msg Message) pong.Game {
	udpAddr, err := net.ResolveUDPAddr("udp", "8000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send a message to the server
	buf, _ := json.Marshal(msg)
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}

	// Read from the connection untill a new line is send
	bufio.NewReader(conn).Read(buf)
	g := pong.NewGame()
	json.Unmarshal(buf, g)
	if err != nil {
		panic(err)
	}
	return *g
}