package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如，连接终止
			continue
		}

		go handleConn(conn)  // 一次处理一个链接
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1 * time.Second)
	}
	// 注意：忽略 input.Err() 中可能的错误

	c.Close()
}
