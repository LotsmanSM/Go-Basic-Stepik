package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	a := bufio.NewScanner(os.Stdin)

	_ = a.Scan()
	razd := a.Text()

	_ = a.Scan()
	str1 := a.Text()

	_ = a.Scan()
	str2 := a.Text()

	_ = a.Scan()
	str3 := a.Text()

	fmt.Print(str1, razd, str2, razd, str3)
}
