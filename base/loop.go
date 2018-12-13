package base

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(12345))
	printFile("abc.txt")
}

func convertToBin(n int) string{
	if n == 0{
		return "0"
	}

	result := ""
	for ; n > 0;n/=2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}