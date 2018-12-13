package base

import "fmt"

func main() {
	fmt.Println(div(12,5))
	q,_ :=div(13,3)
	println(q)
	fmt.Println(eval(3,4,"S"))

}

//函数返回多个值
func div(a int,b int)(q,r int){
	return a / b,a % b
}

//func div(a int,b int)(q,r int){
//	q = a / b
//	r = a % b
//	return
//}
func eval(a,b int,op string)(int,error){
	var result int
	switch op{
	case "+":
		result = a +b
		return result,nil
	case "-":
		result = a -b
		return result,nil
	case "*":
		result = a *b
		return result,nil
	case "/":
		result = a /b
		return result,nil
	default:
		return 0,fmt.Errorf("unsupport operation: %s",op)
	}
}

