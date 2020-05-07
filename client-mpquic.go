package main

import (
	"crypto/tls"
	"fmt"
	"time"
	"encoding/binary"

	utils "./utils"
	config "./config"
	quic "github.com/lucas-clemente/quic-go"
	"gocv.io/x/gocv"
)

const addr = "127.0.0.1:"+config.PORT
var pl = fmt.Println
var p = fmt.Print

func main() {

	quicConfig := &quic.Config{
		CreatePaths: true,
	}


	pl("Trying to connect to: ", addr)
	sess, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, quicConfig)
	utils.HandleError(err)

	stream, err := sess.OpenStream()
	utils.HandleError(err)

	pl("Connection established with server successfully...Starting Video stream")
	defer stream.Close()

	webcam, _ := gocv.VideoCaptureDevice(0)
	img := gocv.NewMat()

	start := time.Now()

	webcam.Read(&img)

	pl("Rows ", img.Rows(), " Cols ", img.Cols())
    var dimens = make([]byte, 4)
	var bs = make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(img.Rows()))
	copy(dimens[0:],bs)
	binary.LittleEndian.PutUint16(bs, uint16(img.Cols()))
	copy(dimens[2:],bs)
	stream.Write(dimens)
	pl("Sending dimensions ", dimens)


	for i:=1;i<=config.MAX_FRAMES;i++{
		webcam.Read(&img)
		pl("\n\nSending frame ", i)
		var b = img.ToBytes()
		pl("first 10 bytes are ", b[:10])
		pl("last 10 bytes are ", b[len(b)-10:])
		for ind:=0;ind<len(b);{
			var end = ind+config.BUFFERSEND
			if end>len(b){
				end = len(b)
			}
			stream.Write(b[ind:end])
			ind = end
		}
		
	}
	stream.Write([]byte{0,0,0,0})
	webcam.Close()

	elapsed := time.Since(start)
	pl("\nEnding Video Stream, Duration : ", elapsed)

	stream.Close()
	stream.Close()
	time.Sleep(2 * time.Second)
	pl("\n\nThank you!")
}
