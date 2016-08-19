package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*百格计算器*/

var arr [10][10]int

func generateBaiGe(x int) [10]int {
	var arr [10]int
	r := rand.New(rand.NewSource(time.Now().Unix() + int64(x)))
	for i := 0; i < 10; i++ {
		arr[i] = r.Intn(100)
	}
	return arr
}

// func dealString(i int) string{
//     if i<
//     s:="  "+i+"  "
// }

func main() {
	// rand_service_handler := rand_generator()
	// fmt.Printf("%d\n", <-rand_service_handler)

	arrW := generateBaiGe(1)
	arrH := generateBaiGe(2)
	fmt.Print("\t")
	for i := 0; i < 10; i++ {
		fmt.Print("\t", arrW[i])
	}
	for j := 0; j < 10; j++ {
		fmt.Print("\n", "\t", arrH[j])
		for i := 0; i < 10; i++ {
			arr[j][i] = arrW[i] + arrH[j]
			fmt.Print("\t", arr[j][i])
		}
	}
	// for i = 0; i < 10; i++ {

	// }
}
