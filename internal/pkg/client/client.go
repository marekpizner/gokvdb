package main

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 8000, "The port to connect to; defaults to 8000.")

type apiResponse struct {
	Reply string
	Item  string
	Items []string
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
	time.Sleep(1)
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
		dec := gob.NewDecoder(conn)
		apiRes := &apiResponse{}
		err := dec.Decode(apiRes)
		fmt.Println(err)
		fmt.Printf("*** %+v \n", apiRes)
	}
}
