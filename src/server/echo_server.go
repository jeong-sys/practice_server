package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if nil != err {
		log.Println(err.Error())
	}
	defer listener.Close()

	for {
		// 클라이언트 요청 대기
		conn, err := listener.Accept()

		// 예외처리
		if err != nil {
			log.Println(err.Error())
			continue
		}
		// defer conn.Close()

		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		// err 처리
		n, err := conn.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				log.Println("err(eof):", err)
				break
			}
			log.Println("err:", err)
			break
		}

		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])

			if err != nil {
				log.Println("err:", err)
				break
			}
		}
	}
}
