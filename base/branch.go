package base

import "fmt"

func main() {
	println("hello world !")
	fmt.Println("hello world !")
	println(eval(3,4,"+"))
	println(eval(12,4,"/"))
	println(eval(12,4,"-"))


}

func eval(a,b int,op string) int{
	var result int
	switch op{
	case "+":
		result = a +b
	case "-":
		result = a -b
	case "*":
		result = a *b
	case "/":
		result = a/b
	default:
		panic("unsupported operator: "+ op)
	}
	return result
}