package main

import (
	"net/rpc"
	"net"
	"net/http"
	"gamePlat/server"
	"fmt"
)

func main(){
	arith := new(server.Arith)
	rpc.Register(arith)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "0.0.0.0:9696")
	if e != nil {
		fmt.Println("listen error:", e)

	}

	go http.Serve(l, nil)




}