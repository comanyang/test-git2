package utils11

import "fmt"

type FamilyAccount struct {
	//声明一个字段，保存接收用户输入的选项
	key  string
	loop bool
	//定义账号余额
	balance float64
	//每次收支的金额
	money float64
	// 每次收支的说明
	note string
	//收支的详情使用字符串来记录
	details string
	//定义一个变量，判断是否有收支
	flag bool
}

// 使用工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说  明",
		flag:    false,
	}
}

// 将显示明细封装到一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("\n------------------显示收支明细------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有任何收支，来一笔吧！")
	}
}

// 将登记收入封装到一个方法
func (this *FamilyAccount) income() {
	this.flag = true
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v    \t%v          \t%v", this.balance, this.money, this.note)

}

// 将登记支出封装到一个方法
func (this *FamilyAccount) pay() {
	this.flag = true
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.balance-this.money < 0 {
		fmt.Println("支出金额超出余额")
		//break
	}
	this.balance -= this.money
	fmt.Print("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v    \t%v          \t%v", this.balance, this.money, this.note)

}

// 将登记退出封装到一个方法
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
	//break
}

// 给该结构体绑定相应的方法
// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {

		fmt.Println("\n----------------家庭收支记账软件----------------")
		fmt.Println("                  1 收支明细                    ")
		fmt.Println("                  2 登记收入                    ")
		fmt.Println("                  3 登记支出                    ")
		fmt.Println("                  4 退出软件                    ")
		fmt.Print("请选择1-4：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if this.loop == false {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}

var herroName string = "宋江"
