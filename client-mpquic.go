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

	for i:=1;i<=config.MAX_FRAMES;i++{
		webcam.Read(&img)
		pl("Sending frame ", i)
		pl("Rows ", img.Rows(), " Cols ", img.Cols())
		var b = img.ToBytes()
		b = append(b, 0,0,0,0)
		copy(b[4:],b[0:])
		var bs = make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, uint16(img.Rows()))
		copy(b[0:],bs)
		binary.LittleEndian.PutUint16(bs, uint16(img.Cols()))
		copy(b[2:],bs)
		pl("last 10 bytes are ", b[len(b)-10:len(b)])
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
