package main

import (
	"fmt"
	"golang-unit-test/helper"
)

func main() {
	test := helper.UserGradeA("username ", 12)
	fmt.Print(test)
}
