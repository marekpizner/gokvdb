package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 8000, "The port to connect to; defaults to 8000.")

type apiResponse struct {
	reply string
	item  string
	items []string
}

func main() {
	flag.Parse()

	dest := *host + ":" + strconv.Itoa(*port)
	fmt.Printf("Connecting to %s...\n", dest)

	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}
	defer conn.Close()

	go readConnection(conn)

	msg := "/SET key1 value1\n"
	fmt.Print(msg)
	setAndGetData(conn, msg)

	msg = "/GET key1\n"
	fmt.Print(msg)
	setAndGetData(conn, msg)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		setAndGetData(conn, text)
	}
}

func setAndGetData(conn net.Conn, msg string) {
	conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, err := conn.Write([]byte(msg))
	fmt.Fprint(conn, msg)
	if err != nil {
		fmt.Println("Error writing to stream.")
	}
}

func readConnection(conn net.Conn) {
	for {
		buf := make([]byte, 2048)
		scanner, _ := conn.Read(buf)

		var message apiResponse
		json.Unmarshal(buf[:scanner], &message)
		fmt.Println(">: ", message.reply, message.item)
		// for {
		// 	ok := scanner.Scan()
		// 	text := scanner.Text()

		// 	command := handleCommands(text)
		// 	if !command {
		// 		fmt.Printf("\b\b** %s\n> ", text)
		// 	}

		// 	if !ok {
		// 		fmt.Println("Reached EOF on server connection.")
		// 		return
		// 	}
		// }
	}
}

func handleCommands(text string) bool {
	r, err := regexp.Compile("^%.*%$")
	if err == nil {
		if r.MatchString(text) {

			switch {
			case text == "%quit%":
				fmt.Println("\b\bServer is leaving. Hanging up.")
				os.Exit(0)
			}

			return true
		}
	}

	return false
}
