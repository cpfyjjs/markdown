package base

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa =3
var ss = "kkk"
var bb = true
const filename = "file.txt"
func main() {
	fmt.Println("Hello, World!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	euler()
	triangle()
	enums()
}
func variableZeroValue(){
	var a int
	var s string
	fmt.Println(a,s)
}

func variableInitialValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a,b,c,s =3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShort(){
	a,b,c,s := 3,4,true,"def"
	fmt.Println(a,b,c,s)

}
func euler(){
	c :=3 +4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
	fmt.Printf("%.3f\n",cmplx.Pow(math.E,1i*math.Pi)+1)
}

func triangle() {
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a +b*b)))
	fmt.Println(c)
}

func consts(){
	const filename = "abc.txt"
	const a,b =3,4
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c)
}

func add(a int,b int){

}

func enums(){
	const(
		cpp = 0
		java = 1
		python =2
		goland = 3
	)
	const(
		b = 1<< (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,python,goland)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
//程序结构








