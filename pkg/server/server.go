package server

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type InputArgs struct {

}

type Reply struct {

}

func (server *Server) PrintSomething(input *InputArgs, reply *Reply) error {
    fmt.Printf("recieved RPC call from another server")

    return nil
}

func CreateServer(serverPortNumber int, portNumberList []int, startupCompleteWG *sync.WaitGroup) {
    l, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPortNumber))
    if err != nil {
        fmt.Printf("Unable to listen to port number: %d", serverPortNumber)
    }
    fmt.Printf("Successfully began listening to port number %d\n", serverPortNumber)

    go http.Serve(l, nil)
    startupCompleteWG.Done()

    startupCompleteWG.Wait()

    for _, serverPort := range portNumberList {
        client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", serverPort))
        if err != nil {
            fmt.Printf("Failed to reach port number %d", serverPort)
        }

        input := &InputArgs{}
        reply := &Reply{}

        err = client.Call("Server.PrintSomething", input, reply)
    }
}
