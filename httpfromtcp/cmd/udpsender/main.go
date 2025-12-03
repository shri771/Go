package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "127.0.0.1:42069"

func main() {

	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Printf("Could not resolve udp addr on port %v: %v \n", port, err)
		return
	}
	fmt.Printf("Staring UDP on prot %v.........\n", port)

	con, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Could not setup connection: %v \n", err)
		return
	}
	defer con.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Could not read from read.io: %v \n", err)
		}
		_, err = con.Write([]byte(line))
		if err != nil {
			log.Printf("Could not write: %v \n", err)
		}

	}

}
