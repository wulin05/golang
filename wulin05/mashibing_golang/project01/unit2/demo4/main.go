package main

import "fmt"

func main() {
	//定义浮点类型的数据
	var num1 float32 = 3.14
	fmt.Println(num1)
	//可以表示正浮点数，也可以表示负的浮点数
	var num2 float32 = -3.14
	fmt.Println(num2)
	//浮点数可以用十进制表示形式，也可以用科学计数法表示形式，E大小写都可以
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	var num5 float64 = 314e+2
	fmt.Println(num5)

	//浮点数可能会有精度的损失，所以通常情况下，建议使用: float64
	var num6 float32 = 256.00000000916
	fmt.Println(num6)
	var num7 float64 = 256.00000000916
	fmt.Println(num7)

	//golang中默认的浮点类型是: float64
	var num8 = 3.14
	fmt.Printf("num9对应的默认类型是: %T", num8)
}
