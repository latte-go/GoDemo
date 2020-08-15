package main

import (
	"GoDemo/ChanDemo"
	"GoDemo/LinkListDemo"
	"GoDemo/SliceDemo"
)

func main()  {

	//测试链表
	linklist := LinkListDemo.LinkList{}
	//linklist.Add("a")
	//linklist.Add("a")
	//linklist.Add("a")
	//linklist.Add("a")
	//linklist.Add("a")
	//linklist.Add("a")
	linklist.Add("ccc")
	linklist.ShowList()


	//测试切片
	SliceDemo.NewSlice()
	ChanDemo.ChanOpertion()

	//测试Chan
	//输出结果
	//1、 执行出chan方法
	//2、 出chan方法，开始阻塞
	//3、 执行入chan方法
	//4、 input----0xc000006038  打印传入chan前变量地址
	//5、 chan 写入完成，结束阻塞、读取chan中值
	//6、 output---0xc000006030  打印读取出chan后变量的内存地址

}