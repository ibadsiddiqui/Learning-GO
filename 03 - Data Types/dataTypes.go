package main
import (
	"fmt"
	"math/cmplx"
)
func main()  {
	var a bool = true
	var b int = 1
	var c string = "hello world"
	var d float32 = 1.222
	var x complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
}