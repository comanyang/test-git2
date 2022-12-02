package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/demo/client/utils"
	"go_code/demo/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userid int, userpwd string, username string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userid
	registerMes.User.UserPwd = userpwd
	registerMes.User.UserName = username
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误err=", err)
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn)  err=", err)
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data) err=", err)
		return
	}
	fmt.Println("RegisterResMes.Code =", registerResMes.Code)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
func (this *UserProcess) Login(userid int, userpwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userid
	loginMes.UserPwd = userpwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err=", err)
	}
	fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))
	n, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
	}
	rf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = rf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn)  err=", err)
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data) err=", err)
		return
	}
	fmt.Println("loginResMes.Code =", loginResMes.Code)
	if loginResMes.Code == 200 {
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			if v == userid {
				continue
			}
			fmt.Println("用户id:", v)
		}
		fmt.Println("\n\n")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
