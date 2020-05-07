package main

import (
	"fmt"
	"time"

	utils "./utils"
	config "./config"
	quic "github.com/lucas-clemente/quic-go"
	// "gocv.io/x/gocv"
)

const addr = "0.0.0.0:" + config.PORT
var pl = fmt.Println
var p = fmt.Print

func main() {

	quicConfig := &quic.Config{
		CreatePaths: true,
	}

	pl("Attaching to: ", addr)
	listener, err := quic.ListenAddr(addr, utils.GenerateTLSConfig(), quicConfig)
	utils.HandleError(err)

	pl("Server listening...")

	sess, err := listener.Accept()
	utils.HandleError(err)
	stream, err := sess.AcceptStream()
	utils.HandleError(err)

	pl("Broadcasting incoming video stream...")
	defer stream.Close()
	
	time.Sleep(10*time.Millisecond)
	start := time.Now()

	buffer := make([]byte, config.BUFFERSIZE)


    // var frame gocv.Mat
    var rows = -1
    var cols = -1

    var dimens = make([]byte, 4)
    stream.Read(dimens)
	pl("Reading dimensions ", dimens)
	rows = int(dimens[1]) << 8 + int(dimens[0])
	cols = int(dimens[3]) << 8 + int(dimens[2])
	pl("Rows ", rows, " Cols ", cols)
	var data = make([]byte, 3*rows*cols)
    var dataind = 0

	var count = 0
	for ;count<config.MAX_FRAMES;{
		var limit = config.BUFFERSIZE
		if limit+dataind >= len(data){
			limit = len(data)-dataind
			var temp = make([]byte, limit)
			stream.Read(temp)
			copy(data[dataind:],temp)
			count++
			pl("\n\nReceived frame ", count, "  size of last buffer ", limit)
			pl("first 10 bytes are ", data[:10])
			pl("last 10 bytes are ", data[len(data)-10:])
			dataind = 0
		} else{
			stream.Read(buffer)
			copy(data[dataind:],buffer)
			dataind = dataind+limit
		}
	}

	elapsed := time.Since(start)
	pl("\n Ending video transmission, Duration: ", elapsed)
	time.Sleep(2 * time.Second)
	stream.Close()
	stream.Close()
}
