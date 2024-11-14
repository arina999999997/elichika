package shutdown

import (
	"fmt"
	"sync/atomic"
	"time"
)

// we perform "graceful" shutdown as follow:
// - Track the work count using an atomic variable, and track the shutdown using a boolean
// - When shutdown flag is set, all incoming requests are rejected.
// - Then we wait for all existing work to finish.
// - And finally shutdown
var isShutdown atomic.Bool
var connectionCount atomic.Int32
var sync = make(chan struct{})

func Shutdown() {
	if isShutdown.Load() {
		fmt.Println("Warning: shutdown is already triggered")
		return
	}
	isShutdown.Store(true)
}

func IsShutdown() bool {
	return isShutdown.Load()
}

func StartConnection() {
	connectionCount.Add(1)
}

func FinishConnection() {
	connectionCount.Add(-1)
}

func WaitForFinish() {
	if !IsShutdown() {
		panic("shutdown flag not set, waiting for it to be set")
	}
	for {
		if connectionCount.Load() == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("No outstanding connection, stopping server")
}

func SendFinalSignal() {
	sync <- struct{}{}
}

func ReceiveFinalSignal() {
	<-sync
}
