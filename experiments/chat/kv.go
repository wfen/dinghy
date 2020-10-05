package main

/*

import (
	"bufio"
	"log"
	"strings"
	"sync"
)

package main

import (
"bufio"
"fmt"
"log"
"net"
"strings"
"sync"
)

var keys = make(map[string]string)
var lock = sync.Mutex{}

const HELP_MSG = `available commands:
  - get [key]
  - put [key] [value]
  - list
`

func get(key string) []byte {
	lock.Lock()
	defer lock.Unlock()
	log.Printf("getting %v from keys", key)
	k, ok := keys[key]
	if !ok {
		return []byte("nil\n")
	}
	return []byte(k + "\n")
}

func put(key string, value string) []byte {
	lock.Lock()
	defer lock.Unlock()
	log.Printf("setting %v to %v", key, value)
	keys[key] = value
	return []byte(fmt.Sprintf("successfully put %s with value of %s!\n", key, value))
}

func list() []byte {
	lock.Lock()
	defer lock.Unlock()
	log.Printf("keys: %+v", keys)
	msg := ""
	for key, value := range keys {
		msg += fmt.Sprintf("key: %v, value: %v\n", key, value)
	}
	return []byte(msg)
}

func readInput(conn net.Conn) {
	for {
		input, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("unable to read input")
		}

		args := strings.Split(strings.TrimSpace(input), " ")
		log.Printf("args: %v", args)

		switch args[0] {
		case "get":
			if len(args) < 2 {
				conn.Write([]byte("what key do you want to get?\n"))
			}
			key := strings.TrimSpace(args[1])
			conn.Write(get(key))
		case "put":
			if len(args) < 3 {
				conn.Write([]byte("error: please supply two args: a key and a value\n"))
			} else {
				key := strings.TrimSpace(args[1])
				value := strings.TrimSpace(args[2])
				conn.Write(put(key, value))
			}
		case "list":
			conn.Write(list())
		default:
			conn.Write([]byte(HELP_MSG))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp","127.0.0.1:4444")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		client := conn.RemoteAddr()
		log.Printf("new client: %v", client)
		go readInput(conn)
	}

}
*/
