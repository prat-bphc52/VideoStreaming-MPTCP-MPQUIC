package main

import (
	"fmt"
	"time"

	utils "./utils"
	config "./config"
	quic "github.com/lucas-clemente/quic-go"
	"gocv.io/x/gocv"
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

	buffer := make([]byte, config.BUFFER_SIZE)


    // var frame gocv.Mat
    var rows = -1
    var cols = -1
	
	window := gocv.NewWindow("Output")
  
    var dimens = make([]byte, 4)
    stream.Read(dimens)
	rows = int(dimens[1]) << 8 + int(dimens[0])
	cols = int(dimens[3]) << 8 + int(dimens[2])
	pl("Video Dimensions : ", rows, " x ", cols)
	var data = make([]byte, 3*rows*cols)
    var dataind = 0

	var count = 0
	for ;count<config.MAX_FRAMES;{
		var limit = config.BUFFER_SIZE
		if limit+dataind >= len(data){
			limit = len(data)-dataind
			var temp = make([]byte, limit)
			stream.Read(temp)
			copy(data[dataind:],temp)
			count++
			dataind = 0
			img,err := gocv.NewMatFromBytes(rows,cols,gocv.MatTypeCV8UC3,data)
			utils.HandleError(err)
			window.IMShow(img)
			window.WaitKey(1)
		} else{
			stream.Read(buffer)
			copy(data[dataind:],buffer)
			dataind = dataind+limit
		}
	}

	elapsed := time.Since(start)
	pl("\nEnding video transmission, Duration: ", elapsed, " Frames Captured ", count)
	stream.Close()
	stream.Close()
	pl("\n\nThank you!")
}
