package logic

import "fmt"

const s string = "constant"

func VarTest2() {
	const n = 500000000
	const dd = 3e20 / n
	fmt.Println(dd)
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f + s)
}
