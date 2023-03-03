package main

import (
	"errors"
	"fmt"
	"net/rpc"
	"raft/pkg/server"
	"sync"
)

func main() {
    startServers(5)
}

func startServers(numberOfServers int) error {
    if numberOfServers > 10 {
        return errors.New("Can only create up to 10 servers")
    }

    serverObject := new(server.Server)
    rpc.Register(serverObject)
    rpc.HandleHTTP()

    fmt.Println("Finished registering server object")
    
    basePort := 8080
    var finalShutDownWG, startupCompleteWG sync.WaitGroup
    portNumberList := make([]int, numberOfServers)

    for i:=0 ; i<numberOfServers ; i++ {
        portNumberList[i] = basePort + i
    }
    startupCompleteWG.Add(numberOfServers)

    for i:=0 ; i<numberOfServers ; i++ {
        finalShutDownWG.Add(1)
        go func (i int) {
            defer finalShutDownWG.Done()
            server.CreateServer(basePort + i, portNumberList, &startupCompleteWG)
        }(i)
    }

    finalShutDownWG.Wait()

    fmt.Println("reached here")
    return nil
}

