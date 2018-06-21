// A simple logging example. As a logging service a natcat utility can be used
// nc -luk 1902

package main

import (
	"log"
	"net"

	"time"
)

func main() {

	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "test log ", f)

	logger.Println("regular log message")
	logger.Panicln("Panicln message")

}
