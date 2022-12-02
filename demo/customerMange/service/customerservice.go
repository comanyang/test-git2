package service

import "go_code/demo/customerMange/model"

// 该CustomerService,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	//什么一个字段，表示当前切片有多少个客户
	//该字段还可以作为新客户的ID+1
	customerNum int
}

// 编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到customers切片中
func (this *CustomerService) Add(customer model.Customer) bool {
	//确定一个分配ID的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//如果index == -1,说明没有这个客户
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true

}

// 根据ID查找客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历切片，this.customers
	for i, v := range this.customers {
		//找到
		if v.Id == id {
			index = i
		}
	}
	return index
}
