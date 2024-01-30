package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	addr  = "0.0.0.0:12345"
	proto = "tcp4"
)

func main() {

	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go client(conn)
	}

}

func client(conn net.Conn) {

	sentences := []string{
		"Don't communicate by sharing memory, share memory by communicating",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.After(3 * time.Second):
				randomSentence := getRandomSentence(sentences)
				conn.Write([]byte(randomSentence + "\n"))
			case <-done:
				return
			}
		}
	}()

}

// Функция для выбора случайной поговорки из списка
func getRandomSentence(sentences []string) string {

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(sentences))
	return sentences[index]

}
