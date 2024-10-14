package main

import (
	"bufio"
	"os"
	//"flag"
	"fmt"
	"net"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	//msg := "Here is a message"

	//msgP := flag.String("msg", "Default message", "The message you want to send") // Pointer
	//flag.Parse()

	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		fmt.Printf("Enter text->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg) // de-reference the pointer by using *msgP
		read(conn)
	}
}
