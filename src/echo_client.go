package main
 
import (
    "fmt"
    "log"
    "net"
)
 
func main() {
    conn, err := net.Dial("tcp", "192.168.50.140:8000")
    if nil != err {
        log.Println("err:", err)
        return
    }
 
    go func() {
        data := make([]byte, 4096)
        for {
            n, err := conn.Read(data)
            if err != nil {
                log.Println(err)
                break
            }
 
            log.Println("Server send:" + string(data[:n]))
        }
    }()
 
    for {
        var s string
        fmt.Scanln(&s)
        conn.Write([]byte(s))
    }
}