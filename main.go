package main

import (
	"context"
	"fmt"
	"net"

	"github.com/lisyaoran51/GoGameServerTest/connection"
	"github.com/lisyaoran51/GoGameServerTest/global"
)

func main() {

	ip := "127.0.0.1"
	port := "8001"

	ipaddr := ip + ":" + port
	listener, err := net.Listen("tcp", ipaddr)
	if err != nil {
		fmt.Printf("fail to listen tcp %v", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("fail to accept listener %v", err)
		return
	}

	gateConnection := connection.NewConnection(&conn)
	connection.GetConnectionManager().AddConnection(global.GATE, gateConnection)

	ctx, cancel := context.WithCancel(context.Background())
	go gateConnection.Start(ctx)

	cancelChan := make(chan int, 0)
	for {
		select {
		case <-cancelChan:
			cancel()
			return
		}
	}

}
