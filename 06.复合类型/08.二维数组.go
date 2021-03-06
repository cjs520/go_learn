package main

import "fmt"

func main() {
	/*
	定义：
		var 变量名 [行长度][列长度]类型
	赋值
	循环取值
	初始化
	注意：还是遵循下标从0开始的
	 */

	// 两行三列的二维数组
	var arr [2][3]int
	fmt.Println(arr)
	// 第一行第二列改为111
	arr[0][1] = 111
	// 第二行第三列改为222
	arr[1][2] = 222
	fmt.Println(arr)

	// for循环取出二维数组的值
	// 外层控制行，内层控制列
	for i:=0;i<len(arr) ;i++  {
		// 列数是每一行的长度，比如二行三列，第一行的长度就是列数，第二行的长度也是列数
		for j:=0;j<len(arr[1]) ;j++  {
			fmt.Println(arr[i][j])

		}
	}

	// 二维数组的初始化
	// 三行四列，就有三个赋值的花括号，每个里面4个值
	// 如果定义的不够就会使用0填充，但是不能多
	b := [3][4]int{{1,2,3,4},{5,6,7,8},{2,3,4,5}}
	fmt.Println(b)

	// 初始化指定的,其他的为0
	c := [3][4]int{{2,3,4},{6,7,8},{4,5}}

}
