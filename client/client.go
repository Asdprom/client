package client //client.go

import (
	"Server/converter"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// Client stores an information about client, such as name (UTF16) and socket (net.Conn)
type Client struct {
	Name   []byte
	Socket net.Conn
}

// RecieveString recievces Utf16 string from socket and converts it to utf8 string
// returns string in utf8 and utf16
func (client Client) RecieveString() (str string, temp []byte) {
	arr := make([]byte, 4)
	_, err := client.Socket.Read(arr)
	checkError(err)
	res := converter.ReadInt32(arr)

	temp = make([]byte, res)
	_, err = client.Socket.Read(temp)
	checkError(err)
	str, err = converter.DecodeUTF16(temp)
	return str, temp
}

// SendMessage отправляет сначала флаг 0, а затем сообщение
func (server Client) SendMessage(message string) {
	barr, err := converter.DecodeUTF8(message)
	checkError(err)
	arr := make([]byte, 4)
	binary.LittleEndian.PutUint32(arr, 0)
	server.Socket.Write(arr)
	binary.LittleEndian.PutUint32(arr, uint32(len(barr)))
	server.Socket.Write(arr)
	server.Socket.Write(barr)
}
func (server Client) SendSimpleString(str string) {

	barr, err := converter.DecodeUTF8(str)
	arr := make([]byte, 4)
	binary.LittleEndian.PutUint32(arr, uint32(len(barr)))
	checkError(err)
	server.Socket.Write(arr)
	server.Socket.Write(barr)
}
func (client Client) SendFlag(flag uint32) {
	arr := make([]byte, 4)
	binary.LittleEndian.PutUint32(arr, flag)
	client.Socket.Write(arr)
}
func (client Client) RecieveFlag() (flag int32) {
	arr := make([]byte, 4)
	_, err := client.Socket.Read(arr)
	checkError(err)
	flag = converter.ReadInt32(arr)
	return flag
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*func (client Client) HandleClient(clients *[]Client) {
	for {
		flag := client.RecieveFlag()
		switch flag {
		case 0:
			message, utf16mess := client.RecieveString()
			fmt.Println(" Message =  ", message)

			for i := 0; i < len(*clients); i++ {
				fmt.Println(" Sending to =  ", (*clients)[i].Name)
				client.SendString((*clients)[i], utf16mess)
			}
		case 1:
			fmt.Println("Fuck you, asshole.")
			client.Socket.Close()
			return
		}
	}
}*/
