package main

import (
	fmt "fmt"
)

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x

}

/*打印长方形*/
func changfangxing(x int, y int) {
	l, h := x, y
	for ; h > 0; h-- {
		for ; l > 0; l-- {
			fmt.Print("*")
		}
		l = x
		fmt.Println("")
	}
}

/*打印正方形*/
func zhengfangxing(x int) {
	changfangxing(x, x)
}

/*打印菱形*/
func lingxing(x int) {
	for i := 1; i <= 2*x-1; i++ {
		for j := 0; j <= abs(x-i); j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*(x-abs(x-i))-1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("----打印长方形----")
	changfangxing(10, 3)
	fmt.Println("----打印正方形----")
	zhengfangxing(2)
	fmt.Println("----打印菱形----")
	lingxing(6)

}
