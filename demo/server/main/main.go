package main

import (
	"fmt"
	"go_code/demo/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯的协程错误 err=", err)
		return
	}
}
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来链接服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)

	}

}
