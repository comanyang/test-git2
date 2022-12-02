package main

import (
	"fmt"
	"go_code/demo/customerMange/model"
	"go_code/demo/customerMange/service"
)

type CustomerView struct {
	//定义必要的字段
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *CustomerView) List() {
	//首先获取到当前所有的客户信息
	customers := this.customerService.List()
	//显示
	fmt.Println("----------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t")  //这样太笨拙
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成-------------")
}

// 得到用户的输入信息，构建新的客户，并完成添加
func (this *CustomerView) add() {
	fmt.Println("----------添加客户-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("Email：")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//ID号没有让用户输入，ID号唯一的，由系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("----------添加客户成功------------")
	} else {
		fmt.Println("----------添加客户失败------------")
	}

}

// 得到用户的输入信息，删除id对应的客户
func (this *CustomerView) Delete() {
	fmt.Println("-----------删除客户----------")
	fmt.Println("请选择待删除客户编号（-1）退出")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" || choice == "Y" {
		//调用customerService 的Delete 方法
		if this.customerService.Delete(id) {
			fmt.Println("删除完成")
		} else {
			fmt.Println("删除失败，输入的ID号不存在，请重新输入")
		}
	}

}

// 显示主菜单
func (this *CustomerView) ShowMenu() {
	for {
		fmt.Println("\n----------------客户信息管理软件----------------")
		fmt.Println("                  1 添加客户                    ")
		fmt.Println("                  2 修改客户                    ")
		fmt.Println("                  3 删除客户                    ")
		fmt.Println("                  4 客户列表                    ")
		fmt.Println("                  5 退   出                    ")
		fmt.Print("请选择1-5：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("添加客户")
		case "3":
			this.Delete()
		case "4":
			this.List()
		case "5":
			fmt.Println("添加客户")
			this.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理系统")
}
func main() {
	fmt.Println("OK!")
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	//这里完成对customerView结构体的customerService 字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.ShowMenu()
}
