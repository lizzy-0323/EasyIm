package main

import (
	"bufio"
	"fmt"
	"go-im/pkg/protocol/pb"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var addr string

func NewCli() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "chat",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runChat()
		},
	}
	rootCmd.Flags().StringVar(&addr, "addr", "ws://localhost:8002", "WebSocket server address")
	return rootCmd
}

func registerDevice() {
	// TODO：
}

func runChat() {
	// Connect to the WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	defer c.Close()
	netDone := make(chan error)

	// Handle incoming messages
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				if err == websocket.ErrCloseSent {
					netDone <- err
				}
				break
			}
			fmt.Println(string(message))
		}
	}()

	fmt.Println("Connected to", addr)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "quit" {
			break
		}
		pbFormatMsg := SignInMsg(line)

		// Marshal to binary
		msg, err := proto.Marshal(pbFormatMsg)
		if err != nil {
			log.Println("marshal error: ", err)
			break
		}

		// Write to WebSocket
		err = c.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	select {
	case <-netDone:
		fmt.Println("Connection closed")
		c.Close()
		return
	default:
	}
}

func SignInMsg(data string) *pb.Input {
	// Using test data
	signInPbMsg := &pb.SignInInput{
		DeviceId: 1,
		UserId:   1,
		Token:    "asdasda",
	}
	signInData, _ := proto.Marshal(signInPbMsg)
	return &pb.Input{
		Type:      pb.PackageType_PT_SIGN_IN,
		Data:      []byte(signInData),
		RequestId: 1,
	}
}

func main() {
	cli := NewCli()
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
