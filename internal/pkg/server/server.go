package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/khan745/gokvdb/internal/pkg/command"
	"github.com/khan745/gokvdb/internal/pkg/storage/memory"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8000, "The port to listen on; default is 8000.")

var strg = memory.New()
var parser = command.NewParser(strg)

func main() {
	flag.Parse()

	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go handleConnection(conn)
		go sendMessage(conn)
	}
}

func sendMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}
		handleMessage(scanner.Text(), conn)
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) {
	if message != "" {
		fmt.Println("> " + message)
	}

	if len(message) > 0 && message[0] == '/' {
		switch {
		case message == "/time":
			resp := "It is " + time.Now().String() + "\n"
			fmt.Print("< " + resp)
			conn.Write([]byte(resp))

		case message == "/quit":
			fmt.Println("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			fmt.Println("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			os.Exit(0)

		default:
			message = strings.Replace(message, "/", "", -1)
			cmd, args, _ := parser.Parse(message)
			// TODO:
			// 	- create api interface for communication
			// 	- pozri na autogen rozhrania https://github.com/gogo/protobuf https://github.com/golang/protobuf
			//  - minimock https://github.com/gojuno/minimock
			if cmd != nil {
				res := cmd.Execute(args...)
				fmt.Println("Comand, arguments", cmd, args, res, message)
				switch t := res.(type) {
				case command.OkResult:
					conn.Write([]byte("Command execute successfully\n"))
				case command.StringReply:
					conn.Write([]byte(t.Value + "\n"))
				case command.ErrResult:
					conn.Write([]byte("Something wen wrong with executing command\n"))
				default:
					fmt.Println("WTF")
				}

			} else {
				conn.Write([]byte("Unrecognized command.\n"))
			}
		}
	}
}
