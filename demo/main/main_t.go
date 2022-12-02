package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//)
//
//func test03() {
//	str := ""
//	for i := 0; i < 100000; i++ {
//		str += "hello" + strconv.Itoa(i)
//	}
//}
//func fbn(n int) []uint64 {
//	fbnSlice := make([]uint64, n)
//	fbnSlice[0] = 1
//	fbnSlice[1] = 1
//	for i := 2; i < n; i++ {
//		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
//
//	}
//	return fbnSlice
//}
//func modifyUser(users map[string]map[string]string, name string, nickname string) {
//	if users[name] != nil {
//		users[name]["psw"] = "888888"
//	} else {
//		users[name] = make(map[string]string, 1)
//		users[name]["nickname"] = nickname
//		users[name]["psw"] = "888888"
//	}
//}
//
//type Cat struct {
//	Name  string
//	Age   int
//	Color string
//	Hobby string
//}
//type Person struct {
//	Name   string
//	Age    int
//	Scores [5]float64
//	ptr    *int
//	slice  []int
//	map1   map[string]string
//}
//type monster struct {
//	Name string
//	Age  int
//}
//type A struct {
//	Name string
//	age  int
//}
//
//func (this *A) SayOk() {
//	fmt.Println("A SayOk", this.Name)
//}
//func (this *A) hello() {
//	fmt.Println("A hello", this.Name)
//}
//
//type B struct {
//	A
//	//	Name string
//}
//type D struct {
//	a A
//}
//type Monster struct {
//	Name string
//	Age  int
//}
//type E struct {
//	Monster
//	int
//	n int
//}
//type Usb interface {
//	Start()
//	Stop()
//}
//
////	type Usb2 interface {
////		Start()
////		Stop()
////		Test()
////	}
////
//// type Phone struct {
//// }
//// type Camera struct {
//// }
////
////	func (this Camera) Start() {
////		fmt.Println("相机2开始工作")
////	}
////
////	func (this Camera) Stop() {
////		fmt.Println("相机2停止工作")
////	}
////
////	func (this Phone) Start() {
////		fmt.Println("手机2开始工作")
////	}
////
////	func (this Phone) Stop() {
////		fmt.Println("手机2222停止工作")
////	}
////
//// type Computer struct {
//// }
////
////	func (this Computer) Working(usb Usb2) {
////		usb.Start()
////		usb.Stop()
////	}
////
////	type Ainterface interface {
////		Say()
////	}
////
////	type Stu struct {
////		Name string
////	}
////
////	func (this Stu) Say() {
////		fmt.Println("Say hello")
////	}
////
//// type integer int
////
////	func (this integer) Say() {
////		fmt.Println("integer Say hello i=", this)
////	}
////
////	type Binterface interface {
////		Hello()
////	}
////
//// type Monster2 struct {
//// }
////
////	func (this Monster2) Hello() {
////		fmt.Println("Binterface Hello")
////	}
////
////	func (this Monster2) Say() {
////		fmt.Println("Binterface Say")
////	}
//type Binterface interface {
//	Test01()
//	Test02()
//}
//type Cinterface interface {
//	Test02()
//	Test03()
//}
//type Ainterface interface {
//	Binterface
//	Cinterface
//}
//
//// type Stu struct {
//// }
////
////	func (this Stu) Test01() {
////		fmt.Println("test01")
////	}
////
////	func (this Stu) Test02() {
////		fmt.Println("test02")
////	}
////
////	func (this Stu) Test03() {
////		fmt.Println("test03")
////	}
////
//// type T interface {
//// }
//type Hero struct {
//	Name string
//	Age  int
//}
//type HeroSlice []Hero
//
//func (this HeroSlice) Len() int {
//	return len(this)
//}
//func (this HeroSlice) Less(i, j int) bool {
//	return this[i].Age < this[j].Age
//}
//func (this HeroSlice) Swap(i, j int) {
//	temp := this[i]
//	this[i] = this[j]
//	this[j] = temp
//}
//func main() {
//	//打开一个新文件，打开文件
//	filePath := "E:\\HELLO.txt"
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		fmt.Printf("open file err=%v\n", err)
//		return
//	}
//	defer file.Close()
//	//写入5句话
//	str := " hello ,gardon\n"
//	writer := bufio.NewWriter(file)
//	for i := 0; i < 5; i++ {
//		writer.WriteString(str)
//	}
//	writer.Flush()
//	//file := "E:\\HELLO.txt"
//	//str, err := ioutil.ReadFile(file)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Printf("%s", str)
//	//file, err := os.Open("E:\\HELLO.txt")
//	//if err != nil {
//	//	fmt.Println("打开文件错误：", err)
//	//}
//	//fmt.Println("file=", file)
//	//
//	////当函数退出时，要及时关闭FILE
//	//defer file.Close() //要及时关闭file句柄，否则会有内存泄漏
//	////创建一个READER
//	//reader := bufio.NewReader(file)
//	////循环的读取文件的内容
//	//for {
//	//	str, err := reader.ReadString('\n') //读到一个换行符就结束
//	//	if err == io.EOF {                  //io.EOF表示文件的末尾
//	//		break
//	//	}
//	//	fmt.Print(str)
//	//}
//	//fmt.Println("文件读取结束。。。")
//
//	//fmt.Println("这个是面向对象的方式完成")
//	//utils.NewFamilyAccount().MainMenu()
//	//var key string
//	//loop := true
//	////定义账号余额
//	//balance := 10000.0
//	////每次收支的金额
//	//money := 0.0
//	//// 每次收支的说明
//	//note := ""
//	////收支的详情使用字符串来记录
//	//details := "收支\t账户金额\t收支金额\t说  明"
//	////定义一个变量，判断是否有收支
//	//flag := false
//	////显示主菜单
//	//for {
//	//
//	//	fmt.Println("\n----------------家庭收支记账软件----------------")
//	//	fmt.Println("                  1 收支明细                    ")
//	//	fmt.Println("                  2 登记收入                    ")
//	//	fmt.Println("                  3 登记支出                    ")
//	//	fmt.Println("                  4 退出软件                    ")
//	//	fmt.Print("请选择1-4：")
//	//	fmt.Scanln(&key)
//	//	switch key {
//	//	case "1":
//	//		fmt.Println("\n------------------显示收支明细------------------")
//	//		if flag {
//	//			fmt.Println(details)
//	//		} else {
//	//			fmt.Println("当前没有任何收支，来一笔吧！")
//	//		}
//	//	case "2":
//	//		flag = true
//	//		fmt.Print("本次收入金额：")
//	//		fmt.Scanln(&money)
//	//		balance += money
//	//		fmt.Print("本次收入说明：")
//	//		fmt.Scanln(&note)
//	//		details += fmt.Sprintf("\n收入\t%v    \t%v          \t%v", balance, money, note)
//	//	case "3":
//	//		flag = true
//	//		fmt.Print("本次支出金额：")
//	//		fmt.Scanln(&money)
//	//		if balance-money < 0 {
//	//			fmt.Println("支出金额超出余额")
//	//			break
//	//		}
//	//		balance -= money
//	//		fmt.Print("本次支出说明：")
//	//		fmt.Scanln(&note)
//	//		details += fmt.Sprintf("\n支出\t%v    \t%v          \t%v", balance, money, note)
//	//	case "4":
//	//		fmt.Println("你确定要退出吗？y/n")
//	//		choice := ""
//	//		for {
//	//			fmt.Scanln(&choice)
//	//			if choice == "y" || choice == "n" {
//	//				break
//	//			}
//	//			fmt.Println("你的输入有误，请重新输入 y/n")
//	//		}
//	//		if choice == "y" {
//	//			loop = false
//	//		}
//	//		break
//	//	default:
//	//		fmt.Println("请输入正确的选项")
//	//	}
//	//	if loop == false {
//	//		break
//	//	}
//	//}
//	//fmt.Println("你退出了家庭记账软件的使用")
//	//var intSlice = []int{1, -1, 3, 5, 10, 7, 9}
//	//sort.Ints(intSlice)
//	//fmt.Println(intSlice)
//	//var heros HeroSlice
//	//for i := 0; i < 10; i++ {
//	//	hero := Hero{
//	//		Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
//	//		Age:  rand.Intn(100),
//	//	}
//	//	heros = append(heros, hero)
//	//}
//	//for _, v := range heros {
//	//	fmt.Println(v)
//	//}
//	//sort.Sort(heros)
//	//for _, v := range heros {
//	//	fmt.Println(v)
//	//}
//	//i := 10
//	//j := 20
//	//i, j = j, i
//	//fmt.Println("i=", i, "j=", j)
//	//var s1 Stu
//	//var a Ainterface = s1
//	//a.Test01()
//	//a.Test02()
//	//a.Test03()
//	//num1 := float64(8.8)
//	//var t T = s1
//	//var t2 T = num1
//	//fmt.Println(t)
//	//fmt.Println(t2)
//	//var m1 Monster2
//	//var b Binterface = m1
//	//var a Ainterface = m1
//	//b.Hello()
//	//a.Say()
//	//var s1 Stu
//	//var a Ainterface = s1
//	//a.Say()
//	//var i1 integer = 20
//	//var b Ainterface = i1
//	//b.Say()
//	//computer := Computer{}
//	//phone := Phone{}
//	//camera := Camera{}
//	//computer.Working(phone)
//	//computer.Working(camera)
//	//var e E
//	//e.Name = "狐狸精"
//	//e.Age = 300
//	//e.int = 20
//	//e.n = 40
//	//fmt.Println(e)
//	//var d D
//	////	d.Name = "tom"
//	//d.a.Name = "tom2"
//	//var b B
//	////b.Name = "SMITH"
//	////b.age = 30
//	//b.A.Name = "jack"
//	//b.A.age = 33
//	//b.SayOk()
//	//b.hello()
//	//var person *Person = new(Person)
//	//(*person).Name = "张三"
//	//person.Age = 300
//	//fmt.Println(person)
//	//var p2 *Person = &Person{}
//	//(*p2).Name = "李四"
//	//p2.Age = 200
//	//fmt.Println(p2)
//	//p2 := Person{}
//	//fmt.Println(p2)
//
//	//定义结构体变量
//	//var p1 Person
//	//fmt.Println(p1)
//	//p1.slice = make([]int, 2)
//	//p1.slice[0] = 100
//	//p1.map1 = make(map[string]string)
//	//p1.map1["key1"] = "小花"
//	//p1.map1["key2"] = "小白"
//	//fmt.Println(p1)
//	//var monster1 monster
//	//monster1.Name = "狐狸精"
//	//monster1.Age = 100
//	//monster2 := monster1
//	//monster2.Age = 200
//	//fmt.Println(monster1)
//	//fmt.Println(monster2)
//	//var cat1 Cat
//	//cat1.Name = "小花"
//	//cat1.Age = 3
//	//cat1.Color = "花色"
//	//cat1.Hobby = "吃<・)))><<"
//	//fmt.Printf("cat1的地址%p\n", &cat1)
//	//fmt.Printf("猫的名字是%v\n", cat1.Name)
//	//fmt.Printf("猫的爱好是%v", cat1.Hobby)
//	//users := make(map[string]map[string]string, 10)
//	//users["smith"] = make(map[string]string, 2)
//	//users["smith"]["nickname"] = "昵称1"
//	//users["smith"]["psw"] = "66"
//	//fmt.Println(users)
//	//modifyUser(users, "smith", "昵称1~")
//	//modifyUser(users, "tom", "昵称2~")
//	//modifyUser(users, "mary", "昵称3~")
//	//fmt.Println(users)
//	//var scores [3][5]float64
//	//for i := 0; i < len(scores); i++ {
//	//	for j := 0; j < len(scores[i]); j++ {
//	//		fmt.Println("请输入某个同学的分数：")
//	//		fmt.Scanln(&scores[i][j])
//	//	}
//	//}
//	//fmt.Println(scores)
//	//
//	//for i := 0; i < len(scores); i++ {
//	//	total_bj := 0.0
//	//	for j := 0; j < len(scores[i]); j++ {
//	//		total_bj += scores[i][j]
//	//	}
//	//	fmt.Printf("第%v个班级的平均分%v:", i+1, total_bj/float64(len(scores[i])))
//	//}
//	//var arr [2][3]int
//	//arr[0][2] = 10
//	//fmt.Printf("arr0的地址:%p\n", &arr[0])
//	//fmt.Printf("arr0的地址:%p\n", &arr[0][0])
//	//fmt.Printf("arr0的地址:%p\n", &arr[0][1])
//	//fmt.Printf("arr1的地址:%p\n", &arr[1])
//	//fmt.Printf("arr1的地址:%p\n", &arr[1][0])
//	//arrInt := [5]int{24, 69, 80, 57, 13}
//	//BubbleSort(&arrInt)
//	//slice := fbn(10)
//	//fmt.Println(slice)
//	//str := "HELLO WORLD"
//	//slice := str[6:]
//	////arr1 := []byte(str)
//	////arr1[0] = 'z'
//	//arr1 := []rune(str)
//	//arr1[0] = '福'
//	//str = string(arr1)
//	////str = arr1
//	//fmt.Println(str)
//	//fmt.Println(slice)
//	//var arr = [5]int{1, 2, 3, 4, 5}
//	//var a []int
//	//a = arr[:]
//	//var slice = make([]int, 10)
//	//copy(slice, a)
//	////slice = a
//	//slice[0] = 29
//	//fmt.Println(slice)
//	//fmt.Println(a)
//	//fmt.Println(arr)
//	//var arrInt = [5]int{1, 2, 3, 4, 5}
//	//slice := arrInt[1:4]
//	//for i := 0; i < len(slice); i++ {
//	//	fmt.Printf("slice[%v]=%v\n", i, slice[i])
//	//}
//	//for i, v := range slice {
//	//	fmt.Printf("slice[%v]=%v\n", i, v)
//	//}
//	//slice2 := slice[:3]
//	//for i := 0; i < len(slice2); i++ {
//	//	fmt.Printf("slice2[%v]=%v\n", i, slice2[i])
//	//}
//	//for i, v := range slice2 {
//	//	fmt.Printf("slice2[%v]=%v\n", i, v)
//	//}
//	//slice2[2] = 88
//	//slice2 = append(slice2, 99, 100, 102, 2338)
//	//slice2 = append(slice2, slice...)
//	//slice2 = append(slice2, slice2...)
//	//fmt.Println(slice2)
//	//fmt.Println(arrInt)
//	//var arr [5]int
//	//rand.Seed(int64(time.Now().Unix()))
//	//for i := 0; i < len(arr); i++ {
//	//
//	//	arr[i] = rand.Intn(1000)
//	//}
//	//len := len(arr)
//	//var temp int
//	//fmt.Println("前:\n", arr)
//	//for i := 0; i < len/2; i++ {
//	//	temp = arr[len-1-i]
//	//	arr[len-1-i] = arr[i]
//	//	arr[i] = temp
//	//}
//	//fmt.Println("后：", arr)
//	//for i := len(arr) - 1; i >= 0; i-- {
//	//	fmt.Printf("arr[%v]=%v\n", i, arr[i])
//	//}
//
//	//var intArr [3]int
//	//fmt.Printf("%v", &intArr)
//
//	//start := time.Now().Unix()
//	//test03()
//	//end := time.Now().Unix()
//	//fmt.Printf("执行test03（）耗费时间为%v秒\n", end-start)
//	////0.1秒打印一个数
//	//	i := 0
//	//	for {
//	//		i++
//	//		fmt.Println(i)
//	//		time.Sleep(time.Millisecond * 100)
//	//		if i == 100 {
//	//			break
//	//		}
//	//	}
//	//now := time.Now()
//	//fmt.Printf("unix时间戳=%v unixnao时间戳=%v", now.Unix(), now.UnixNano())
//
//	//fmt.Printf("type =%T val= %v\n", now, now)
//	//fmt.Printf("年=%v\n", now.Year())
//	//fmt.Printf("月=%v\n", int(now.Month()))
//	//fmt.Printf("日=%v\n", now.Day())
//	//fmt.Printf("时=%v\n", now.Hour())
//	//fmt.Printf("分=%v\n", now.Minute())
//	//fmt.Printf("秒=%v\n", now.Second())
//	//datestr := fmt.Sprintf("当前年月日: %d-%d-%d %d:%d:%d \n", now.Year(),
//	//	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
//	//fmt.Println("datestr=", datestr)
//	//fmt.Printf(now.Format("01"))
//	//strArr := strings.HasSuffix(" hello,WOttRLD,OK   !", "! ")
//	//fmt.Printf("b=%v\n", strArr)
//	//for i := 0; i < len(strArr); i++ {
//	//	fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
//	//}
//
//	//str := strconv.FormatInt(123, 16)
//	//fmt.Println("123对应的二进制是：", str)
//
//	//str := string([]byte{97, 98, 99})
//	//fmt.Println("str=", str)
//
//	//var bytes = []byte("hello go")
//	//fmt.Printf("bytes=%v", bytes)
//
//	//var total float64 = 100000
//	//var num int
//	//for {
//	//	if total > 50000 {
//	//		total = total * 0.95
//	//		num++
//	//		continue
//	//	}
//	//	if total < 50000 {
//	//		total -= 1000
//	//		num++
//	//	}
//	//	if total <= 0 {
//	//		break
//	//	}
//	//}
//	//fmt.Println("一共经过多少个路口：", num)
//	//var n int
//	//var positiveCount int
//	//var NegativeCount int
//	//for {
//	//	fmt.Println("请输入一个随机整数")
//	//	fmt.Scanln(&n)
//	//	if n == 0 {
//	//		break
//	//	}
//	//	if n > 0 {
//	//		positiveCount++
//	//		continue
//	//	}
//	//	if n < 0 {
//	//		NegativeCount++
//	//	}
//	//}
//	//fmt.Println("输入的正数个数为：", positiveCount)
//	//fmt.Println("输入的负数个数为：", NegativeCount)
//
//	////打印1-100之间的奇数
//	//for i := 1; i <= 100; i++ {
//	//	if i%2 != 0 {
//	//		fmt.Println("i=", i)
//	//	} else {
//	//		continue
//	//	}
//	//}
//
//	////continue
//	//lable1:
//	//for i := 0; i < 4; i++ {
//	//	//lable2:
//	//	for j := 0; j < 10; j++ {
//	//		if j == 2 {
//	//			continue lable1
//	//		}
//	//		fmt.Println("j=", j)
//	//	}
//	//}
//
//	////break 练习
//	//var name string
//	//var pwd string
//	//for i := 1; i <= 3; i++ {
//	//	fmt.Println("请输入用户名：")
//	//	fmt.Scanln(&name)
//	//	fmt.Println("请输入密码：")
//	//	fmt.Scanln(&pwd)
//	//	if name == "张无忌" && pwd == "888" {
//	//		fmt.Printf("登录成功")
//	//		break
//	//	} else {
//	//		fmt.Printf("还有%v次机会\n", 3-i)
//	//	}
//	//}
//
//	////100以内的数相加，第一次大于20时的数是多少
//	//var sum int
//	//for i := 1; i < 100; i++ {
//	//	sum += i
//	//	if sum >= 20 {
//	//		fmt.Println("第一次大于20时的数是：", i)
//	//		break
//	//	}
//	//}
//
//	//break
//	//lable1:
//	//for i := 0; i < 4; i++ {
//	//	//lable2:
//	//	for j := 0; j < 10; j++ {
//	//		if j == 2 {
//	//			break lable1
//	//		}
//	//		fmt.Println("j=", j)
//	//	}
//	//}
//
//	//随机生成一个1-100的数 ，直到生成99这个数，看看一共用了多少次？
//	//为了生成一个随机数，还需要给RAND设置一个种子
//	//time.Now().Unix():返回一个从1970:01:01的0时0秒到现在的秒数
//	//rand.Seed(time.Now().Unix())
//	//fmt.Println(time.Now().Unix())
//	////如何随机生成1-100的整数
//	//n := rand.Intn(100) + 1
//	//fmt.Println("n=", n)
//	//var count int
//	//for {
//	//	rand.Seed(time.Now().UnixNano())
//	//	n := rand.Intn(100) + 1
//	//	count++
//	//	if n == 99 {
//	//		break
//	//	}
//	//}
//	//fmt.Printf("生成99一共用了%v次", count)
//
//	//var height float64
//	//var weight float64
//	//var weight_standard float64
//	//fmt.Println("请输入身高：")
//	//fmt.Scan(&height)
//	//fmt.Println("请输入体重：")
//	//fmt.Scan(&weight)
//	//weight_standard = (height * 1.08)
//	//fmt.Println("weight=", weight)
//	//fmt.Println("weight_standard=", weight_standard)
//	//if weight > weight_standard {
//	//	fmt.Println("您的体重超重")
//	//}
//	////打印99乘法表
//	//for i := 1; i <= 3; i++ {
//	//	for j := 1; j <= i; j++ {
//	//		fmt.Printf("%v * %v= %v  ", j, i, i*j)
//	//	}
//	//	fmt.Println()
//	//}
//
//	////打印空心菱形
//	//var totalLevel = 9
//	//for i := 1; i <= totalLevel; i++ {
//	//	//在打印*前先打印空格
//	//	for k := 1; k <= totalLevel-i; k++ {
//	//		fmt.Print(" ")
//	//	}
//	//	for j := 1; j <= 2*i-1; j++ {
//	//		if j == 1 || j == 2*i-1 {
//	//			fmt.Print("*")
//	//		} else {
//	//			fmt.Print(" ")
//	//		}
//	//	}
//	//	fmt.Println()
//	//}
//	////先打印上半部分，再打印下半部分，下半部分行数少一
//	//totalLevel2 := totalLevel - 1
//	//for i := 1; i <= totalLevel2; i++ {
//	//	//在打印*前先打印空格
//	//	for k := 1; k <= i; k++ {
//	//		fmt.Print(" ")
//	//	}
//	//	for j := 1; j <= 2*(totalLevel2-i)+1; j++ {
//	//		if j == 1 || j == 2*(totalLevel2-i)+1 {
//	//			fmt.Print("*")
//	//		} else {
//	//
//	//			fmt.Print(" ")
//	//		}
//	//	}
//	//
//	//	fmt.Println()
//	//}
//
//	//打印空闲金字塔
//	//var totalLevel = 9
//	//for i := 1; i <= totalLevel; i++ {
//	//	//在打印*前先打印空格
//	//	for k := 1; k <= totalLevel-i; k++ {
//	//		fmt.Print(" ")
//	//	}
//	//	for j := 1; j <= 2*i-1; j++ {
//	//		if j == 1 || j == 2*i-1 || i == totalLevel {
//	//			fmt.Print("*")
//	//		} else {
//	//
//	//			fmt.Print(" ")
//	//		}
//	//	}
//	//	fmt.Println()
//	//}
//}
//func test(b byte) byte {
//	return b + 1
//}
