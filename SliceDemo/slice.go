package main

import "fmt"

//通过下标修改切片里面的值
func ModifySlice(O []string){
	O[0] = "10"
}
//通过append 内置方法修改切片里面的值
func AppendSlice(O []string){
	O = append(O,"20")
}

func NewSlice()  {
	var UserSlice = [] string {"1","2","3","4"}

	AppendSlice(UserSlice)  //掉用append内置方法修改，原切片内数值并未增加， 值传递
	fmt.Println(UserSlice)

	ModifySlice(UserSlice)  //掉用通过下标修改，原切片内数值被改变，指针传递
	fmt.Println(UserSlice)
}

func main()  {
	NewSlice()

}