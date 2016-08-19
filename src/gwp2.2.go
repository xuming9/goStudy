package main

import (
	"fmt"
)

func main() {

	// fmt.Println("a"<-2)
	fmt.Println(1 << 7)

	// // fmt.Println(split(17))
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
}
func initVar() {
	var v1 int32
	var v11, v12 int64
	var v21 float32 = 0.01
	var v31, v32 float64 = 0.31, 0.31
	var v41, v42 = 0.41, 0.42
	v51, v52 := 51, 52
	v1 = 1
	v11 = 11
	v12 = 12
	fmt.Println(v1)
	fmt.Println(v11)
	fmt.Println(v12)
	fmt.Println(v21)
	fmt.Println(v31)
	fmt.Println(v32)
	fmt.Println(v41)
	fmt.Println(v42)
	fmt.Println(v51)
	fmt.Println(v52)
}

func initMap() {
	// vMap := make(map[string]int)
	var vMap map[string]string
	vMap["one"] = "1"
	vMap["two"] = "2"

	for k, v := range vMap {
		fmt.Println("key = " + k + "  value = " + v)
	}
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
