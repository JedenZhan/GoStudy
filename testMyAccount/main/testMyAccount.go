package main

import (
	"fmt"
)

func main() {
	//声明变量接收输入
	key := ""
	loop := true
	for {
		fmt.Println(`
			------家庭收支软件------
			1.收支明细
			2.登记收入
			3.登记支出
			4.退出
		`)
		fmt.Print("请选择:")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----收支明细----")
		case "2":
			fmt.Println("----登记收入----")
		case "3":
			fmt.Println("----登记支出----")
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你已滚蛋...")
}