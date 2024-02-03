package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if nil != err {
		log.Println(err)
		return
	}

	go func() {
		data := make([]byte, 65536)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		var input []byte

		for {
			char, err := reader.ReadByte()
			if err != nil {
				log.Println(err)
				break
			}

			// 개행 문자를 만나면 입력 종료
			if char == '\n' {
				break
			}

			input = append(input, char)
		}

		_, err = conn.Write(input)

		if err != nil {
			log.Println("Failed to send data:", err)
			continue
		}
	}
}
