package main

import (
	"flag"
	"fmt"
	"github.com/anchor/picolog"
	"github.com/fractalcat/emogo"
	zmq "github.com/pebbe/zmq4"
	"os"
)

var Logger *picolog.Logger

func readFrames(e *emogo.EmokitContext, out chan *emogo.EmokitFrame) {
	for {
		f, err := e.WaitGetFrame()
		if err != nil {
			fmt.Printf("error reading frame: %v", err)
			return
		}
		out <- f
	}
}

func main() {
	listen := flag.String("listen", "tcp://*:9424", "ZMQ URI to listen on.")
	flag.Parse()
	Logger = picolog.NewLogger(picolog.LogDebug, "insenced", os.Stdout)
	eeg, err := emogo.NewEmokitContext()
	if err != nil {
		Logger.Fatalf("Could not initialize emokit context: %v", err)
	}
	Logger.Debugf("EEG initialized.")
	Logger.Debugf("Detected %d EEG devices connected.", eeg.Count())
	sock, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		Logger.Fatalf("Could not create ZMQ socket: %v", err)
	}
	err = sock.Bind(*listen)
	if err != nil {
		Logger.Fatalf("Could not bind to %s: %v", listen, err)
	}
	frameChan := make(chan *emogo.EmokitFrame, 0)
	readFrames(eeg, frameChan)
	for {
		f := <-frameChan
		Logger.Debugf("Got frame. Sending.")
		_, err := sock.SendBytes(f.Raw(), 0)
		if err != nil {
			fmt.Printf("Error sending raw frame: %v", err)
		}
	}
}
