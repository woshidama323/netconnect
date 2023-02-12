package netconnect

import (
	"fmt"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	_, err := NewWebsocketInfra("wss://bsc-mainnet.nodereal.io/ws/v1/86143292859b49279bffe736e6981bfa")
	if err != nil {
		fmt.Println("new failed err:", err)
	}

	time.Sleep(100 * time.Second)

}
