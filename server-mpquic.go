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

	var count = 0

	count = 0
    // var frame gocv.Mat
    var rows1 = -1
    var rows2 = -1
    var cols1 = -1
    var cols2 = -1
    var rows = -1
    var cols = -1

    var data []byte
    var dataind = 0

    // var i = 0
    outer:
	  for{
		stream.Read(buffer)
		// i++
		// pl("Iteration : ", i)
		for ind:=0;ind<len(buffer);{
			if rows1 == -1{
				pl(buffer[ind:ind+4])
				rows1 = int(buffer[ind])
				ind++
			}
			if rows2 == -1{
				rows2 = int(buffer[ind])
				ind++
				rows = rows2 << 8 + rows1
			}
			if cols1 == -1{
				cols1 = int(buffer[ind])
				ind++
			}
			if cols2 == -1{
				cols2 = int(buffer[ind])
				ind++
				cols = cols2 << 8 + cols1
				pl("Rows ", rows, " Cols ", cols)
				if rows!=0 && rows != 480{
					break outer
				}
				data = make([]byte, 3*rows*cols)
				
			}

			var limit = len(buffer) - ind
			if limit+dataind>len(data){
				limit = len(data)-dataind
			}
			if limit!=0{
				copy(data[dataind:],buffer[ind:(ind+limit)])
			}
			dataind = dataind+limit
			ind = ind+limit

			if dataind==len(data){
				count++
				pl("Received frame, Total Count : ", count)
				pl("last 10 bytes are ", data[len(data)-10:len(data)])
				rows = 0
				cols = 0
				rows1 = -1
				rows2 = -1
				cols1 = -1
				cols2 = -1
				dataind = 0
				// continue
			}
		}
	  }

	elapsed := time.Since(start)
	pl("\n Ending video transmission, Duration: ", elapsed)
	time.Sleep(2 * time.Second)
	stream.Close()
	stream.Close()
}
