package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gordonklaus/portaudio"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("missing required argument:  output file name")
	// 	return
	// }
	fmt.Println("Recording.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	nSamples := 0

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 999)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	for {
		chk(stream.Read())
		// chk(binary.Write(f, binary.BigEndian, in))
		nSamples += len(in)
		fmt.Println("writing", nSamples)
		select {
		case <-sig:
			return
		default:
		}
	}
	fmt.Println("writing  end~")
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
