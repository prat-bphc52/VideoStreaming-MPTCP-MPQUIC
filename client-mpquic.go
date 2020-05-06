package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	utils "./utils"
	config "./config"
	quic "github.com/lucas-clemente/quic-go"
)

const addr = "127.0.0.1:"+config.PORT
const threshold = 5 * 1024  // 5KB
//TODO: set this threshold dynamically, based on network conditions


func main() {

	quicConfig := &quic.Config{
		CreatePaths: true,
	}

	fileToSend := "sample.txt"

	fmt.Println("Sending File: ", fileToSend)


	file, err := os.Open(fileToSend)
	utils.HandleError(err)

	fileInfo, err := file.Stat()
	utils.HandleError(err)

	if fileInfo.Size() <= threshold {
		quicConfig.CreatePaths = false
		fmt.Println("File is small, using single path only.")
	} else {
		fmt.Println("file is large, using multipath now.")
	}
	file.Close()

	fmt.Println("Trying to connect to: ", addr)
	sess, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, quicConfig)
	utils.HandleError(err)

	fmt.Println("session created: ", sess.RemoteAddr())

	stream, err := sess.OpenStream()
	utils.HandleError(err)

	fmt.Println("stream created...")
	fmt.Println("Client connected")
	sendFile(stream, fileToSend)
	time.Sleep(2 * time.Second)

}

func sendFile(stream quic.Stream, fileToSend string) {
	fmt.Println("A client has connected!")
	defer stream.Close()

	file, err := os.Open(fileToSend)
	utils.HandleError(err)

	fileInfo, err := file.Stat()
	utils.HandleError(err)

	fileSize := utils.FillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := utils.FillString(fileInfo.Name(), 64)

	fmt.Println("Sending filename and filesize!")
	stream.Write([]byte(fileSize))
	stream.Write([]byte(fileName))

	sendBuffer := make([]byte, config.BUFFERSIZE)
	fmt.Println("Start sending file!\n")

	var sentBytes int64
	start := time.Now()

	for {
		sentSize, err := file.Read(sendBuffer)
		if err != nil {
			break
		}

		stream.Write(sendBuffer)
		if err != nil {
			break
		}


		sentBytes += int64(sentSize)
		fmt.Printf("\033[2K\rSent: %d / %d", sentBytes, fileInfo.Size())
	}
	elapsed := time.Since(start)
	fmt.Println("\nTransfer took: ", elapsed)

	stream.Close()
	stream.Close()
	time.Sleep(2 * time.Second)
	fmt.Println("\n\nFile has been sent, closing stream!")
	return
}
