package main

import (
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"log"
	"os"
	"strconv"
	"time"
)

func help() {
	fmt.Println("kk-uuid <name> <0.0.0.0:8080>")
}

const twepoch = int64(1424016000000000)

func milliseconds() int64 {
	return time.Now().UnixNano() / 1e3
}

func main() {

	var args = os.Args
	var name string = ""
	var address string = ""

	if len(args) > 2 {
		name = args[1]
		address = args[2]
	} else {
		help()
		return
	}

	var id = milliseconds()

	var replay func(message *kk.Message) bool = nil

	replay, _ = kk.TCPClientConnect(name, address, map[string]interface{}{"exclusive": true, "title": "uuid"}, func(message *kk.Message) {
		if message.Method == "REQUEST" {
			id = id + 1
			var v = kk.Message{message.Method, name, message.From, "text", []byte(strconv.FormatInt(id, 10))}
			replay(&v)
		} else {
			var v = kk.Message{"NOIMPLEMENT", message.To, message.From, "", []byte("")}
			log.Println(v)
			replay(&v)
		}
	})

	kk.DispatchMain()

}
