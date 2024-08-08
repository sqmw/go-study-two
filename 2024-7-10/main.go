package main

import "fmt"

// Address 地址信息
type Address struct {
	City    string
	Country string
}

// Student 学生信息
type Student struct {
	Name    string
	Age     int
	Address // 匿名嵌套
}

// Teacher 教师信息
type Teacher struct {
	Name       string
	Subject    string
	HomeAddr   Address // 命名嵌套
	SchoolAddr Address // 命名嵌套
}

func main() {
	// 使用匿名嵌套的结构体
	student := Student{
		Name: "小明",
		Age:  18,
		Address: Address{
			City:    "北京",
			Country: "中国",
		},
	}

	// 可以直接访问匿名嵌套的字段
	fmt.Printf("学生%s来自%s,%s\n", student.Name, student.City, student.Country)

	// 使用命名嵌套的结构体
	teacher := Teacher{
		Name:    "张老师",
		Subject: "数学",
		HomeAddr: Address{
			City:    "上海",
			Country: "中国",
		},
		SchoolAddr: Address{
			City:    "广州",
			Country: "中国",
		},
	}

	// 必须通过完整路径访问命名嵌套的字段
	fmt.Printf("老师%s的家在%s,%s\n", teacher.Name, teacher.HomeAddr.City, teacher.HomeAddr.Country)
	fmt.Printf("老师%s在%s,%s的学校教%s\n", teacher.Name, teacher.SchoolAddr.City, teacher.SchoolAddr.Country, teacher.Subject)
}
