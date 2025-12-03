package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

const port = ":42069"

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to listen on %v: %v\n", port, err)
		return
	}
	defer lis.Close()

	fmt.Printf("Listening for TCP traffic on %v\n", port)

	for {
		conec, err := lis.Accept()
		if err != nil {
			fmt.Printf("Failde to accept contion: %v \n", err)
			return
		}
		fmt.Printf("Accepted connection from 127.0.0.1:%v \n", port)

		linesChan := getLinesChannel(conec)

		for line := range linesChan {
			fmt.Printf("%s", line)
		}
	}

}

func getLinesChannel(c net.Conn) <-chan string {

	linesChan := make(chan string)

	go func() {
		currentLineContents := ""
		defer c.Close()
		defer close(linesChan)

		for {

			b := make([]byte, 8, 8)
			n, err := c.Read(b)
			if err != nil {
				if currentLineContents != "" {
					linesChan <- fmt.Sprintf("read: %s\n", currentLineContents)
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			str := string(b[:n])
			part := strings.Split(str, "\n")

			for i := 0; i < len(part)-1; i++ {
				linesChan <- fmt.Sprintf("%s%s", currentLineContents, part[i])
				currentLineContents = ""
			}
			currentLineContents += part[len(part)-1]
		}
	}()

	return linesChan

}
