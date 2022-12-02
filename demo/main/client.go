package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "locaohose:6379")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("conn 成功", conn)
	//conn, err := net.Dial("tcp", "127.0.0.1:8888")
	//if err != nil {
	//	fmt.Println("Dial err=", err)
	//	return
	//}
	//fmt.Println("成功conn=", conn)
	////fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("reader.ReadString err=", err)
	//	}
	//	line = strings.Trim(line, "\r\n")
	//	if line == "exit" {
	//		fmt.Println("客户端退出...")
	//		break
	//	}
	//	_, err = conn.Write([]byte(line + "\n"))
	//	if err != nil {
	//		fmt.Println("conn.Write err=", err)
	//	}
	//	//fmt.Printf("客户端发送了%v个字节的数据", n)
	//}

}
