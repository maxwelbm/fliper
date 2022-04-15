package main

import (
	"fmt"
	"testing"
)

// go test -count=1 -failfast -v -run ^Test_GenLines
func Test_GenLines(t *testing.T) {

	fmt.Print(genLines("1233224323", "123323422"))

}
