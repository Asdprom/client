package initClient

import (
	"clientWidget2/client"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/therecipe/qt/widgets"
)

var isLoggedIn bool

const (
	OkFlag    = 4
	ErrorFlag = 5
	RegFlag   = 2
	LoginFlag = 3
)

type ChatClient struct {
	Server     client.Client
	isLoggedIn bool
}

func (chatClient *ChatClient) Init(port string) {
	var err error
	chatClient.Server.Socket, err = net.Dial("tcp", ":"+port)
	chatClient.isLoggedIn = false
	checkError(err)
}

func (chatClient *ChatClient) Login(username, password string) bool {
	chatClient.Server.SendFlag(LoginFlag)
	chatClient.Server.SendSimpleString(username)
	chatClient.Server.SendSimpleString(password)
	response := chatClient.Server.RecieveFlag()
	if response == OkFlag {
		chatClient.isLoggedIn = true
		return true
	}
	return false
}

func (chatClient *ChatClient) Registration(username, password string) bool {
	chatClient.Server.SendFlag(RegFlag)
	chatClient.Server.SendSimpleString(username)
	chatClient.Server.SendSimpleString(password)
	response := chatClient.Server.RecieveFlag()
	if response == OkFlag {
		return true
	}
	return false
}

/*func InitClient() {
	service := ":1213"

	isLoggedIn = false
	username := "User"
	var err error

	var server client.Client
	server.Socket, err = net.Dial("tcp", service)
	checkError(err)
	reader := bufio.NewReader(os.Stdin)
	for {
		if isLoggedIn == true {
			break
		}
		var flag int
		fmt.Scan(&flag)
		switch flag {
		case 0: // Registration0
			server.SendFlag(RegFlag)
			server.SendSimpleString(username)
			server.SendSimpleString("tipidor")
			response := server.RecieveFlag()
			if response == OkFlag {
				fmt.Println("Registration was suckassful!")
			}
		case 1: // Login
			server.SendFlag(LoginFlag)
			server.SendSimpleString(username)
			server.SendSimpleString("password")
			response := server.RecieveFlag()
			if response == OkFlag {
				fmt.Println("Login was suckassful!")
				isLoggedIn = true
			} else {
				fmt.Println("No such combination login/password")
			}

		case 2: // ??
			break
		}
	}
	go waitForResponse(server)

	for {
		fmt.Print("Me: ")
		message, _ := reader.ReadString('\n')

		server.SendMessage(message)
	}

}*/
func (chat ChatClient) SendMessage(message string) {
	chat.Server.SendMessage(message)
}

func splitMessage(message string) (name string, mess string) {
	arr := strings.SplitAfter(message, "%&%")
	arr[0] = strings.Trim(arr[0], "%&%")

	return arr[0], arr[1]
}
func WaitForResponse(server client.Client, chat *widgets.QTextEdit) {
	for {
		flag := server.RecieveFlag()
		switch flag {
		case 0:
			message, _ := server.RecieveString()
			name, message := splitMessage(message)
			previousText := chat.ToPlainText()
			previousText += "\n" + name + ": " + message
			chat.SetText(previousText)
		case 1:
			return
		case 3:

		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
